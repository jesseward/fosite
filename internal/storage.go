// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: Storage)

package internal

import (
	gomock "github.com/golang/mock/gomock"
	client "github.com/ory-am/fosite/client"
)

// Mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockStorageRecorder
}

// Recorder for MockStorage (not exported)
type _MockStorageRecorder struct {
	mock *MockStorage
}

func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &_MockStorageRecorder{mock}
	return mock
}

func (_m *MockStorage) EXPECT() *_MockStorageRecorder {
	return _m.recorder
}

func (_m *MockStorage) GetClient(_param0 string) (client.Client, error) {
	ret := _m.ctrl.Call(_m, "GetClient", _param0)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStorageRecorder) GetClient(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClient", arg0)
}
