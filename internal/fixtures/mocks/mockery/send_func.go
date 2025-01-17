// Code generated by mockigo. DO NOT EDIT.

package mockery

import context "context"
import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type SendFunc struct {
	mock *mock.Mock
}

func NewSendFunc(t mock.Testing) *SendFunc {
	t.Helper()
	return &SendFunc{mock: mock.NewMock(t)}
}

type _SendFunc_Expecter struct {
	mock *mock.Mock
}

func (_mock *SendFunc) EXPECT() _SendFunc_Expecter {
	 return _SendFunc_Expecter{mock: _mock.mock}
}

type _SendFunc_Execute_Call struct {
	*mock.Call
}

func (_mock *SendFunc) Execute(ctx context.Context, data string) (int, error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Execute", ctx, data)
	_r0 := _results.Get(0).(int)
	_r1 := _results.Error(1)
	return _r0, _r1
}

func (_expecter _SendFunc_Expecter) Execute(ctx match.Arg[context.Context], data match.Arg[string]) _SendFunc_Execute_Call {
	return _SendFunc_Execute_Call{Call: _expecter.mock.ExpectCall("Execute", ctx.Arg, data.Arg)}
}

func (_call _SendFunc_Execute_Call) Return(_r0 int, _r1 error) _SendFunc_Execute_Call {
	_call.Call.Return(_r0, _r1)
	return _call
}

func (_call _SendFunc_Execute_Call) RunReturn(f func(ctx context.Context, data string) (int, error)) _SendFunc_Execute_Call {
	_call.Call.RunReturn(f)
	return _call
}
