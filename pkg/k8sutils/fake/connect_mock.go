// Code generated by MockGen. DO NOT EDIT.
// Source: connect.go

// Package fake is a generated GoMock package.
package fake

import (
	config "keptn-contrib/job-executor-service/pkg/config"
	k8sutils "keptn-contrib/job-executor-service/pkg/k8sutils"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v0_2_0 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

// MockK8s is a mock of K8s interface.
type MockK8s struct {
	ctrl     *gomock.Controller
	recorder *MockK8sMockRecorder
}

// MockK8sMockRecorder is the mock recorder for MockK8s.
type MockK8sMockRecorder struct {
	mock *MockK8s
}

// NewMockK8s creates a new mock instance.
func NewMockK8s(ctrl *gomock.Controller) *MockK8s {
	mock := &MockK8s{ctrl: ctrl}
	mock.recorder = &MockK8sMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8s) EXPECT() *MockK8sMockRecorder {
	return m.recorder
}

// AwaitK8sJobDone mocks base method.
func (m *MockK8s) AwaitK8sJobDone(jobName string, maxPollDuration, pollIntervalInSeconds int, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AwaitK8sJobDone", jobName, maxPollDuration, pollIntervalInSeconds, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// AwaitK8sJobDone indicates an expected call of AwaitK8sJobDone.
func (mr *MockK8sMockRecorder) AwaitK8sJobDone(jobName, maxPollDuration, pollIntervalInSeconds, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AwaitK8sJobDone", reflect.TypeOf((*MockK8s)(nil).AwaitK8sJobDone), jobName, maxPollDuration, pollIntervalInSeconds, namespace)
}

// ConnectToCluster mocks base method.
func (m *MockK8s) ConnectToCluster() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToCluster")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectToCluster indicates an expected call of ConnectToCluster.
func (mr *MockK8sMockRecorder) ConnectToCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToCluster", reflect.TypeOf((*MockK8s)(nil).ConnectToCluster))
}

// CreateK8sJob mocks base method.
func (m *MockK8s) CreateK8sJob(jobName string, action *config.Action, task config.Task, eventData *v0_2_0.EventData, jobSettings k8sutils.JobSettings, jsonEventData interface{}, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateK8sJob", jobName, action, task, eventData, jobSettings, jsonEventData, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateK8sJob indicates an expected call of CreateK8sJob.
func (mr *MockK8sMockRecorder) CreateK8sJob(jobName, action, task, eventData, jobSettings, jsonEventData, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateK8sJob", reflect.TypeOf((*MockK8s)(nil).CreateK8sJob), jobName, action, task, eventData, jobSettings, jsonEventData, namespace)
}

// DeleteK8sJob mocks base method.
func (m *MockK8s) DeleteK8sJob(jobName, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteK8sJob", jobName, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteK8sJob indicates an expected call of DeleteK8sJob.
func (mr *MockK8sMockRecorder) DeleteK8sJob(jobName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteK8sJob", reflect.TypeOf((*MockK8s)(nil).DeleteK8sJob), jobName, namespace)
}

// GetLogsOfPod mocks base method.
func (m *MockK8s) GetLogsOfPod(jobName, namespace string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogsOfPod", jobName, namespace)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogsOfPod indicates an expected call of GetLogsOfPod.
func (mr *MockK8sMockRecorder) GetLogsOfPod(jobName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogsOfPod", reflect.TypeOf((*MockK8s)(nil).GetLogsOfPod), jobName, namespace)
}
