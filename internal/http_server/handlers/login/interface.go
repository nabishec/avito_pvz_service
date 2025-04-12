package login

import "github.com/google/uuid"

//go:generate minimock -i PostLogin
type PostLogin interface {
	Login(email string, password string) (userID uuid.UUID, role string, err error)
}
