package closelastreceptions

import "github.com/google/uuid"

//go:generate minimock -i PostCloseLastReceptions
type PostCloseLastReceptions interface {
	CloseLastReceptions(pvzID uuid.UUID) error
}
