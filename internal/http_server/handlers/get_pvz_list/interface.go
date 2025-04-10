package getpvzlist

import (
	"time"

	"github.com/nabishec/avito_pvz_service/internal/model"
)

type GetPVZ interface {
	GetPVZList(startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
}
