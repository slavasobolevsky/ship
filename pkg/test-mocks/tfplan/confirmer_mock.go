// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/replicatedhq/ship/pkg/lifecycle/terraform/tfplan (interfaces: PlanConfirmer)

// Package tfplan is a generated GoMock package.
package tfplan

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/replicatedhq/ship/pkg/api"
	daemontypes "github.com/replicatedhq/ship/pkg/lifecycle/daemon/daemontypes"
	tfplan "github.com/replicatedhq/ship/pkg/lifecycle/terraform/tfplan"
)

// MockPlanConfirmer is a mock of PlanConfirmer interface
type MockPlanConfirmer struct {
	ctrl     *gomock.Controller
	recorder *MockPlanConfirmerMockRecorder
}

// MockPlanConfirmerMockRecorder is the mock recorder for MockPlanConfirmer
type MockPlanConfirmerMockRecorder struct {
	mock *MockPlanConfirmer
}

// NewMockPlanConfirmer creates a new mock instance
func NewMockPlanConfirmer(ctrl *gomock.Controller) *MockPlanConfirmer {
	mock := &MockPlanConfirmer{ctrl: ctrl}
	mock.recorder = &MockPlanConfirmerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlanConfirmer) EXPECT() *MockPlanConfirmerMockRecorder {
	return m.recorder
}

// ConfirmPlan mocks base method
func (m *MockPlanConfirmer) ConfirmPlan(arg0 context.Context, arg1 string, arg2 api.Release, arg3 chan bool) (bool, error) {
	ret := m.ctrl.Call(m, "ConfirmPlan", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmPlan indicates an expected call of ConfirmPlan
func (mr *MockPlanConfirmerMockRecorder) ConfirmPlan(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmPlan", reflect.TypeOf((*MockPlanConfirmer)(nil).ConfirmPlan), arg0, arg1, arg2, arg3)
}

// WithStatusReceiver mocks base method
func (m *MockPlanConfirmer) WithStatusReceiver(arg0 daemontypes.StatusReceiver) tfplan.PlanConfirmer {
	ret := m.ctrl.Call(m, "WithStatusReceiver", arg0)
	ret0, _ := ret[0].(tfplan.PlanConfirmer)
	return ret0
}

// WithStatusReceiver indicates an expected call of WithStatusReceiver
func (mr *MockPlanConfirmerMockRecorder) WithStatusReceiver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStatusReceiver", reflect.TypeOf((*MockPlanConfirmer)(nil).WithStatusReceiver), arg0)
}
