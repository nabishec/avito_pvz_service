package login

import (
	"context"

	"github.com/google/uuid"
)

//go:generate minimock -i PostLogin
type PostLogin interface {
	Login(ctx context.Context, email string, password string) (userID uuid.UUID, role string, err error)
}
