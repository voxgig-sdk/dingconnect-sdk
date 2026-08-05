package core

type DingconnectError struct {
	IsDingconnectError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewDingconnectError(code string, msg string, ctx *Context) *DingconnectError {
	return &DingconnectError{
		IsDingconnectError: true,
		Sdk:              "Dingconnect",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *DingconnectError) Error() string {
	return e.Msg
}
