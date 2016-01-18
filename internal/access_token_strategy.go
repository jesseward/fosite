// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite/handler/core (interfaces: AccessTokenStrategy)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	context "golang.org/x/net/context"
	http "net/http"
)

// Mock of AccessTokenStrategy interface
type MockAccessTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessTokenStrategyRecorder
}

// Recorder for MockAccessTokenStrategy (not exported)
type _MockAccessTokenStrategyRecorder struct {
	mock *MockAccessTokenStrategy
}

func NewMockAccessTokenStrategy(ctrl *gomock.Controller) *MockAccessTokenStrategy {
	mock := &MockAccessTokenStrategy{ctrl: ctrl}
	mock.recorder = &_MockAccessTokenStrategyRecorder{mock}
	return mock
}

func (_m *MockAccessTokenStrategy) EXPECT() *_MockAccessTokenStrategyRecorder {
	return _m.recorder
}

func (_m *MockAccessTokenStrategy) GenerateAccessToken(_param0 context.Context, _param1 *http.Request, _param2 fosite.Requester) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GenerateAccessToken", _param0, _param1, _param2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAccessTokenStrategyRecorder) GenerateAccessToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateAccessToken", arg0, arg1, arg2)
}

func (_m *MockAccessTokenStrategy) ValidateAccessToken(_param0 string, _param1 context.Context, _param2 *http.Request, _param3 fosite.Requester) (string, error) {
	ret := _m.ctrl.Call(_m, "ValidateAccessToken", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessTokenStrategyRecorder) ValidateAccessToken(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateAccessToken", arg0, arg1, arg2, arg3)
}