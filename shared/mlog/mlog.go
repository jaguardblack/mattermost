// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Package mlog provides a simple wrapper around Logr.
package mlog

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/mattermost/logr/v2"
	logrcfg "github.com/mattermost/logr/v2/config"
)

const (
	ShutdownTimeout                = time.Second * 15
	FlushTimeout                   = time.Second * 15
	DefaultMaxQueueSize            = 1000
	DefaultMetricsUpdateFreqMillis = 15000
)

type LoggerIFace interface {
	IsLevelEnabled(Level) bool
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...Field)
	Critical(string, ...Field)
	Log(Level, string, ...Field)
	LogM([]Level, string, ...Field)
}

// Type and function aliases from Logr to limit the spread of dependencies.
type Field = logr.Field
type Level = logr.Level
type Option = logr.Option
type Target = logr.Target
type TargetInfo = logr.TargetInfo
type LogRec = logr.LogRec
type LogCloner = logr.LogCloner
type MetricsCollector = logr.MetricsCollector
type TargetCfg = logrcfg.TargetCfg
type Sugar = logr.Sugar

// LoggerConfiguration is a map of LogTarget configurations.
type LoggerConfiguration map[string]TargetCfg

func (lc LoggerConfiguration) Append(cfg LoggerConfiguration) {
	for k, v := range cfg {
		lc[k] = v
	}
}

func (lc LoggerConfiguration) toTargetCfg() map[string]logrcfg.TargetCfg {
	tcfg := make(map[string]logrcfg.TargetCfg)
	for k, v := range lc {
		tcfg[k] = v
	}
	return tcfg
}

// Any picks the best supported field type based on type of val.
// For best performance when passing a struct (or struct pointer),
// implement `logr.LogWriter` on the struct, otherwise reflection
// will be used to generate a string representation.
var Any = logr.Any

// Int64 constructs a field containing a key and Int64 value.
var Int64 = logr.Int64

// Int32 constructs a field containing a key and Int32 value.
var Int32 = logr.Int32

// Int constructs a field containing a key and Int value.
var Int = logr.Int

// Uint64 constructs a field containing a key and Uint64 value.
var Uint64 = logr.Uint64

// Uint32 constructs a field containing a key and Uint32 value.
var Uint32 = logr.Uint32

// Uint constructs a field containing a key and Uint value.
var Uint = logr.Uint

// Float64 constructs a field containing a key and Float64 value.
var Float64 = logr.Float64

// Float32 constructs a field containing a key and Float32 value.
var Float32 = logr.Float32

// String constructs a field containing a key and String value.
var String = logr.String

// Stringer constructs a field containing a key and a fmt.Stringer value.
// The fmt.Stringer's `String` method is called lazily.
var Stringer = logr.Stringer

// Err constructs a field containing a default key ("error") and error value.
var Err = logr.Err

// NamedErr constructs a field containing a key and error value.
var NamedErr = logr.NamedErr

// Bool constructs a field containing a key and bool value.
var Bool = logr.Bool

// Time constructs a field containing a key and time.Time value.
var Time = logr.Time

// Duration constructs a field containing a key and time.Duration value.
var Duration = logr.Duration

// Millis constructs a field containing a key and timestamp value.
// The timestamp is expected to be milliseconds since Jan 1, 1970 UTC.
var Millis = logr.Millis

// Array constructs a field containing a key and array value.
var Array = logr.Array

// Map constructs a field containing a key and map value.
var Map = logr.Map

// Logger provides a thin wrapper around a Logr instance. This is a struct instead of an interface
// so that there are no allocations on the heap each interface method invocation. Normally not
// something to be concerned about, but logging calls for disabled levels should have as little CPU
// and memory impact as possible. Most of these wrapper calls will be inlined as well.
type Logger struct {
	log        *logr.Logger
	lockConfig *int32
}

// NewLogger creates a new Logger instance which can be configured via `(*Logger).Configure`.
func NewLogger(options ...Option) *Logger {
	options = append(options, logr.StackFilter(logr.GetPackageName("NewLogger")))

	lgr, _ := logr.New(options...)
	log := lgr.NewLogger()
	var lockConfig int32

	return &Logger{
		log:        &log,
		lockConfig: &lockConfig,
	}
}

// Configure provides a new configuration for this logger.
// Zero or more sources of config can be provided:
//   cfgFile    - path to file containing JSON
//   cfgEscaped - JSON string probably from ENV var
//
// For each case JSON containing log targets is provided. Target name collisions are resolved
// using the following precedence:
//     cfgFile > cfgEscaped
func (l *Logger) Configure(cfgFile string, cfgEscaped string) error {
	if atomic.LoadInt32(l.lockConfig) != 0 {
		return ConfigurationLockedError{}
	}

	cfgMap := make(LoggerConfiguration)

	// Add config from file
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return fmt.Errorf("error reading logger config file %s: %w", cfgFile, err)
		}

		var mapCfgFile LoggerConfiguration
		if err := json.Unmarshal(b, &mapCfgFile); err != nil {
			return fmt.Errorf("error decoding logger config file %s: %w", cfgFile, err)
		}
		cfgMap.Append(mapCfgFile)
	}

	// Add config from escaped json string
	if cfgEscaped != "" {
		var mapCfgEscaped LoggerConfiguration
		if err := json.Unmarshal([]byte(cfgEscaped), &mapCfgEscaped); err != nil {
			return fmt.Errorf("error decoding logger config as escaped json: %w", err)
		}
		cfgMap.Append(mapCfgEscaped)
	}

	if len(cfgMap) == 0 {
		return nil
	}

	return logrcfg.ConfigureTargets(l.log.Logr(), cfgMap.toTargetCfg(), nil)
}

// ConfigureTargets provides a new configuration for this logger via a `LoggerConfig` map.
// Typically `mlog.Configure` is used instead which accepts JSON formatted configuration.
func (l *Logger) ConfigureTargets(cfg LoggerConfiguration) error {
	if atomic.LoadInt32(l.lockConfig) != 0 {
		return ConfigurationLockedError{}
	}
	return logrcfg.ConfigureTargets(l.log.Logr(), cfg.toTargetCfg(), nil)
}

// LockConfiguration disallows further configuration changes until `UnlockConfiguration`
// is called. The previous locked stated is returned.
func (l *Logger) LockConfiguration() bool {
	old := atomic.SwapInt32(l.lockConfig, 1)
	return old != 0
}

// UnlockConfiguration allows configuration changes. The previous locked stated is returned.
func (l *Logger) UnlockConfiguration() bool {
	old := atomic.SwapInt32(l.lockConfig, 0)
	return old != 0
}

// IsConfigurationLocked returns the current state of the configuration lock.
func (l *Logger) IsConfigurationLocked() bool {
	return atomic.LoadInt32(l.lockConfig) != 0
}

// With creates a new Logger with the specified fields. This is a light-weight
// operation and can be called on demand.
func (l *Logger) With(fields ...Field) *Logger {
	logWith := l.log.With(fields...)
	return &Logger{
		log:        &logWith,
		lockConfig: l.lockConfig,
	}
}

// IsLevelEnabled returns true only if at least one log target is
// configured to emit the specified log level. Use this check when
// gathering the log info may be expensive.
//
// Note, transformations and serializations done via fields are already
// lazily evaluated and don't require this check beforehand.
func (l *Logger) IsLevelEnabled(level Level) bool {
	return l.log.IsLevelEnabled(level)
}

// Log emits the log record for any targets configured for the specified level.
func (l *Logger) Log(level Level, msg string, fields ...Field) {
	l.log.Log(level, msg, fields...)
}

// LogM emits the log record for any targets configured for the specified levels.
// Equivalent to calling `Log` once for each level.
func (l *Logger) LogM(levels []Level, msg string, fields ...Field) {
	l.log.LogM(levels, msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Trace` level.
func (l *Logger) Trace(msg string, fields ...Field) {
	l.log.Trace(msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Debug` level.
func (l *Logger) Debug(msg string, fields ...Field) {
	l.log.Debug(msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Info` level.
func (l *Logger) Info(msg string, fields ...Field) {
	l.log.Info(msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Warn` level.
func (l *Logger) Warn(msg string, fields ...Field) {
	l.log.Warn(msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Error` level.
func (l *Logger) Error(msg string, fields ...Field) {
	l.log.Error(msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Critical` level.
func (l *Logger) Critical(msg string, fields ...Field) {
	l.log.Log(LvlCritical, msg, fields...)
}

// Convenience method equivalent to calling `Log` with the `Fatal` level,
// followed by `os.Exit(1)`.
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.log.Log(logr.Fatal, msg, fields...)
	_ = l.Shutdown()
	os.Exit(1)
}

// HasTargets returns true if at least one log target has been added.
func (l *Logger) HasTargets() bool {
	return l.log.Logr().HasTargets()
}

// StdLogger creates a standard logger backed by this logger.
// All log records are output with the specified level.
func (l *Logger) StdLogger(level Level) *log.Logger {
	return l.log.StdLogger(level)
}

// StdLogWriter returns a writer that can be hooked up to the output of a golang standard logger
// anything written will be interpreted as log entries and passed to this logger.
func (l *Logger) StdLogWriter() io.Writer {
	return &logWriter{
		logger: l,
	}
}

// RedirectStdLog redirects output from the standard library's package-global logger
// to this logger at the specified level and with zero or more Field's. Since this logger already
// handles caller annotations, timestamps, etc., it automatically disables the standard
// library's annotations and prefixing.
// A function is returned that restores the original prefix and flags and resets the standard
// library's output to os.Stdout.
func (l *Logger) RedirectStdLog(level Level, fields ...Field) func() {
	return l.log.Logr().RedirectStdLog(level, fields...)
}

// RemoveTargets safely removes one or more targets based on the filtering method.
// `f` should return true to delete the target, false to keep it.
// When removing a target, best effort is made to write any queued log records before
// closing, with cxt determining how much time can be spent in total.
// Note, keep the timeout short since this method blocks certain logging operations.
func (l *Logger) RemoveTargets(ctx context.Context, f func(ti TargetInfo) bool) error {
	return l.log.Logr().RemoveTargets(ctx, f)
}

// SetMetricsCollector sets (or resets) the metrics collector to be used for gathering
// metrics for all targets. Only targets added after this call will use the collector.
//
// To ensure all targets use a collector, use the `SetMetricsCollector` option when
// creating the Logger instead, or configure/reconfigure the Logger after calling this method.
func (l *Logger) SetMetricsCollector(collector MetricsCollector, updateFrequencyMillis int64) {
	l.log.Logr().SetMetricsCollector(collector, updateFrequencyMillis)
}

// Sugar creates a new `Logger` with a less structured API. Any fields are preserved.
func (l *Logger) Sugar(fields ...Field) Sugar {
	return l.log.Sugar(fields...)
}

// Flush forces all targets to write out any queued log records with a default timeout.
func (l *Logger) Flush() error {
	ctx, cancel := context.WithTimeout(context.Background(), FlushTimeout)
	defer cancel()
	return l.log.Logr().FlushWithTimeout(ctx)
}

// Flush forces all targets to write out any queued log records with the specfified timeout.
func (l *Logger) FlushWithTimeout(ctx context.Context) error {
	return l.log.Logr().FlushWithTimeout(ctx)
}

// Shutdown shuts down the logger after making best efforts to flush any
// remaining records.
func (l *Logger) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return l.log.Logr().ShutdownWithTimeout(ctx)
}

// Shutdown shuts down the logger after making best efforts to flush any
// remaining records.
func (l *Logger) ShutdownWithTimeout(ctx context.Context) error {
	return l.log.Logr().ShutdownWithTimeout(ctx)
}

// GetPackageName reduces a fully qualified function name to the package name
// By sirupsen: https://github.com/sirupsen/logrus/blob/master/entry.go
func GetPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}
	return f
}

type logWriter struct {
	logger *Logger
}

func (lw *logWriter) Write(p []byte) (int, error) {
	lw.logger.Info(string(p))
	return len(p), nil
}

// ConfigurationLockedError is returned when one of a logger's configuration APIs is called
// while the configuration is locked.
type ConfigurationLockedError struct {
}

func (e ConfigurationLockedError) Error() string {
	return "configuration is locked"
}
