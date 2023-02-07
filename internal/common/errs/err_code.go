package errs

type ErrNo struct {
	Code    int64
	Message string
}

// Error. error msg
func (err *ErrNo) Error() string {
	return err.Message
}
