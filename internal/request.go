// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: Requester)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	client "github.com/ory-am/fosite/client"
	url "net/url"
	time "time"
)

// Mock of Requester interface
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *_MockRequesterRecorder
}

// Recorder for MockRequester (not exported)
type _MockRequesterRecorder struct {
	mock *MockRequester
}

func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &_MockRequesterRecorder{mock}
	return mock
}

func (_m *MockRequester) EXPECT() *_MockRequesterRecorder {
	return _m.recorder
}

func (_m *MockRequester) GetClient() client.Client {
	ret := _m.ctrl.Call(_m, "GetClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClient")
}

func (_m *MockRequester) GetGrantedScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetGrantedScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetGrantedScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGrantedScopes")
}

func (_m *MockRequester) GetRequestForm() url.Values {
	ret := _m.ctrl.Call(_m, "GetRequestForm")
	ret0, _ := ret[0].(url.Values)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetRequestForm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestForm")
}

func (_m *MockRequester) GetRequestedAt() time.Time {
	ret := _m.ctrl.Call(_m, "GetRequestedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetRequestedAt() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRequestedAt")
}

func (_m *MockRequester) GetScopes() fosite.Arguments {
	ret := _m.ctrl.Call(_m, "GetScopes")
	ret0, _ := ret[0].(fosite.Arguments)
	return ret0
}

func (_mr *_MockRequesterRecorder) GetScopes() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetScopes")
}

func (_m *MockRequester) GetSession() interface{} {
	ret := _m.ctrl.Call(_m, "GetSession")
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockRequesterRecorder) GetSession() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSession")
}

func (_m *MockRequester) GrantScope(_param0 string) {
	_m.ctrl.Call(_m, "GrantScope", _param0)
}

func (_mr *_MockRequesterRecorder) GrantScope(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GrantScope", arg0)
}

func (_m *MockRequester) SetScopes(_param0 fosite.Arguments) {
	_m.ctrl.Call(_m, "SetScopes", _param0)
}

func (_mr *_MockRequesterRecorder) SetScopes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetScopes", arg0)
}

func (_m *MockRequester) SetSession(_param0 interface{}) {
	_m.ctrl.Call(_m, "SetSession", _param0)
}

func (_mr *_MockRequesterRecorder) SetSession(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetSession", arg0)
}
