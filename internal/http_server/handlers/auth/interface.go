package auth

import (
	"context"

	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i PostAuth
type PostAuth interface {
	CreateUser(ctx context.Context, email string, password string, role string) (*model.RegisterResp, error)
}
