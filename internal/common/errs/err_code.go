package errs

type ErrNo struct {
	Code    int64
	Message string
}

func (err *ErrNo) Error() string {
	return err.Message
}
