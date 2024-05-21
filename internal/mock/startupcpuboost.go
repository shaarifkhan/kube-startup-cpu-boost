// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kube-startup-cpu-boost/internal/boost (interfaces: StartupCPUBoost)
//
// Generated by this command:
//
//	mockgen -package mock --copyright_file hack/boilerplate.go.txt --destination internal/mock/startupcpuboost.go github.com/google/kube-startup-cpu-boost/internal/boost StartupCPUBoost
//

package mock

import (
	context "context"
	reflect "reflect"

	boost "github.com/google/kube-startup-cpu-boost/internal/boost"
	duration "github.com/google/kube-startup-cpu-boost/internal/boost/duration"
	resource "github.com/google/kube-startup-cpu-boost/internal/boost/resource"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockStartupCPUBoost is a mock of StartupCPUBoost interface.
type MockStartupCPUBoost struct {
	ctrl     *gomock.Controller
	recorder *MockStartupCPUBoostMockRecorder
}

// MockStartupCPUBoostMockRecorder is the mock recorder for MockStartupCPUBoost.
type MockStartupCPUBoostMockRecorder struct {
	mock *MockStartupCPUBoost
}

// NewMockStartupCPUBoost creates a new mock instance.
func NewMockStartupCPUBoost(ctrl *gomock.Controller) *MockStartupCPUBoost {
	mock := &MockStartupCPUBoost{ctrl: ctrl}
	mock.recorder = &MockStartupCPUBoostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStartupCPUBoost) EXPECT() *MockStartupCPUBoostMockRecorder {
	return m.recorder
}

// DeletePod mocks base method.
func (m *MockStartupCPUBoost) DeletePod(arg0 context.Context, arg1 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockStartupCPUBoostMockRecorder) DeletePod(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockStartupCPUBoost)(nil).DeletePod), arg0, arg1)
}

// DurationPolicies mocks base method.
func (m *MockStartupCPUBoost) DurationPolicies() map[string]duration.Policy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DurationPolicies")
	ret0, _ := ret[0].(map[string]duration.Policy)
	return ret0
}

// DurationPolicies indicates an expected call of DurationPolicies.
func (mr *MockStartupCPUBoostMockRecorder) DurationPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DurationPolicies", reflect.TypeOf((*MockStartupCPUBoost)(nil).DurationPolicies))
}

// Matches mocks base method.
func (m *MockStartupCPUBoost) Matches(arg0 *v1.Pod) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Matches", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Matches indicates an expected call of Matches.
func (mr *MockStartupCPUBoostMockRecorder) Matches(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Matches", reflect.TypeOf((*MockStartupCPUBoost)(nil).Matches), arg0)
}

// Name mocks base method.
func (m *MockStartupCPUBoost) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockStartupCPUBoostMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockStartupCPUBoost)(nil).Name))
}

// Namespace mocks base method.
func (m *MockStartupCPUBoost) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockStartupCPUBoostMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockStartupCPUBoost)(nil).Namespace))
}

// Pod mocks base method.
func (m *MockStartupCPUBoost) Pod(arg0 string) (*v1.Pod, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod", arg0)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Pod indicates an expected call of Pod.
func (mr *MockStartupCPUBoostMockRecorder) Pod(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockStartupCPUBoost)(nil).Pod), arg0)
}

// ResourcePolicy mocks base method.
func (m *MockStartupCPUBoost) ResourcePolicy(arg0 string) (resource.ContainerPolicy, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourcePolicy", arg0)
	ret0, _ := ret[0].(resource.ContainerPolicy)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ResourcePolicy indicates an expected call of ResourcePolicy.
func (mr *MockStartupCPUBoostMockRecorder) ResourcePolicy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourcePolicy", reflect.TypeOf((*MockStartupCPUBoost)(nil).ResourcePolicy), arg0)
}

// RevertResources mocks base method.
func (m *MockStartupCPUBoost) RevertResources(arg0 context.Context, arg1 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevertResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevertResources indicates an expected call of RevertResources.
func (mr *MockStartupCPUBoostMockRecorder) RevertResources(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevertResources", reflect.TypeOf((*MockStartupCPUBoost)(nil).RevertResources), arg0, arg1)
}

// Stats mocks base method.
func (m *MockStartupCPUBoost) Stats() boost.StartupCPUBoostStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(boost.StartupCPUBoostStats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockStartupCPUBoostMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockStartupCPUBoost)(nil).Stats))
}

// UpsertPod mocks base method.
func (m *MockStartupCPUBoost) UpsertPod(arg0 context.Context, arg1 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPod indicates an expected call of UpsertPod.
func (mr *MockStartupCPUBoostMockRecorder) UpsertPod(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPod", reflect.TypeOf((*MockStartupCPUBoost)(nil).UpsertPod), arg0, arg1)
}

// ValidatePolicy mocks base method.
func (m *MockStartupCPUBoost) ValidatePolicy(arg0 context.Context, arg1 string) []*v1.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePolicy", arg0, arg1)
	ret0, _ := ret[0].([]*v1.Pod)
	return ret0
}

// ValidatePolicy indicates an expected call of ValidatePolicy.
func (mr *MockStartupCPUBoostMockRecorder) ValidatePolicy(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePolicy", reflect.TypeOf((*MockStartupCPUBoost)(nil).ValidatePolicy), arg0, arg1)
}
