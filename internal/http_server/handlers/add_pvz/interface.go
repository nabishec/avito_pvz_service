package addpvz

import (
	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i PostPVZ
type PostPVZ interface {
	AddPVZ(city string) (*model.PVZResp, error)
}
