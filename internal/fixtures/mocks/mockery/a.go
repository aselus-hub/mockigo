// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"
import mockery "github.com/subtle-byte/mockigo/internal/fixtures/mockery"

var _ = match.Any[int]

type A struct {
	mock *mock.Mock
}

func NewA(t mock.Testing) *A {
	t.Helper()
	return &A{mock: mock.NewMock(t)}
}

type _A_Expecter struct {
	mock *mock.Mock
}

func (_mock *A) EXPECT() _A_Expecter {
	 return _A_Expecter{mock: _mock.mock}
}

type _A_Call_Call struct {
	*mock.Call
}

func (_mock *A) Call() (mockery.B, error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Call", )
	_r0 := _results.Get(0).(mockery.B)
	_r1 := _results.Error(1)
	return _r0, _r1
}

func (_expecter _A_Expecter) Call() _A_Call_Call {
	return _A_Call_Call{Call: _expecter.mock.ExpectCall("Call", )}
}

func (_call _A_Call_Call) Return(_r0 mockery.B, _r1 error) _A_Call_Call {
	_call.Call.Return(_r0, _r1)
	return _call
}

func (_call _A_Call_Call) RunReturn(f func() (mockery.B, error)) _A_Call_Call {
	_call.Call.RunReturn(f)
	return _call
}
