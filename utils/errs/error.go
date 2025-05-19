package errs

import "errors"

var (
	ErrValidationError  = errors.New("validation error")
	ErrPhoneAlreadyUsed = errors.New("nomor telepon sudah terdaftar")
	ErrDataNotFound     = errors.New("data tidak ditemukan")
	ErrInvalidLogin     = errors.New("nomor telepon atau pin tidak cocok")
)
