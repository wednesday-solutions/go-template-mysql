// Code generated by MockGen. DO NOT EDIT.	
// Source: github.com/go-playground/validator (interfaces: FieldError)	

// Package mock_validator is a generated GoMock package.	
package testutls	

import (	
	reflect "reflect"	

	universal_translator "github.com/go-playground/universal-translator"	
	gomock "github.com/golang/mock/gomock"	
)	

// MockFieldError is a mock of FieldError interface.	
type MockFieldError struct {	
	ctrl     *gomock.Controller	
	recorder *MockFieldErrorMockRecorder	
}	

// MockFieldErrorMockRecorder is the mock recorder for MockFieldError.	
type MockFieldErrorMockRecorder struct {	
	mock *MockFieldError	
}	

// NewMockFieldError creates a new mock instance.	
func NewMockFieldError(ctrl *gomock.Controller) *MockFieldError {	
	mock := &MockFieldError{ctrl: ctrl}	
	mock.recorder = &MockFieldErrorMockRecorder{mock}	
	return mock	
}	

// EXPECT returns an object that allows the caller to indicate expected use.	
func (m *MockFieldError) EXPECT() *MockFieldErrorMockRecorder {	
	return m.recorder	
}	

// ActualTag mocks base method.	
func (m *MockFieldError) ActualTag() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "ActualTag")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// ActualTag indicates an expected call of ActualTag.	
func (mr *MockFieldErrorMockRecorder) ActualTag() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActualTag", reflect.TypeOf((*MockFieldError)(nil).ActualTag))	
}	

// Field mocks base method.	
func (m *MockFieldError) Field() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Field")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// Field indicates an expected call of Field.	
func (mr *MockFieldErrorMockRecorder) Field() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Field", reflect.TypeOf((*MockFieldError)(nil).Field))	
}	

// Kind mocks base method.	
func (m *MockFieldError) Kind() reflect.Kind {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Kind")	
	ret0, _ := ret[0].(reflect.Kind)	
	return ret0	
}	

// Kind indicates an expected call of Kind.	
func (mr *MockFieldErrorMockRecorder) Kind() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockFieldError)(nil).Kind))	
}	

// Namespace mocks base method.	
func (m *MockFieldError) Namespace() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Namespace")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// Namespace indicates an expected call of Namespace.	
func (mr *MockFieldErrorMockRecorder) Namespace() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockFieldError)(nil).Namespace))	
}	

// Param mocks base method.	
func (m *MockFieldError) Param() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Param")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// Param indicates an expected call of Param.	
func (mr *MockFieldErrorMockRecorder) Param() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockFieldError)(nil).Param))	
}	

// StructField mocks base method.	
func (m *MockFieldError) StructField() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "StructField")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// StructField indicates an expected call of StructField.	
func (mr *MockFieldErrorMockRecorder) StructField() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructField", reflect.TypeOf((*MockFieldError)(nil).StructField))	
}	

// StructNamespace mocks base method.	
func (m *MockFieldError) StructNamespace() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "StructNamespace")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// StructNamespace indicates an expected call of StructNamespace.	
func (mr *MockFieldErrorMockRecorder) StructNamespace() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructNamespace", reflect.TypeOf((*MockFieldError)(nil).StructNamespace))	
}	

// Tag mocks base method.	
func (m *MockFieldError) Tag() string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Tag")	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// Tag indicates an expected call of Tag.	
func (mr *MockFieldErrorMockRecorder) Tag() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockFieldError)(nil).Tag))	
}	

// Translate mocks base method.	
func (m *MockFieldError) Translate(arg0 universal_translator.Translator) string {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Translate", arg0)	
	ret0, _ := ret[0].(string)	
	return ret0	
}	

// Translate indicates an expected call of Translate.	
func (mr *MockFieldErrorMockRecorder) Translate(arg0 interface{}) *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockFieldError)(nil).Translate), arg0)	
}	

// Type mocks base method.	
func (m *MockFieldError) Type() reflect.Type {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Type")	
	ret0, _ := ret[0].(reflect.Type)	
	return ret0	
}	

// Type indicates an expected call of Type.	
func (mr *MockFieldErrorMockRecorder) Type() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockFieldError)(nil).Type))	
}	

// Value mocks base method.	
func (m *MockFieldError) Value() interface{} {	
	m.ctrl.T.Helper()	
	ret := m.ctrl.Call(m, "Value")	
	ret0, _ := ret[0].(interface{})	
	return ret0	
}	

// Value indicates an expected call of Value.	
func (mr *MockFieldErrorMockRecorder) Value() *gomock.Call {	
	mr.mock.ctrl.T.Helper()	
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockFieldError)(nil).Value))	
}	
