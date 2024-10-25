// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nxenon/xquic-go (interfaces: PacketHandler)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/nxenon/xquic-go -destination mock_packet_handler_test.go github.com/nxenon/xquic-go PacketHandler
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	qerr "github.com/nxenon/xquic-go/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacketHandler is a mock of PacketHandler interface.
type MockPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerMockRecorder
}

// MockPacketHandlerMockRecorder is the mock recorder for MockPacketHandler.
type MockPacketHandlerMockRecorder struct {
	mock *MockPacketHandler
}

// NewMockPacketHandler creates a new mock instance.
func NewMockPacketHandler(ctrl *gomock.Controller) *MockPacketHandler {
	mock := &MockPacketHandler{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketHandler) EXPECT() *MockPacketHandlerMockRecorder {
	return m.recorder
}

// closeWithTransportError mocks base method.
func (m *MockPacketHandler) closeWithTransportError(arg0 qerr.TransportErrorCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "closeWithTransportError", arg0)
}

// closeWithTransportError indicates an expected call of closeWithTransportError.
func (mr *MockPacketHandlerMockRecorder) closeWithTransportError(arg0 any) *MockPacketHandlercloseWithTransportErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeWithTransportError", reflect.TypeOf((*MockPacketHandler)(nil).closeWithTransportError), arg0)
	return &MockPacketHandlercloseWithTransportErrorCall{Call: call}
}

// MockPacketHandlercloseWithTransportErrorCall wrap *gomock.Call
type MockPacketHandlercloseWithTransportErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketHandlercloseWithTransportErrorCall) Return() *MockPacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketHandlercloseWithTransportErrorCall) Do(f func(qerr.TransportErrorCode)) *MockPacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketHandlercloseWithTransportErrorCall) DoAndReturn(f func(qerr.TransportErrorCode)) *MockPacketHandlercloseWithTransportErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// destroy mocks base method.
func (m *MockPacketHandler) destroy(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy.
func (mr *MockPacketHandlerMockRecorder) destroy(arg0 any) *MockPacketHandlerdestroyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockPacketHandler)(nil).destroy), arg0)
	return &MockPacketHandlerdestroyCall{Call: call}
}

// MockPacketHandlerdestroyCall wrap *gomock.Call
type MockPacketHandlerdestroyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketHandlerdestroyCall) Return() *MockPacketHandlerdestroyCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketHandlerdestroyCall) Do(f func(error)) *MockPacketHandlerdestroyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketHandlerdestroyCall) DoAndReturn(f func(error)) *MockPacketHandlerdestroyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// handlePacket mocks base method.
func (m *MockPacketHandler) handlePacket(arg0 receivedPacket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket.
func (mr *MockPacketHandlerMockRecorder) handlePacket(arg0 any) *MockPacketHandlerhandlePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockPacketHandler)(nil).handlePacket), arg0)
	return &MockPacketHandlerhandlePacketCall{Call: call}
}

// MockPacketHandlerhandlePacketCall wrap *gomock.Call
type MockPacketHandlerhandlePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPacketHandlerhandlePacketCall) Return() *MockPacketHandlerhandlePacketCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPacketHandlerhandlePacketCall) Do(f func(receivedPacket)) *MockPacketHandlerhandlePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPacketHandlerhandlePacketCall) DoAndReturn(f func(receivedPacket)) *MockPacketHandlerhandlePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
