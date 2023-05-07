// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libp2p/go-libp2p/core/network (interfaces: StreamManagementScope)

// Package mocknetwork is a generated GoMock package.
package mocknetwork

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	network "github.com/libp2p/go-libp2p/core/network"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
)

// MockStreamManagementScope is a mock of StreamManagementScope interface.
type MockStreamManagementScope struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagementScopeMockRecorder
}

// MockStreamManagementScopeMockRecorder is the mock recorder for MockStreamManagementScope.
type MockStreamManagementScopeMockRecorder struct {
	mock *MockStreamManagementScope
}

// NewMockStreamManagementScope creates a new mock instance.
func NewMockStreamManagementScope(ctrl *gomock.Controller) *MockStreamManagementScope {
	mock := &MockStreamManagementScope{ctrl: ctrl}
	mock.recorder = &MockStreamManagementScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManagementScope) EXPECT() *MockStreamManagementScopeMockRecorder {
	return m.recorder
}

// BeginSpan mocks base method.
func (m *MockStreamManagementScope) BeginSpan() (network.ResourceScopeSpan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSpan")
	ret0, _ := ret[0].(network.ResourceScopeSpan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSpan indicates an expected call of BeginSpan.
func (mr *MockStreamManagementScopeMockRecorder) BeginSpan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSpan", reflect.TypeOf((*MockStreamManagementScope)(nil).BeginSpan))
}

// Done mocks base method.
func (m *MockStreamManagementScope) Done() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Done")
}

// Done indicates an expected call of Done.
func (mr *MockStreamManagementScopeMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockStreamManagementScope)(nil).Done))
}

// PeerScope mocks base method.
func (m *MockStreamManagementScope) PeerScope() network.PeerScope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeerScope")
	ret0, _ := ret[0].(network.PeerScope)
	return ret0
}

// PeerScope indicates an expected call of PeerScope.
func (mr *MockStreamManagementScopeMockRecorder) PeerScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerScope", reflect.TypeOf((*MockStreamManagementScope)(nil).PeerScope))
}

// ProtocolScope mocks base method.
func (m *MockStreamManagementScope) ProtocolScope() network.ProtocolScope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtocolScope")
	ret0, _ := ret[0].(network.ProtocolScope)
	return ret0
}

// ProtocolScope indicates an expected call of ProtocolScope.
func (mr *MockStreamManagementScopeMockRecorder) ProtocolScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtocolScope", reflect.TypeOf((*MockStreamManagementScope)(nil).ProtocolScope))
}

// ReleaseMemory mocks base method.
func (m *MockStreamManagementScope) ReleaseMemory(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseMemory", arg0)
}

// ReleaseMemory indicates an expected call of ReleaseMemory.
func (mr *MockStreamManagementScopeMockRecorder) ReleaseMemory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseMemory", reflect.TypeOf((*MockStreamManagementScope)(nil).ReleaseMemory), arg0)
}

// ReserveMemory mocks base method.
func (m *MockStreamManagementScope) ReserveMemory(arg0 int, arg1 byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveMemory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveMemory indicates an expected call of ReserveMemory.
func (mr *MockStreamManagementScopeMockRecorder) ReserveMemory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveMemory", reflect.TypeOf((*MockStreamManagementScope)(nil).ReserveMemory), arg0, arg1)
}

// ServiceScope mocks base method.
func (m *MockStreamManagementScope) ServiceScope() network.ServiceScope {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceScope")
	ret0, _ := ret[0].(network.ServiceScope)
	return ret0
}

// ServiceScope indicates an expected call of ServiceScope.
func (mr *MockStreamManagementScopeMockRecorder) ServiceScope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceScope", reflect.TypeOf((*MockStreamManagementScope)(nil).ServiceScope))
}

// SetProtocol mocks base method.
func (m *MockStreamManagementScope) SetProtocol(arg0 protocol.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProtocol", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProtocol indicates an expected call of SetProtocol.
func (mr *MockStreamManagementScopeMockRecorder) SetProtocol(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProtocol", reflect.TypeOf((*MockStreamManagementScope)(nil).SetProtocol), arg0)
}

// SetService mocks base method.
func (m *MockStreamManagementScope) SetService(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetService indicates an expected call of SetService.
func (mr *MockStreamManagementScopeMockRecorder) SetService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetService", reflect.TypeOf((*MockStreamManagementScope)(nil).SetService), arg0)
}

// Stat mocks base method.
func (m *MockStreamManagementScope) Stat() network.ScopeStat {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(network.ScopeStat)
	return ret0
}

// Stat indicates an expected call of Stat.
func (mr *MockStreamManagementScopeMockRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockStreamManagementScope)(nil).Stat))
}
