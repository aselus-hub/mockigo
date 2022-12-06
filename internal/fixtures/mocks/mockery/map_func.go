// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type MapFunc struct {
	mock *mock.Mock
}

func NewMapFunc(t mock.Testing) *MapFunc {
	t.Helper()
	return &MapFunc{mock: mock.NewMock(t)}
}

type _MapFunc_Expecter struct {
	mock *mock.Mock
}

func (_mock *MapFunc) EXPECT() _MapFunc_Expecter {
	 return _MapFunc_Expecter{mock: _mock.mock}
}

type _MapFunc_Get_Call struct {
	*mock.Call
}

func (_mock *MapFunc) Get(m map[string]func(string) string) (error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Get", m)
	_r0 := _results.Error(0)
	return _r0
}

func (_expecter _MapFunc_Expecter) Get(m match.Arg[map[string]func(string) string]) _MapFunc_Get_Call {
	return _MapFunc_Get_Call{Call: _expecter.mock.ExpectCall("Get", m.Arg)}
}

func (_call _MapFunc_Get_Call) Return(_r0 error) _MapFunc_Get_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _MapFunc_Get_Call) RunReturn(f func(m map[string]func(string) string) (error)) _MapFunc_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}
