// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nxenon/h3spacex (interfaces: AckFrameSource)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/nxenon/h3spacex -destination mock_ack_frame_source_test.go github.com/nxenon/h3spacex AckFrameSource
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/nxenon/h3spacex/internal/protocol"
	wire "github.com/nxenon/h3spacex/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockAckFrameSource is a mock of AckFrameSource interface.
type MockAckFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockAckFrameSourceMockRecorder
}

// MockAckFrameSourceMockRecorder is the mock recorder for MockAckFrameSource.
type MockAckFrameSourceMockRecorder struct {
	mock *MockAckFrameSource
}

// NewMockAckFrameSource creates a new mock instance.
func NewMockAckFrameSource(ctrl *gomock.Controller) *MockAckFrameSource {
	mock := &MockAckFrameSource{ctrl: ctrl}
	mock.recorder = &MockAckFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAckFrameSource) EXPECT() *MockAckFrameSourceMockRecorder {
	return m.recorder
}

// GetAckFrame mocks base method.
func (m *MockAckFrameSource) GetAckFrame(arg0 protocol.EncryptionLevel, arg1 bool) *wire.AckFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckFrame", arg0, arg1)
	ret0, _ := ret[0].(*wire.AckFrame)
	return ret0
}

// GetAckFrame indicates an expected call of GetAckFrame.
func (mr *MockAckFrameSourceMockRecorder) GetAckFrame(arg0, arg1 any) *MockAckFrameSourceGetAckFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckFrame", reflect.TypeOf((*MockAckFrameSource)(nil).GetAckFrame), arg0, arg1)
	return &MockAckFrameSourceGetAckFrameCall{Call: call}
}

// MockAckFrameSourceGetAckFrameCall wrap *gomock.Call
type MockAckFrameSourceGetAckFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAckFrameSourceGetAckFrameCall) Return(arg0 *wire.AckFrame) *MockAckFrameSourceGetAckFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAckFrameSourceGetAckFrameCall) Do(f func(protocol.EncryptionLevel, bool) *wire.AckFrame) *MockAckFrameSourceGetAckFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAckFrameSourceGetAckFrameCall) DoAndReturn(f func(protocol.EncryptionLevel, bool) *wire.AckFrame) *MockAckFrameSourceGetAckFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
