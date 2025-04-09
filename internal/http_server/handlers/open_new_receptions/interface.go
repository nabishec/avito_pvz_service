package opennewreceptions

import (
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

type PostReceptions interface {
	AddReception(pvzID uuid.UUID) (*model.ReceptionsResp, error)
}
