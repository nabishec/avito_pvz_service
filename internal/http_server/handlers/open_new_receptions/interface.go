package opennewreceptions

import (
	"context"

	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i PostReceptions
type PostReceptions interface {
	AddReception(ctx context.Context, pvzID uuid.UUID) (*model.ReceptionsResp, error)
}
