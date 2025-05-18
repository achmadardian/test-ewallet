package errs

import "errors"

var (
	ErrValidationError  = errors.New("validation error")
	ErrPhoneAlreadyUsed = errors.New("nomor telepon sudah terdaftar")
)
