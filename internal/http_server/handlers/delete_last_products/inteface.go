package deletelastproducts

import "github.com/google/uuid"

type PostDeleteLastProducts interface {
	DeleteLastProducts(pvzID uuid.UUID) error
}
