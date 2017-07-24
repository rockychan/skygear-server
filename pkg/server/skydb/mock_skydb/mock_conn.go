// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/skygeario/skygear-server/pkg/server/skydb (interfaces: Conn)

package mock_skydb

import (
	gomock "github.com/golang/mock/gomock"
	skydb "github.com/skygeario/skygear-server/pkg/server/skydb"
	reflect "reflect"
	time "time"
)

// MockConn is a mock of Conn interface
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConn) EXPECT() *MockConnMockRecorder {
	return _m.recorder
}

// AddRelation mocks base method
func (_m *MockConn) AddRelation(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "AddRelation", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRelation indicates an expected call of AddRelation
func (_mr *MockConnMockRecorder) AddRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AddRelation", reflect.TypeOf((*MockConn)(nil).AddRelation), arg0, arg1, arg2)
}

// AssignRoles mocks base method
func (_m *MockConn) AssignRoles(_param0 []string, _param1 []string) error {
	ret := _m.ctrl.Call(_m, "AssignRoles", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignRoles indicates an expected call of AssignRoles
func (_mr *MockConnMockRecorder) AssignRoles(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "AssignRoles", reflect.TypeOf((*MockConn)(nil).AssignRoles), arg0, arg1)
}

// Close mocks base method
func (_m *MockConn) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockConnMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}

// CreateUser mocks base method
func (_m *MockConn) CreateUser(_param0 *skydb.AuthInfo) error {
	ret := _m.ctrl.Call(_m, "CreateUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (_mr *MockConnMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateUser", reflect.TypeOf((*MockConn)(nil).CreateUser), arg0)
}

// DeleteDevice mocks base method
func (_m *MockConn) DeleteDevice(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteDevice", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice
func (_mr *MockConnMockRecorder) DeleteDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteDevice", reflect.TypeOf((*MockConn)(nil).DeleteDevice), arg0)
}

// DeleteDevicesByToken mocks base method
func (_m *MockConn) DeleteDevicesByToken(_param0 string, _param1 time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteDevicesByToken", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevicesByToken indicates an expected call of DeleteDevicesByToken
func (_mr *MockConnMockRecorder) DeleteDevicesByToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteDevicesByToken", reflect.TypeOf((*MockConn)(nil).DeleteDevicesByToken), arg0, arg1)
}

// DeleteEmptyDevicesByTime mocks base method
func (_m *MockConn) DeleteEmptyDevicesByTime(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "DeleteEmptyDevicesByTime", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmptyDevicesByTime indicates an expected call of DeleteEmptyDevicesByTime
func (_mr *MockConnMockRecorder) DeleteEmptyDevicesByTime(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteEmptyDevicesByTime", reflect.TypeOf((*MockConn)(nil).DeleteEmptyDevicesByTime), arg0)
}

// DeleteUser mocks base method
func (_m *MockConn) DeleteUser(_param0 string) error {
	ret := _m.ctrl.Call(_m, "DeleteUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (_mr *MockConnMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteUser", reflect.TypeOf((*MockConn)(nil).DeleteUser), arg0)
}

// GetAdminRoles mocks base method
func (_m *MockConn) GetAdminRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetAdminRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminRoles indicates an expected call of GetAdminRoles
func (_mr *MockConnMockRecorder) GetAdminRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAdminRoles", reflect.TypeOf((*MockConn)(nil).GetAdminRoles))
}

// GetAsset mocks base method
func (_m *MockConn) GetAsset(_param0 string, _param1 *skydb.Asset) error {
	ret := _m.ctrl.Call(_m, "GetAsset", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAsset indicates an expected call of GetAsset
func (_mr *MockConnMockRecorder) GetAsset(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAsset", reflect.TypeOf((*MockConn)(nil).GetAsset), arg0, arg1)
}

// GetAssets mocks base method
func (_m *MockConn) GetAssets(_param0 []string) ([]skydb.Asset, error) {
	ret := _m.ctrl.Call(_m, "GetAssets", _param0)
	ret0, _ := ret[0].([]skydb.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssets indicates an expected call of GetAssets
func (_mr *MockConnMockRecorder) GetAssets(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAssets", reflect.TypeOf((*MockConn)(nil).GetAssets), arg0)
}

// GetDefaultRoles mocks base method
func (_m *MockConn) GetDefaultRoles() ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetDefaultRoles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultRoles indicates an expected call of GetDefaultRoles
func (_mr *MockConnMockRecorder) GetDefaultRoles() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDefaultRoles", reflect.TypeOf((*MockConn)(nil).GetDefaultRoles))
}

// GetDevice mocks base method
func (_m *MockConn) GetDevice(_param0 string, _param1 *skydb.Device) error {
	ret := _m.ctrl.Call(_m, "GetDevice", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetDevice indicates an expected call of GetDevice
func (_mr *MockConnMockRecorder) GetDevice(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDevice", reflect.TypeOf((*MockConn)(nil).GetDevice), arg0, arg1)
}

// GetRecordAccess mocks base method
func (_m *MockConn) GetRecordAccess(_param0 string) (skydb.RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordAccess", _param0)
	ret0, _ := ret[0].(skydb.RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordAccess indicates an expected call of GetRecordAccess
func (_mr *MockConnMockRecorder) GetRecordAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordAccess", reflect.TypeOf((*MockConn)(nil).GetRecordAccess), arg0)
}

// GetRecordDefaultAccess mocks base method
func (_m *MockConn) GetRecordDefaultAccess(_param0 string) (skydb.RecordACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordDefaultAccess", _param0)
	ret0, _ := ret[0].(skydb.RecordACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordDefaultAccess indicates an expected call of GetRecordDefaultAccess
func (_mr *MockConnMockRecorder) GetRecordDefaultAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordDefaultAccess", reflect.TypeOf((*MockConn)(nil).GetRecordDefaultAccess), arg0)
}

// GetRecordFieldAccess mocks base method
func (_m *MockConn) GetRecordFieldAccess() (skydb.FieldACL, error) {
	ret := _m.ctrl.Call(_m, "GetRecordFieldAccess")
	ret0, _ := ret[0].(skydb.FieldACL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordFieldAccess indicates an expected call of GetRecordFieldAccess
func (_mr *MockConnMockRecorder) GetRecordFieldAccess() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRecordFieldAccess", reflect.TypeOf((*MockConn)(nil).GetRecordFieldAccess))
}

// GetUser mocks base method
func (_m *MockConn) GetUser(_param0 string, _param1 *skydb.AuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetUser", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (_mr *MockConnMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUser", reflect.TypeOf((*MockConn)(nil).GetUser), arg0, arg1)
}

// GetUserByPrincipalID mocks base method
func (_m *MockConn) GetUserByPrincipalID(_param0 string, _param1 *skydb.AuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetUserByPrincipalID", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserByPrincipalID indicates an expected call of GetUserByPrincipalID
func (_mr *MockConnMockRecorder) GetUserByPrincipalID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByPrincipalID", reflect.TypeOf((*MockConn)(nil).GetUserByPrincipalID), arg0, arg1)
}

// GetUserByUsernameEmail mocks base method
func (_m *MockConn) GetUserByUsernameEmail(_param0 string, _param1 string, _param2 *skydb.AuthInfo) error {
	ret := _m.ctrl.Call(_m, "GetUserByUsernameEmail", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetUserByUsernameEmail indicates an expected call of GetUserByUsernameEmail
func (_mr *MockConnMockRecorder) GetUserByUsernameEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByUsernameEmail", reflect.TypeOf((*MockConn)(nil).GetUserByUsernameEmail), arg0, arg1, arg2)
}

// PrivateDB mocks base method
func (_m *MockConn) PrivateDB(_param0 string) skydb.Database {
	ret := _m.ctrl.Call(_m, "PrivateDB", _param0)
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

// PrivateDB indicates an expected call of PrivateDB
func (_mr *MockConnMockRecorder) PrivateDB(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PrivateDB", reflect.TypeOf((*MockConn)(nil).PrivateDB), arg0)
}

// PublicDB mocks base method
func (_m *MockConn) PublicDB() skydb.Database {
	ret := _m.ctrl.Call(_m, "PublicDB")
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

// PublicDB indicates an expected call of PublicDB
func (_mr *MockConnMockRecorder) PublicDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "PublicDB", reflect.TypeOf((*MockConn)(nil).PublicDB))
}

// QueryDevicesByUser mocks base method
func (_m *MockConn) QueryDevicesByUser(_param0 string) ([]skydb.Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUser", _param0)
	ret0, _ := ret[0].([]skydb.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDevicesByUser indicates an expected call of QueryDevicesByUser
func (_mr *MockConnMockRecorder) QueryDevicesByUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryDevicesByUser", reflect.TypeOf((*MockConn)(nil).QueryDevicesByUser), arg0)
}

// QueryDevicesByUserAndTopic mocks base method
func (_m *MockConn) QueryDevicesByUserAndTopic(_param0 string, _param1 string) ([]skydb.Device, error) {
	ret := _m.ctrl.Call(_m, "QueryDevicesByUserAndTopic", _param0, _param1)
	ret0, _ := ret[0].([]skydb.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryDevicesByUserAndTopic indicates an expected call of QueryDevicesByUserAndTopic
func (_mr *MockConnMockRecorder) QueryDevicesByUserAndTopic(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryDevicesByUserAndTopic", reflect.TypeOf((*MockConn)(nil).QueryDevicesByUserAndTopic), arg0, arg1)
}

// QueryRelation mocks base method
func (_m *MockConn) QueryRelation(_param0 string, _param1 string, _param2 string, _param3 skydb.QueryConfig) []skydb.AuthInfo {
	ret := _m.ctrl.Call(_m, "QueryRelation", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].([]skydb.AuthInfo)
	return ret0
}

// QueryRelation indicates an expected call of QueryRelation
func (_mr *MockConnMockRecorder) QueryRelation(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryRelation", reflect.TypeOf((*MockConn)(nil).QueryRelation), arg0, arg1, arg2, arg3)
}

// QueryRelationCount mocks base method
func (_m *MockConn) QueryRelationCount(_param0 string, _param1 string, _param2 string) (uint64, error) {
	ret := _m.ctrl.Call(_m, "QueryRelationCount", _param0, _param1, _param2)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRelationCount indicates an expected call of QueryRelationCount
func (_mr *MockConnMockRecorder) QueryRelationCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryRelationCount", reflect.TypeOf((*MockConn)(nil).QueryRelationCount), arg0, arg1, arg2)
}

// QueryUser mocks base method
func (_m *MockConn) QueryUser(_param0 []string, _param1 []string) ([]skydb.AuthInfo, error) {
	ret := _m.ctrl.Call(_m, "QueryUser", _param0, _param1)
	ret0, _ := ret[0].([]skydb.AuthInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUser indicates an expected call of QueryUser
func (_mr *MockConnMockRecorder) QueryUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "QueryUser", reflect.TypeOf((*MockConn)(nil).QueryUser), arg0, arg1)
}

// RemoveRelation mocks base method
func (_m *MockConn) RemoveRelation(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "RemoveRelation", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRelation indicates an expected call of RemoveRelation
func (_mr *MockConnMockRecorder) RemoveRelation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RemoveRelation", reflect.TypeOf((*MockConn)(nil).RemoveRelation), arg0, arg1, arg2)
}

// RevokeRoles mocks base method
func (_m *MockConn) RevokeRoles(_param0 []string, _param1 []string) error {
	ret := _m.ctrl.Call(_m, "RevokeRoles", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRoles indicates an expected call of RevokeRoles
func (_mr *MockConnMockRecorder) RevokeRoles(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RevokeRoles", reflect.TypeOf((*MockConn)(nil).RevokeRoles), arg0, arg1)
}

// SaveAsset mocks base method
func (_m *MockConn) SaveAsset(_param0 *skydb.Asset) error {
	ret := _m.ctrl.Call(_m, "SaveAsset", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAsset indicates an expected call of SaveAsset
func (_mr *MockConnMockRecorder) SaveAsset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SaveAsset", reflect.TypeOf((*MockConn)(nil).SaveAsset), arg0)
}

// SaveDevice mocks base method
func (_m *MockConn) SaveDevice(_param0 *skydb.Device) error {
	ret := _m.ctrl.Call(_m, "SaveDevice", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDevice indicates an expected call of SaveDevice
func (_mr *MockConnMockRecorder) SaveDevice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SaveDevice", reflect.TypeOf((*MockConn)(nil).SaveDevice), arg0)
}

// SetAdminRoles mocks base method
func (_m *MockConn) SetAdminRoles(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetAdminRoles", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAdminRoles indicates an expected call of SetAdminRoles
func (_mr *MockConnMockRecorder) SetAdminRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetAdminRoles", reflect.TypeOf((*MockConn)(nil).SetAdminRoles), arg0)
}

// SetDefaultRoles mocks base method
func (_m *MockConn) SetDefaultRoles(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "SetDefaultRoles", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultRoles indicates an expected call of SetDefaultRoles
func (_mr *MockConnMockRecorder) SetDefaultRoles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetDefaultRoles", reflect.TypeOf((*MockConn)(nil).SetDefaultRoles), arg0)
}

// SetRecordAccess mocks base method
func (_m *MockConn) SetRecordAccess(_param0 string, _param1 skydb.RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordAccess", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordAccess indicates an expected call of SetRecordAccess
func (_mr *MockConnMockRecorder) SetRecordAccess(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordAccess", reflect.TypeOf((*MockConn)(nil).SetRecordAccess), arg0, arg1)
}

// SetRecordDefaultAccess mocks base method
func (_m *MockConn) SetRecordDefaultAccess(_param0 string, _param1 skydb.RecordACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordDefaultAccess", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordDefaultAccess indicates an expected call of SetRecordDefaultAccess
func (_mr *MockConnMockRecorder) SetRecordDefaultAccess(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordDefaultAccess", reflect.TypeOf((*MockConn)(nil).SetRecordDefaultAccess), arg0, arg1)
}

// SetRecordFieldAccess mocks base method
func (_m *MockConn) SetRecordFieldAccess(_param0 skydb.FieldACL) error {
	ret := _m.ctrl.Call(_m, "SetRecordFieldAccess", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecordFieldAccess indicates an expected call of SetRecordFieldAccess
func (_mr *MockConnMockRecorder) SetRecordFieldAccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRecordFieldAccess", reflect.TypeOf((*MockConn)(nil).SetRecordFieldAccess), arg0)
}

// Subscribe mocks base method
func (_m *MockConn) Subscribe(_param0 chan skydb.RecordEvent) error {
	ret := _m.ctrl.Call(_m, "Subscribe", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (_mr *MockConnMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Subscribe", reflect.TypeOf((*MockConn)(nil).Subscribe), arg0)
}

// UnionDB mocks base method
func (_m *MockConn) UnionDB() skydb.Database {
	ret := _m.ctrl.Call(_m, "UnionDB")
	ret0, _ := ret[0].(skydb.Database)
	return ret0
}

// UnionDB indicates an expected call of UnionDB
func (_mr *MockConnMockRecorder) UnionDB() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UnionDB", reflect.TypeOf((*MockConn)(nil).UnionDB))
}

// UpdateUser mocks base method
func (_m *MockConn) UpdateUser(_param0 *skydb.AuthInfo) error {
	ret := _m.ctrl.Call(_m, "UpdateUser", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (_mr *MockConnMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateUser", reflect.TypeOf((*MockConn)(nil).UpdateUser), arg0)
}