package errs

import "errors"

// ErrDBInit 初始化错误
var ErrDBInit = errors.New("ErrDBInit")

//MError defined error
type MError struct {
	Code int
	Msg  string
}

func (e *MError) Error() string {
	return e.Msg
}
