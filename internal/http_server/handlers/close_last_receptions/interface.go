package closelastreceptions

import "github.com/google/uuid"

type PostCloseLastReceptions interface {
	CloseLastReceptions(pvzID uuid.UUID) error
}
