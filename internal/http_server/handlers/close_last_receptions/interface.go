package closelastreceptions

import (
	"context"

	"github.com/google/uuid"
)

//go:generate minimock -i PostCloseLastReceptions
type PostCloseLastReceptions interface {
	CloseLastReceptions(ctx context.Context, pvzID uuid.UUID) error
}
