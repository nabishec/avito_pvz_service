package grpcgetpvz

import "github.com/nabishec/avito_pvz_service/internal/model"

type GetPVZ interface {
	GetPVZList() ([]*model.PVZResp, error)
}
