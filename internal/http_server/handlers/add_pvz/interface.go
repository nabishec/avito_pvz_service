package addpvz

import (
	"context"

	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i PostPVZ
type PostPVZ interface {
	AddPVZ(ctx context.Context, city string) (*model.PVZResp, error)
}
