package addproducts

import (
	"context"

	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

//go:generate minimock -i PostProducts
type PostProducts interface {
	AddProduct(ctx context.Context, pvzID uuid.UUID, productType string) (*model.ProductsResp, error)
}
