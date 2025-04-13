package getpvzlist

import (
	"time"

	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i GetPVZ
type GetPVZ interface {
	GetPVZListWithRecep(startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
}
