// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nxenon/h3spacex (interfaces: Packer)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/nxenon/h3spacex -destination mock_packer_test.go github.com/nxenon/h3spacex Packer
//

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	ackhandler "github.com/nxenon/h3spacex/internal/ackhandler"
	protocol "github.com/nxenon/h3spacex/internal/protocol"
	qerr "github.com/nxenon/h3spacex/internal/qerr"
	gomock "go.uber.org/mock/gomock"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// AppendPacket mocks base method.
func (m *MockPacker) AppendPacket(arg0 *packetBuffer, arg1 protocol.ByteCount, arg2 protocol.Version) (shortHeaderPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendPacket indicates an expected call of AppendPacket.
func (mr *MockPackerMockRecorder) AppendPacket(arg0, arg1, arg2 any) *MockPackerAppendPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendPacket", reflect.TypeOf((*MockPacker)(nil).AppendPacket), arg0, arg1, arg2)
	return &MockPackerAppendPacketCall{Call: call}
}

// MockPackerAppendPacketCall wrap *gomock.Call
type MockPackerAppendPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerAppendPacketCall) Return(arg0 shortHeaderPacket, arg1 error) *MockPackerAppendPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerAppendPacketCall) Do(f func(*packetBuffer, protocol.ByteCount, protocol.Version) (shortHeaderPacket, error)) *MockPackerAppendPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerAppendPacketCall) DoAndReturn(f func(*packetBuffer, protocol.ByteCount, protocol.Version) (shortHeaderPacket, error)) *MockPackerAppendPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MaybePackProbePacket mocks base method.
func (m *MockPacker) MaybePackProbePacket(arg0 protocol.EncryptionLevel, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePackProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePackProbePacket indicates an expected call of MaybePackProbePacket.
func (mr *MockPackerMockRecorder) MaybePackProbePacket(arg0, arg1, arg2 any) *MockPackerMaybePackProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePackProbePacket", reflect.TypeOf((*MockPacker)(nil).MaybePackProbePacket), arg0, arg1, arg2)
	return &MockPackerMaybePackProbePacketCall{Call: call}
}

// MockPackerMaybePackProbePacketCall wrap *gomock.Call
type MockPackerMaybePackProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerMaybePackProbePacketCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerMaybePackProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerMaybePackProbePacketCall) Do(f func(protocol.EncryptionLevel, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerMaybePackProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerMaybePackProbePacketCall) DoAndReturn(f func(protocol.EncryptionLevel, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerMaybePackProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackAckOnlyPacket mocks base method.
func (m *MockPacker) PackAckOnlyPacket(arg0 protocol.ByteCount, arg1 protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackAckOnlyPacket", arg0, arg1)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackAckOnlyPacket indicates an expected call of PackAckOnlyPacket.
func (mr *MockPackerMockRecorder) PackAckOnlyPacket(arg0, arg1 any) *MockPackerPackAckOnlyPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackAckOnlyPacket", reflect.TypeOf((*MockPacker)(nil).PackAckOnlyPacket), arg0, arg1)
	return &MockPackerPackAckOnlyPacketCall{Call: call}
}

// MockPackerPackAckOnlyPacketCall wrap *gomock.Call
type MockPackerPackAckOnlyPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackAckOnlyPacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackAckOnlyPacketCall) Do(f func(protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackAckOnlyPacketCall) DoAndReturn(f func(protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackAckOnlyPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackApplicationClose mocks base method.
func (m *MockPacker) PackApplicationClose(arg0 *qerr.ApplicationError, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackApplicationClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackApplicationClose indicates an expected call of PackApplicationClose.
func (mr *MockPackerMockRecorder) PackApplicationClose(arg0, arg1, arg2 any) *MockPackerPackApplicationCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackApplicationClose", reflect.TypeOf((*MockPacker)(nil).PackApplicationClose), arg0, arg1, arg2)
	return &MockPackerPackApplicationCloseCall{Call: call}
}

// MockPackerPackApplicationCloseCall wrap *gomock.Call
type MockPackerPackApplicationCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackApplicationCloseCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackApplicationCloseCall) Do(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackApplicationCloseCall) DoAndReturn(f func(*qerr.ApplicationError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackApplicationCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackCoalescedPacket mocks base method.
func (m *MockPacker) PackCoalescedPacket(arg0 bool, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackCoalescedPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackCoalescedPacket indicates an expected call of PackCoalescedPacket.
func (mr *MockPackerMockRecorder) PackCoalescedPacket(arg0, arg1, arg2 any) *MockPackerPackCoalescedPacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackCoalescedPacket", reflect.TypeOf((*MockPacker)(nil).PackCoalescedPacket), arg0, arg1, arg2)
	return &MockPackerPackCoalescedPacketCall{Call: call}
}

// MockPackerPackCoalescedPacketCall wrap *gomock.Call
type MockPackerPackCoalescedPacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackCoalescedPacketCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackCoalescedPacketCall) Do(f func(bool, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackCoalescedPacketCall) DoAndReturn(f func(bool, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackCoalescedPacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackConnectionClose mocks base method.
func (m *MockPacker) PackConnectionClose(arg0 *qerr.TransportError, arg1 protocol.ByteCount, arg2 protocol.Version) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackConnectionClose", arg0, arg1, arg2)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackConnectionClose indicates an expected call of PackConnectionClose.
func (mr *MockPackerMockRecorder) PackConnectionClose(arg0, arg1, arg2 any) *MockPackerPackConnectionCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackConnectionClose", reflect.TypeOf((*MockPacker)(nil).PackConnectionClose), arg0, arg1, arg2)
	return &MockPackerPackConnectionCloseCall{Call: call}
}

// MockPackerPackConnectionCloseCall wrap *gomock.Call
type MockPackerPackConnectionCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackConnectionCloseCall) Return(arg0 *coalescedPacket, arg1 error) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackConnectionCloseCall) Do(f func(*qerr.TransportError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackConnectionCloseCall) DoAndReturn(f func(*qerr.TransportError, protocol.ByteCount, protocol.Version) (*coalescedPacket, error)) *MockPackerPackConnectionCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PackMTUProbePacket mocks base method.
func (m *MockPacker) PackMTUProbePacket(arg0 ackhandler.Frame, arg1 protocol.ByteCount, arg2 protocol.Version) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackMTUProbePacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackMTUProbePacket indicates an expected call of PackMTUProbePacket.
func (mr *MockPackerMockRecorder) PackMTUProbePacket(arg0, arg1, arg2 any) *MockPackerPackMTUProbePacketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackMTUProbePacket", reflect.TypeOf((*MockPacker)(nil).PackMTUProbePacket), arg0, arg1, arg2)
	return &MockPackerPackMTUProbePacketCall{Call: call}
}

// MockPackerPackMTUProbePacketCall wrap *gomock.Call
type MockPackerPackMTUProbePacketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerPackMTUProbePacketCall) Return(arg0 shortHeaderPacket, arg1 *packetBuffer, arg2 error) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerPackMTUProbePacketCall) Do(f func(ackhandler.Frame, protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerPackMTUProbePacketCall) DoAndReturn(f func(ackhandler.Frame, protocol.ByteCount, protocol.Version) (shortHeaderPacket, *packetBuffer, error)) *MockPackerPackMTUProbePacketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetToken mocks base method.
func (m *MockPacker) SetToken(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockPackerMockRecorder) SetToken(arg0 any) *MockPackerSetTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockPacker)(nil).SetToken), arg0)
	return &MockPackerSetTokenCall{Call: call}
}

// MockPackerSetTokenCall wrap *gomock.Call
type MockPackerSetTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockPackerSetTokenCall) Return() *MockPackerSetTokenCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockPackerSetTokenCall) Do(f func([]byte)) *MockPackerSetTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockPackerSetTokenCall) DoAndReturn(f func([]byte)) *MockPackerSetTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
