// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jahvon/tuikit/io (interfaces: Logger)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_logger.go -package=mocks . Logger
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	io "github.com/jahvon/tuikit/io"
	gomock "go.uber.org/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Debugx mocks base method.
func (m *MockLogger) Debugx(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugx", varargs...)
}

// Debugx indicates an expected call of Debugx.
func (mr *MockLoggerMockRecorder) Debugx(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugx", reflect.TypeOf((*MockLogger)(nil).Debugx), varargs...)
}

// Error mocks base method.
func (m *MockLogger) Error(arg0 error, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", arg0, arg1)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), arg0, arg1)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Errorx mocks base method.
func (m *MockLogger) Errorx(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorx", varargs...)
}

// Errorx indicates an expected call of Errorx.
func (mr *MockLoggerMockRecorder) Errorx(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorx", reflect.TypeOf((*MockLogger)(nil).Errorx), varargs...)
}

// FatalErr mocks base method.
func (m *MockLogger) FatalErr(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FatalErr", arg0)
}

// FatalErr indicates an expected call of FatalErr.
func (mr *MockLoggerMockRecorder) FatalErr(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FatalErr", reflect.TypeOf((*MockLogger)(nil).FatalErr), arg0)
}

// Fatalf mocks base method.
func (m *MockLogger) Fatalf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockLoggerMockRecorder) Fatalf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// Fatalx mocks base method.
func (m *MockLogger) Fatalx(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalx", varargs...)
}

// Fatalx indicates an expected call of Fatalx.
func (mr *MockLoggerMockRecorder) Fatalx(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalx", reflect.TypeOf((*MockLogger)(nil).Fatalx), varargs...)
}

// Flush mocks base method.
func (m *MockLogger) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockLoggerMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockLogger)(nil).Flush))
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Infox mocks base method.
func (m *MockLogger) Infox(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infox", varargs...)
}

// Infox indicates an expected call of Infox.
func (mr *MockLoggerMockRecorder) Infox(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infox", reflect.TypeOf((*MockLogger)(nil).Infox), varargs...)
}

// LogMode mocks base method.
func (m *MockLogger) LogMode() io.LogMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogMode")
	ret0, _ := ret[0].(io.LogMode)
	return ret0
}

// LogMode indicates an expected call of LogMode.
func (mr *MockLoggerMockRecorder) LogMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogMode", reflect.TypeOf((*MockLogger)(nil).LogMode))
}

// PlainTextDebug mocks base method.
func (m *MockLogger) PlainTextDebug(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PlainTextDebug", arg0)
}

// PlainTextDebug indicates an expected call of PlainTextDebug.
func (mr *MockLoggerMockRecorder) PlainTextDebug(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlainTextDebug", reflect.TypeOf((*MockLogger)(nil).PlainTextDebug), arg0)
}

// PlainTextError mocks base method.
func (m *MockLogger) PlainTextError(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PlainTextError", arg0)
}

// PlainTextError indicates an expected call of PlainTextError.
func (mr *MockLoggerMockRecorder) PlainTextError(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlainTextError", reflect.TypeOf((*MockLogger)(nil).PlainTextError), arg0)
}

// PlainTextInfo mocks base method.
func (m *MockLogger) PlainTextInfo(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PlainTextInfo", arg0)
}

// PlainTextInfo indicates an expected call of PlainTextInfo.
func (mr *MockLoggerMockRecorder) PlainTextInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlainTextInfo", reflect.TypeOf((*MockLogger)(nil).PlainTextInfo), arg0)
}

// PlainTextSuccess mocks base method.
func (m *MockLogger) PlainTextSuccess(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PlainTextSuccess", arg0)
}

// PlainTextSuccess indicates an expected call of PlainTextSuccess.
func (mr *MockLoggerMockRecorder) PlainTextSuccess(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlainTextSuccess", reflect.TypeOf((*MockLogger)(nil).PlainTextSuccess), arg0)
}

// PlainTextWarn mocks base method.
func (m *MockLogger) PlainTextWarn(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PlainTextWarn", arg0)
}

// PlainTextWarn indicates an expected call of PlainTextWarn.
func (mr *MockLoggerMockRecorder) PlainTextWarn(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlainTextWarn", reflect.TypeOf((*MockLogger)(nil).PlainTextWarn), arg0)
}

// Println mocks base method.
func (m *MockLogger) Println(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Println", arg0)
}

// Println indicates an expected call of Println.
func (mr *MockLoggerMockRecorder) Println(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Println", reflect.TypeOf((*MockLogger)(nil).Println), arg0)
}

// SetLevel mocks base method.
func (m *MockLogger) SetLevel(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLevel", arg0)
}

// SetLevel indicates an expected call of SetLevel.
func (mr *MockLoggerMockRecorder) SetLevel(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLevel", reflect.TypeOf((*MockLogger)(nil).SetLevel), arg0)
}

// SetMode mocks base method.
func (m *MockLogger) SetMode(arg0 io.LogMode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMode", arg0)
}

// SetMode indicates an expected call of SetMode.
func (mr *MockLoggerMockRecorder) SetMode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMode", reflect.TypeOf((*MockLogger)(nil).SetMode), arg0)
}

// Warnf mocks base method.
func (m *MockLogger) Warnf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockLoggerMockRecorder) Warnf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}

// Warnx mocks base method.
func (m *MockLogger) Warnx(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnx", varargs...)
}

// Warnx indicates an expected call of Warnx.
func (mr *MockLoggerMockRecorder) Warnx(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnx", reflect.TypeOf((*MockLogger)(nil).Warnx), varargs...)
}
