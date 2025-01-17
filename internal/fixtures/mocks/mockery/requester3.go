// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type Requester3 struct {
	mock *mock.Mock
}

func NewRequester3(t mock.Testing) *Requester3 {
	t.Helper()
	return &Requester3{mock: mock.NewMock(t)}
}

type _Requester3_Expecter struct {
	mock *mock.Mock
}

func (_mock *Requester3) EXPECT() _Requester3_Expecter {
	 return _Requester3_Expecter{mock: _mock.mock}
}

type _Requester3_Get_Call struct {
	*mock.Call
}

func (_mock *Requester3) Get() (error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Get", )
	_r0 := _results.Error(0)
	return _r0
}

func (_expecter _Requester3_Expecter) Get() _Requester3_Get_Call {
	return _Requester3_Get_Call{Call: _expecter.mock.ExpectCall("Get", )}
}

func (_call _Requester3_Get_Call) Return(_r0 error) _Requester3_Get_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _Requester3_Get_Call) RunReturn(f func() (error)) _Requester3_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}
