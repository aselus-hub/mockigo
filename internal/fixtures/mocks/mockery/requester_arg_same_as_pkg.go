// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type RequesterArgSameAsPkg struct {
	mock *mock.Mock
}

func NewRequesterArgSameAsPkg(t mock.Testing) *RequesterArgSameAsPkg {
	t.Helper()
	return &RequesterArgSameAsPkg{mock: mock.NewMock(t)}
}

type _RequesterArgSameAsPkg_Expecter struct {
	mock *mock.Mock
}

func (_mock *RequesterArgSameAsPkg) EXPECT() _RequesterArgSameAsPkg_Expecter {
	 return _RequesterArgSameAsPkg_Expecter{mock: _mock.mock}
}

type _RequesterArgSameAsPkg_Get_Call struct {
	*mock.Call
}

func (_mock *RequesterArgSameAsPkg) Get(test string) () {
	_mock.mock.T.Helper()
	_mock.mock.Called("Get", test)
}

func (_expecter _RequesterArgSameAsPkg_Expecter) Get(test match.Arg[string]) _RequesterArgSameAsPkg_Get_Call {
	return _RequesterArgSameAsPkg_Get_Call{Call: _expecter.mock.ExpectCall("Get", test.Arg)}
}

func (_call _RequesterArgSameAsPkg_Get_Call) Return() _RequesterArgSameAsPkg_Get_Call {
	_call.Call.Return()
	return _call
}

func (_call _RequesterArgSameAsPkg_Get_Call) RunReturn(f func(test string) ()) _RequesterArgSameAsPkg_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}
