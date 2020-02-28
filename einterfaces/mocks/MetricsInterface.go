// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

// AddMemCacheHitCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheHitCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// AddMemCacheMissCounter provides a mock function with given fields: cacheName, amount
func (_m *MetricsInterface) AddMemCacheMissCounter(cacheName string, amount float64) {
	_m.Called(cacheName, amount)
}

// DecrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) DecrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementChannelIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementChannelIndexCounter() {
	_m.Called()
}

// IncrementClusterEventType provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementClusterEventType(eventType string) {
	_m.Called(eventType)
}

// IncrementClusterRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementClusterRequest() {
	_m.Called()
}

// IncrementEtagHitCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagHitCounter(route string) {
	_m.Called(route)
}

// IncrementEtagMissCounter provides a mock function with given fields: route
func (_m *MetricsInterface) IncrementEtagMissCounter(route string) {
	_m.Called(route)
}

// IncrementHttpError provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHttpError() {
	_m.Called()
}

// IncrementHttpRequest provides a mock function with given fields:
func (_m *MetricsInterface) IncrementHttpRequest() {
	_m.Called()
}

// IncrementLogin provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLogin() {
	_m.Called()
}

// IncrementLoginFail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementLoginFail() {
	_m.Called()
}

// IncrementMemCacheHitCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheHitCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheHitCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheHitCounterSession() {
	_m.Called()
}

// IncrementMemCacheInvalidationCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheInvalidationCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheInvalidationCounterSession() {
	_m.Called()
}

// IncrementMemCacheMissCounter provides a mock function with given fields: cacheName
func (_m *MetricsInterface) IncrementMemCacheMissCounter(cacheName string) {
	_m.Called(cacheName)
}

// IncrementMemCacheMissCounterSession provides a mock function with given fields:
func (_m *MetricsInterface) IncrementMemCacheMissCounterSession() {
	_m.Called()
}

// IncrementPostBroadcast provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostBroadcast() {
	_m.Called()
}

// IncrementPostCreate provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostCreate() {
	_m.Called()
}

// IncrementPostFileAttachment provides a mock function with given fields: count
func (_m *MetricsInterface) IncrementPostFileAttachment(count int) {
	_m.Called(count)
}

// IncrementPostIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostIndexCounter() {
	_m.Called()
}

// IncrementPostSentEmail provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentEmail() {
	_m.Called()
}

// IncrementPostSentPush provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostSentPush() {
	_m.Called()
}

// IncrementPostsSearchCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementPostsSearchCounter() {
	_m.Called()
}

// IncrementUserIndexCounter provides a mock function with given fields:
func (_m *MetricsInterface) IncrementUserIndexCounter() {
	_m.Called()
}

// IncrementWebSocketBroadcast provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebSocketBroadcast(eventType string) {
	_m.Called(eventType)
}

// IncrementWebSocketBroadcastBufferSize provides a mock function with given fields: hub, amount
func (_m *MetricsInterface) IncrementWebSocketBroadcastBufferSize(hub string, amount float64) {
	_m.Called(hub, amount)
}

// IncrementWebhookPost provides a mock function with given fields:
func (_m *MetricsInterface) IncrementWebhookPost() {
	_m.Called()
}

// IncrementWebsocketEvent provides a mock function with given fields: eventType
func (_m *MetricsInterface) IncrementWebsocketEvent(eventType string) {
	_m.Called(eventType)
}

// ObserveApiEndpointDuration provides a mock function with given fields: endpoint, method, elapsed
func (_m *MetricsInterface) ObserveApiEndpointDuration(endpoint string, method string, elapsed float64) {
	_m.Called(endpoint, method, elapsed)
}

// ObserveClusterRequestDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveClusterRequestDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveHttpRequestDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObserveHttpRequestDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObservePluginApiDuration provides a mock function with given fields: pluginID, apiName, success, elapsed
func (_m *MetricsInterface) ObservePluginApiDuration(pluginID string, apiName string, success bool, elapsed float64) {
	_m.Called(pluginID, apiName, success, elapsed)
}

// ObservePluginHookDuration provides a mock function with given fields: pluginID, hookName, success, elapsed
func (_m *MetricsInterface) ObservePluginHookDuration(pluginID string, hookName string, success bool, elapsed float64) {
	_m.Called(pluginID, hookName, success, elapsed)
}

// ObservePluginMultiHookDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePluginMultiHookDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObservePluginMultiHookIterationDuration provides a mock function with given fields: pluginID, elapsed
func (_m *MetricsInterface) ObservePluginMultiHookIterationDuration(pluginID string, elapsed float64) {
	_m.Called(pluginID, elapsed)
}

// ObservePostsSearchDuration provides a mock function with given fields: elapsed
func (_m *MetricsInterface) ObservePostsSearchDuration(elapsed float64) {
	_m.Called(elapsed)
}

// ObserveStoreMethodDuration provides a mock function with given fields: method, success, elapsed
func (_m *MetricsInterface) ObserveStoreMethodDuration(method string, success string, elapsed float64) {
	_m.Called(method, success, elapsed)
}

// StartServer provides a mock function with given fields:
func (_m *MetricsInterface) StartServer() {
	_m.Called()
}

// StopServer provides a mock function with given fields:
func (_m *MetricsInterface) StopServer() {
	_m.Called()
}
