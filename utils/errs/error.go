package errs

import "errors"

var (
	ErrValidationError  = errors.New("validation error")
	ErrPhoneAlreadyUsed = errors.New("nomor telepon sudah terdaftar")
	ErrDataNotFound     = errors.New("data tidak ditemukan")
	ErrInvalidLogin     = errors.New("nomor telepon atau pin tidak cocok")
	ErrInvalidToken     = errors.New("invalid token")
	ErrInsufficientFund = errors.New("saldo tidak cukup")
)
