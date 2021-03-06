// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-cli/stemcell (interfaces: CloudStemcell,Manager)

package mocks

import (
	stemcell "github.com/cloudfoundry/bosh-cli/stemcell"
	ui "github.com/cloudfoundry/bosh-cli/ui"
	gomock "github.com/golang/mock/gomock"
)

// Mock of CloudStemcell interface
type MockCloudStemcell struct {
	ctrl     *gomock.Controller
	recorder *_MockCloudStemcellRecorder
}

// Recorder for MockCloudStemcell (not exported)
type _MockCloudStemcellRecorder struct {
	mock *MockCloudStemcell
}

func NewMockCloudStemcell(ctrl *gomock.Controller) *MockCloudStemcell {
	mock := &MockCloudStemcell{ctrl: ctrl}
	mock.recorder = &_MockCloudStemcellRecorder{mock}
	return mock
}

func (_m *MockCloudStemcell) EXPECT() *_MockCloudStemcellRecorder {
	return _m.recorder
}

func (_m *MockCloudStemcell) CID() string {
	ret := _m.ctrl.Call(_m, "CID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCloudStemcellRecorder) CID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CID")
}

func (_m *MockCloudStemcell) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudStemcellRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockCloudStemcell) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCloudStemcellRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockCloudStemcell) PromoteAsCurrent() error {
	ret := _m.ctrl.Call(_m, "PromoteAsCurrent")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudStemcellRecorder) PromoteAsCurrent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PromoteAsCurrent")
}

func (_m *MockCloudStemcell) Version() string {
	ret := _m.ctrl.Call(_m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockCloudStemcellRecorder) Version() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Version")
}

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) DeleteUnused(_param0 ui.Stage) error {
	ret := _m.ctrl.Call(_m, "DeleteUnused", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) DeleteUnused(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUnused", arg0)
}

func (_m *MockManager) FindCurrent() ([]stemcell.CloudStemcell, error) {
	ret := _m.ctrl.Call(_m, "FindCurrent")
	ret0, _ := ret[0].([]stemcell.CloudStemcell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) FindCurrent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindCurrent")
}

func (_m *MockManager) FindUnused() ([]stemcell.CloudStemcell, error) {
	ret := _m.ctrl.Call(_m, "FindUnused")
	ret0, _ := ret[0].([]stemcell.CloudStemcell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) FindUnused() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindUnused")
}

func (_m *MockManager) Upload(_param0 stemcell.ExtractedStemcell, _param1 ui.Stage) (stemcell.CloudStemcell, error) {
	ret := _m.ctrl.Call(_m, "Upload", _param0, _param1)
	ret0, _ := ret[0].(stemcell.CloudStemcell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Upload", arg0, arg1)
}
