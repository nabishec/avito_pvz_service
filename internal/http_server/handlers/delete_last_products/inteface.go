package deletelastproducts

import (
	"context"

	"github.com/google/uuid"
)

//go:generate minimock -i PostDeleteLastProducts
type PostDeleteLastProducts interface {
	DeleteLastProducts(ctx context.Context, pvzID uuid.UUID) error
}
