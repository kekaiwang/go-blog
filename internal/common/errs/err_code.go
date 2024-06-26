package errs

// ErrNo. err no struct
type ErrNo struct {
	Code    int64  // err code
	Message string // err msg
}

// Error. error msg
func (err *ErrNo) Error() string {
	return err.Message // err msg
}
