package errors

import (
	"fmt"
)

// CustomError 自定义错误
type CustomError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *CustomError {
	return &CustomError{
		Code: code,
		Msg:  msg,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code %d err %s", e.Code, e.Msg)
}
