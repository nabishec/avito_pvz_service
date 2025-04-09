package addproducts

import (
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

type PostProducts interface {
	AddProduct(pvzID uuid.UUID, productType string) (*model.ProductsResp, error)
}
