package errors

type CustomError struct {
	Err int
	Msg string
}

//var (
//	LOGIN_UNKNOWN = New(202, "用户不存在")
//)

func (e *CustomError) Error() string {
	return e.Msg
}

func New(err int, msg string) *CustomError {
	return &CustomError{
		Err: err,
		Msg: msg,
	}
}
