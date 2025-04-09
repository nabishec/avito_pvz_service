package storage

import "errors"

var (
	ErrPVZNotExist                 = errors.New("pvz not exist")
	ErrPreviousReceptionNotClosed  = errors.New("previous reception not closed")
	ErrOpenReceptionNotExist       = errors.New("open reception not exist")
	ErrProductsInReceptionNotExist = errors.New("products in reception not exist")
	ErrPasswordIsEmpty             = errors.New("password is empty")
	ErrUserAlreadyExist            = errors.New("user already exist")
	ErrUserNotExist                = errors.New("user not exist")
	ErrPasswordIsWrong             = errors.New("password is wrong")
)
