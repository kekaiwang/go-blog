package errs

var (
	OK    = &ErrNo{Code: 0, Message: "OK"}
	Error = &ErrNo{Code: 1, Message: "ERROR"}

	// enum err msg
	ErrInvalidParam        = &ErrNo{Code: 1001, Message: "Invalid params"}
	ErrBindJson            = &ErrNo{Code: 1002, Message: "Bind json error"}
	ErrRecordNotFound      = &ErrNo{Code: 1003, Message: "Record not found"}
	ErrQueryModel          = &ErrNo{Code: 1004, Message: "Query model err."}
	ErrUpdateRecord        = &ErrNo{Code: 1005, Message: "update model err."}
	ErrRecordAlreadyExists = &ErrNo{Code: 1006, Message: "Already exists."}
	ErrRecordyExists       = &ErrNo{Code: 1006, Message: "Already exists."}
)
