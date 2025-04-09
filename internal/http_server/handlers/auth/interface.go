package auth

import "github.com/nabishec/avito_pvz_service/internal/model"

type PostAuth interface {
	CreateUser(email string, password string, role string) (*model.RegisterResp, error)
}
