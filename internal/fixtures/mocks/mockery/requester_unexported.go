// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type Requester_unexported struct {
	mock *mock.Mock
}

func NewRequester_unexported(t mock.Testing) *Requester_unexported {
	t.Helper()
	return &Requester_unexported{mock: mock.NewMock(t)}
}

type _Requester_unexported_Expecter struct {
	mock *mock.Mock
}

func (_mock *Requester_unexported) EXPECT() _Requester_unexported_Expecter {
	 return _Requester_unexported_Expecter{mock: _mock.mock}
}

type _Requester_unexported_Get_Call struct {
	*mock.Call
}

func (_mock *Requester_unexported) Get() () {
	_mock.mock.T.Helper()
	_mock.mock.Called("Get", )
}

func (_expecter _Requester_unexported_Expecter) Get() _Requester_unexported_Get_Call {
	return _Requester_unexported_Get_Call{Call: _expecter.mock.ExpectCall("Get", )}
}

func (_call _Requester_unexported_Get_Call) Return() _Requester_unexported_Get_Call {
	_call.Call.Return()
	return _call
}

func (_call _Requester_unexported_Get_Call) RunReturn(f func() ()) _Requester_unexported_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}
