// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package cache

import (
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack/v5"

	"github.com/bluele/gcache"
)

// LFU cache implementation using the Ristretto library.
type LFU struct {
	name                   string
	size                   int
	cache                  gcache.Cache
	defaultExpiry          time.Duration
	invalidateClusterEvent string
}

// LFUOptions contains options for initializing LFU cache
type LFUOptions struct {
	Name                   string
	Size                   int
	DefaultExpiry          time.Duration
	InvalidateClusterEvent string
}

// NewLFU creates an LFU of the given size.
func NewLFU(opts *LFUOptions) Cache {
	lfu := gcache.New(opts.Size).LFU().Build()

	return &LFU{
		name:                   opts.Name,
		size:                   opts.Size,
		cache:                  lfu,
		defaultExpiry:          opts.DefaultExpiry,
		invalidateClusterEvent: opts.InvalidateClusterEvent,
	}
}

// Purge is used to completely clear the cache.
func (l *LFU) Purge() error {
	l.cache.Purge()
	return nil
}

// Set adds the given key and value to the store without an expiry. If the key already exists,
// it will overwrite the previous value.
func (l *LFU) Set(key string, value interface{}) error {
	return l.set(key, value, 0)
}

// SetWithDefaultExpiry adds the given key and value to the store with the default expiry. If
// the key already exists, it will overwrite the previoous value
func (l *LFU) SetWithDefaultExpiry(key string, value interface{}) error {
	return l.SetWithExpiry(key, value, l.defaultExpiry)
}

// SetWithExpiry adds the given key and value to the cache with the given expiry. If the key
// already exists, it will overwrite the previoous value
func (l *LFU) SetWithExpiry(key string, value interface{}, ttl time.Duration) error {
	return l.set(key, value, ttl)
}

// Get the content stored in the cache for the given key, and decode it into the value interface.
// return ErrKeyNotFound if the key is missing from the cache
func (l *LFU) Get(key string, value interface{}) error {
	return l.get(key, value)
}

// Remove deletes the value for a key.
func (l *LFU) Remove(key string) error {
	return l.remove(key)
}

// Keys returns a slice of the keys in the cache.
func (l *LFU) Keys() ([]string, error) {
	keys := l.cache.Keys(false)
	var keysStr = make([]string, len(keys))
	for i := range keys {
		keysStr[i] = keys[i].(string)
	}
	return keysStr, nil
}

// Len returns the number of items in the cache.
func (l *LFU) Len() (int, error) {
	return l.cache.Len(false), nil
}

// GetInvalidateClusterEvent returns the cluster event configured when this cache was created.
func (l *LFU) GetInvalidateClusterEvent() string {
	return l.invalidateClusterEvent
}

// Name returns the name of the cache
func (l *LFU) Name() string {
	return l.name
}

func (l *LFU) set(key string, value interface{}, ttl time.Duration) error {
	var buf []byte
	var err error

	// We use a fast path for hot structs.
	if msgpVal, ok := value.(msgp.Marshaler); ok {
		buf, err = msgpVal.MarshalMsg(nil)
		if err != nil {
			return err
		}
	} else {
		// Slow path for other structs.
		buf, err = msgpack.Marshal(value)
		if err != nil {
			return err
		}
	}

	var ok error
	if ttl != 0 {
		ok = l.cache.SetWithExpire(key, buf, ttl)
	} else {
		ok = l.cache.Set(key, buf)
	}

	return ok
}

func (l *LFU) get(key string, value interface{}) error {
	val, found := l.cache.Get(key)

	if found == gcache.KeyNotFoundError {
		return ErrKeyNotFound
	}

	// We use a fast path for hot structs.
	if msgpVal, ok := value.(msgp.Unmarshaler); ok {
		_, err := msgpVal.UnmarshalMsg(val.([]byte))
		return err
	}

	// This is ugly and makes the cache package aware of the model package.
	// But this is due to 2 things.
	// 1. The msgp package works on methods on structs rather than functions.
	// 2. Our cache interface passes pointers to empty pointers, and not pointers
	// to values. This is mainly how all our model structs are passed around.
	// It might be technically possible to use values _just_ for hot structs
	// like these and then return a pointer while returning from the cache function,
	// but it will make the codebase inconsistent, and has some edge-cases to take care of.
	switch v := value.(type) {
	case **model.User:
		var u model.User
		_, err := u.UnmarshalMsg(val.([]byte))
		*v = &u
		return err
	case **model.Session:
		var s model.Session
		_, err := s.UnmarshalMsg(val.([]byte))
		*v = &s
		return err
	}

	// Slow path for other structs.
	return msgpack.Unmarshal(val.([]byte), value)
}

func (l *LFU) remove(key string) error {
	l.cache.Remove(key)
	return nil
}
