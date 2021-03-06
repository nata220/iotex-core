// Code generated by MockGen. DO NOT EDIT.
// Source: ./consensus/consensusfsm/context.go

// Package consensusfsm is a generated GoMock package.
package consensusfsm

import (
	gomock "github.com/golang/mock/gomock"
	go_fsm "github.com/iotexproject/go-fsm"
	zap "go.uber.org/zap"
	reflect "reflect"
	time "time"
)

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// IsStaleEvent mocks base method
func (m *MockContext) IsStaleEvent(arg0 *ConsensusEvent) bool {
	ret := m.ctrl.Call(m, "IsStaleEvent", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStaleEvent indicates an expected call of IsStaleEvent
func (mr *MockContextMockRecorder) IsStaleEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStaleEvent", reflect.TypeOf((*MockContext)(nil).IsStaleEvent), arg0)
}

// IsFutureEvent mocks base method
func (m *MockContext) IsFutureEvent(arg0 *ConsensusEvent) bool {
	ret := m.ctrl.Call(m, "IsFutureEvent", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFutureEvent indicates an expected call of IsFutureEvent
func (mr *MockContextMockRecorder) IsFutureEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFutureEvent", reflect.TypeOf((*MockContext)(nil).IsFutureEvent), arg0)
}

// IsStaleUnmatchedEvent mocks base method
func (m *MockContext) IsStaleUnmatchedEvent(arg0 *ConsensusEvent) bool {
	ret := m.ctrl.Call(m, "IsStaleUnmatchedEvent", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStaleUnmatchedEvent indicates an expected call of IsStaleUnmatchedEvent
func (mr *MockContextMockRecorder) IsStaleUnmatchedEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStaleUnmatchedEvent", reflect.TypeOf((*MockContext)(nil).IsStaleUnmatchedEvent), arg0)
}

// Logger mocks base method
func (m *MockContext) Logger() *zap.Logger {
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(*zap.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockContextMockRecorder) Logger() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockContext)(nil).Logger))
}

// Height mocks base method
func (m *MockContext) Height() uint64 {
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Height indicates an expected call of Height
func (mr *MockContextMockRecorder) Height() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockContext)(nil).Height))
}

// NewConsensusEvent mocks base method
func (m *MockContext) NewConsensusEvent(arg0 go_fsm.EventType, arg1 interface{}) *ConsensusEvent {
	ret := m.ctrl.Call(m, "NewConsensusEvent", arg0, arg1)
	ret0, _ := ret[0].(*ConsensusEvent)
	return ret0
}

// NewConsensusEvent indicates an expected call of NewConsensusEvent
func (mr *MockContextMockRecorder) NewConsensusEvent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsensusEvent", reflect.TypeOf((*MockContext)(nil).NewConsensusEvent), arg0, arg1)
}

// NewBackdoorEvt mocks base method
func (m *MockContext) NewBackdoorEvt(arg0 go_fsm.State) *ConsensusEvent {
	ret := m.ctrl.Call(m, "NewBackdoorEvt", arg0)
	ret0, _ := ret[0].(*ConsensusEvent)
	return ret0
}

// NewBackdoorEvt indicates an expected call of NewBackdoorEvt
func (mr *MockContextMockRecorder) NewBackdoorEvt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBackdoorEvt", reflect.TypeOf((*MockContext)(nil).NewBackdoorEvt), arg0)
}

// Broadcast mocks base method
func (m *MockContext) Broadcast(arg0 interface{}) {
	m.ctrl.Call(m, "Broadcast", arg0)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockContextMockRecorder) Broadcast(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockContext)(nil).Broadcast), arg0)
}

// Prepare mocks base method
func (m *MockContext) Prepare() (bool, interface{}, bool, bool, time.Duration, error) {
	ret := m.ctrl.Call(m, "Prepare")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(interface{})
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(bool)
	ret4, _ := ret[4].(time.Duration)
	ret5, _ := ret[5].(error)
	return ret0, ret1, ret2, ret3, ret4, ret5
}

// Prepare indicates an expected call of Prepare
func (mr *MockContextMockRecorder) Prepare() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockContext)(nil).Prepare))
}

// NewProposalEndorsement mocks base method
func (m *MockContext) NewProposalEndorsement(arg0 interface{}) (interface{}, error) {
	ret := m.ctrl.Call(m, "NewProposalEndorsement", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProposalEndorsement indicates an expected call of NewProposalEndorsement
func (mr *MockContextMockRecorder) NewProposalEndorsement(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProposalEndorsement", reflect.TypeOf((*MockContext)(nil).NewProposalEndorsement), arg0)
}

// NewLockEndorsement mocks base method
func (m *MockContext) NewLockEndorsement(arg0 interface{}) (interface{}, error) {
	ret := m.ctrl.Call(m, "NewLockEndorsement", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLockEndorsement indicates an expected call of NewLockEndorsement
func (mr *MockContextMockRecorder) NewLockEndorsement(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLockEndorsement", reflect.TypeOf((*MockContext)(nil).NewLockEndorsement), arg0)
}

// NewPreCommitEndorsement mocks base method
func (m *MockContext) NewPreCommitEndorsement(arg0 interface{}) (interface{}, error) {
	ret := m.ctrl.Call(m, "NewPreCommitEndorsement", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPreCommitEndorsement indicates an expected call of NewPreCommitEndorsement
func (mr *MockContextMockRecorder) NewPreCommitEndorsement(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPreCommitEndorsement", reflect.TypeOf((*MockContext)(nil).NewPreCommitEndorsement), arg0)
}

// Commit mocks base method
func (m *MockContext) Commit(arg0 interface{}) (bool, error) {
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockContextMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockContext)(nil).Commit), arg0)
}
