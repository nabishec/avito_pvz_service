package auth

import "github.com/nabishec/avito_pvz_service/internal/model"

//go:generate minimock -i PostAuth
type PostAuth interface {
	CreateUser(email string, password string, role string) (*model.RegisterResp, error)
}
