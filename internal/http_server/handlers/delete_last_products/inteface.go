package deletelastproducts

import "github.com/google/uuid"

//go:generate minimock -i PostDeleteLastProducts
type PostDeleteLastProducts interface {
	DeleteLastProducts(pvzID uuid.UUID) error
}
