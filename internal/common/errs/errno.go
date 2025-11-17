package errs

var (
	OK    = &ErrNo{Code: 0, Message: "OK"}
	Error = &ErrNo{Code: 1, Message: "ERROR"}

	// enum err msg
	ErrInvalidParam        = &ErrNo{Code: 1001, Message: "Invalid params"}    // invalid params
	ErrBindJson            = &ErrNo{Code: 1002, Message: "Bind json error"}   // err bind json
	ErrRecordNotFound      = &ErrNo{Code: 1003, Message: "Record not found"}  // record not found
	ErrQueryModel          = &ErrNo{Code: 1004, Message: "Query model err."}  // query err
	ErrUpdateRecord        = &ErrNo{Code: 1005, Message: "update model err."} // err update
	ErrRecordAlreadyExists = &ErrNo{Code: 1006, Message: "Already exists."}   // already exists
	ErrRecordyExists       = &ErrNo{Code: 1006, Message: "Already exists."}
	ErrRawsFound           = &ErrNo{Code: 1008, Message: "Raws count err."} // raws count err
	ErrRawsFounds           = &ErrNo{Code: 1008, Message: "Raws count err."} // raws count err
	ErrRawsFoundCount          = &ErrNo{Code: 1008, Message: "Raws count err."} // raws count err
)
