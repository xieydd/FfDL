// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/IBM/FfDL/commons/service (interfaces: LifecycleManager_GetTrainingLogStreamClient,LifecycleManagerClient)

package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/IBM/FfDL/commons/service"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// Mock of LifecycleManager_GetTrainingLogStreamClient interface
type MockLifecycleManager_GetTrainingLogStreamClient struct {
	ctrl     *gomock.Controller
	recorder *_MockLifecycleManager_GetTrainingLogStreamClientRecorder
}

// Recorder for MockLifecycleManager_GetTrainingLogStreamClient (not exported)
type _MockLifecycleManager_GetTrainingLogStreamClientRecorder struct {
	mock *MockLifecycleManager_GetTrainingLogStreamClient
}

func NewMockLifecycleManager_GetTrainingLogStreamClient(ctrl *gomock.Controller) *MockLifecycleManager_GetTrainingLogStreamClient {
	mock := &MockLifecycleManager_GetTrainingLogStreamClient{ctrl: ctrl}
	mock.recorder = &_MockLifecycleManager_GetTrainingLogStreamClientRecorder{mock}
	return mock
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) EXPECT() *_MockLifecycleManager_GetTrainingLogStreamClientRecorder {
	return _m.recorder
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) CloseSend() error {
	ret := _m.ctrl.Call(_m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) CloseSend() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloseSend")
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Context")
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) Header() (metadata.MD, error) {
	ret := _m.ctrl.Call(_m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) Header() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Header")
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) Recv() (*service.TrainerLogStreamResponse, error) {
	ret := _m.ctrl.Call(_m, "Recv")
	ret0, _ := ret[0].(*service.TrainerLogStreamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) Recv() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Recv")
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) RecvMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "RecvMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecvMsg", arg0)
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) SendMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SendMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendMsg", arg0)
}

func (_m *MockLifecycleManager_GetTrainingLogStreamClient) Trailer() metadata.MD {
	ret := _m.ctrl.Call(_m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

func (_mr *_MockLifecycleManager_GetTrainingLogStreamClientRecorder) Trailer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Trailer")
}

// Mock of LifecycleManagerClient interface
type MockLifecycleManagerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockLifecycleManagerClientRecorder
}

// Recorder for MockLifecycleManagerClient (not exported)
type _MockLifecycleManagerClientRecorder struct {
	mock *MockLifecycleManagerClient
}

func NewMockLifecycleManagerClient(ctrl *gomock.Controller) *MockLifecycleManagerClient {
	mock := &MockLifecycleManagerClient{ctrl: ctrl}
	mock.recorder = &_MockLifecycleManagerClientRecorder{mock}
	return mock
}

func (_m *MockLifecycleManagerClient) EXPECT() *_MockLifecycleManagerClientRecorder {
	return _m.recorder
}

func (_m *MockLifecycleManagerClient) DeployTrainingJob(_param0 context.Context, _param1 *service.JobDeploymentRequest, _param2 ...grpc.CallOption) (*service.JobDeploymentResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeployTrainingJob", _s...)
	ret0, _ := ret[0].(*service.JobDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManagerClientRecorder) DeployTrainingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeployTrainingJob", _s...)
}

func (_m *MockLifecycleManagerClient) GetTrainingLogStream(_param0 context.Context, _param1 *service.TrainerContainerInfosRequest, _param2 ...grpc.CallOption) (service.LifecycleManager_GetTrainingLogStreamClient, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetTrainingLogStream", _s...)
	ret0, _ := ret[0].(service.LifecycleManager_GetTrainingLogStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManagerClientRecorder) GetTrainingLogStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTrainingLogStream", _s...)
}

func (_m *MockLifecycleManagerClient) HaltTrainingJob(_param0 context.Context, _param1 *service.JobHaltRequest, _param2 ...grpc.CallOption) (*service.JobHaltResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "HaltTrainingJob", _s...)
	ret0, _ := ret[0].(*service.JobHaltResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManagerClientRecorder) HaltTrainingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HaltTrainingJob", _s...)
}

func (_m *MockLifecycleManagerClient) KillTrainingJob(_param0 context.Context, _param1 *service.JobKillRequest, _param2 ...grpc.CallOption) (*service.JobKillResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "KillTrainingJob", _s...)
	ret0, _ := ret[0].(*service.JobKillResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManagerClientRecorder) KillTrainingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillTrainingJob", _s...)
}

// 	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
func (_m *MockLifecycleManagerClient) GetMetrics(_param0 context.Context, _param1 *service.GetMetricsRequest, _param2 ...grpc.CallOption) (*service.GetMetricsResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetMetrics", _s...)
	ret0, _ := ret[0].(*service.GetMetricsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLifecycleManagerClientRecorder) GetMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetrics", _s...)
}

//CreateEventEndpoint ... mock implementation
func (_m *MockLifecycleManagerClient) CreateEventEndpoint(_param0 context.Context, _param1 *service.CreateEventEndpointRequest, _param2 ...grpc.CallOption) (*service.CreateEventEndpointResponse, error) {
	return nil, nil
}

//DeleteEventEndpoint ... mock impl
func (_m *MockLifecycleManagerClient) DeleteEventEndpoint(_param0 context.Context, _param1 *service.DeleteEventEndpointRequest, _param2 ...grpc.CallOption) (*service.DeleteEventEndpointResponse, error) {
	return nil, nil
}

//GetEventTypeEndpoints ... mock impl
func (_m *MockLifecycleManagerClient) GetEventTypeEndpoints(_param0 context.Context, _param1 *service.GetEventTypeEndpointsRequest, _param2 ...grpc.CallOption) (*service.GetEventTypeEndpointsResponse, error) {
	return nil, nil
}

//GetEventEndpoint ... mock impl
func (_m *MockLifecycleManagerClient) GetEventEndpoint(_param0 context.Context, _param1 *service.GetEventEndpointRequest, _param2 ...grpc.CallOption) (*service.GetEventEndpointResponse, error) {
	return nil, nil
}
