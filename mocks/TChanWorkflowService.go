package mocks

import minions "code.uber.internal/devexp/minions-client-go.git/.gen/go/minions"
import mock "github.com/stretchr/testify/mock"
import thrift "github.com/uber/tchannel-go/thrift"

// TChanWorkflowService is an autogenerated mock type for the TChanWorkflowService type
type TChanWorkflowService struct {
	mock.Mock
}

// PollForActivityTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForActivityTask(ctx thrift.Context, pollRequest *minions.PollForActivityTaskRequest) (*minions.PollForActivityTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *minions.PollForActivityTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.PollForActivityTaskRequest) *minions.PollForActivityTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*minions.PollForActivityTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *minions.PollForActivityTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForDecisionTask provides a mock function with given fields: ctx, pollRequest
func (_m *TChanWorkflowService) PollForDecisionTask(ctx thrift.Context, pollRequest *minions.PollForDecisionTaskRequest) (*minions.PollForDecisionTaskResponse, error) {
	ret := _m.Called(ctx, pollRequest)

	var r0 *minions.PollForDecisionTaskResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.PollForDecisionTaskRequest) *minions.PollForDecisionTaskResponse); ok {
		r0 = rf(ctx, pollRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*minions.PollForDecisionTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *minions.PollForDecisionTaskRequest) error); ok {
		r1 = rf(ctx, pollRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordActivityTaskHeartbeat provides a mock function with given fields: ctx, heartbeatRequest
func (_m *TChanWorkflowService) RecordActivityTaskHeartbeat(ctx thrift.Context, heartbeatRequest *minions.RecordActivityTaskHeartbeatRequest) (*minions.RecordActivityTaskHeartbeatResponse, error) {
	ret := _m.Called(ctx, heartbeatRequest)

	var r0 *minions.RecordActivityTaskHeartbeatResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.RecordActivityTaskHeartbeatRequest) *minions.RecordActivityTaskHeartbeatResponse); ok {
		r0 = rf(ctx, heartbeatRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*minions.RecordActivityTaskHeartbeatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *minions.RecordActivityTaskHeartbeatRequest) error); ok {
		r1 = rf(ctx, heartbeatRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RespondActivityTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondActivityTaskCompleted(ctx thrift.Context, completeRequest *minions.RespondActivityTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.RespondActivityTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskFailed provides a mock function with given fields: ctx, failRequest
func (_m *TChanWorkflowService) RespondActivityTaskFailed(ctx thrift.Context, failRequest *minions.RespondActivityTaskFailedRequest) error {
	ret := _m.Called(ctx, failRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.RespondActivityTaskFailedRequest) error); ok {
		r0 = rf(ctx, failRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondDecisionTaskCompleted provides a mock function with given fields: ctx, completeRequest
func (_m *TChanWorkflowService) RespondDecisionTaskCompleted(ctx thrift.Context, completeRequest *minions.RespondDecisionTaskCompletedRequest) error {
	ret := _m.Called(ctx, completeRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.RespondDecisionTaskCompletedRequest) error); ok {
		r0 = rf(ctx, completeRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartWorkflowExecution provides a mock function with given fields: ctx, startRequest
func (_m *TChanWorkflowService) StartWorkflowExecution(ctx thrift.Context, startRequest *minions.StartWorkflowExecutionRequest) (*minions.StartWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, startRequest)

	var r0 *minions.StartWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(thrift.Context, *minions.StartWorkflowExecutionRequest) *minions.StartWorkflowExecutionResponse); ok {
		r0 = rf(ctx, startRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*minions.StartWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(thrift.Context, *minions.StartWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, startRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ minions.TChanWorkflowService = (*TChanWorkflowService)(nil)