package getpvzlist

import (
	"context"
	"time"

	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i GetPVZ
type GetPVZ interface {
	GetPVZListWithRecep(ctx context.Context, startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
}
