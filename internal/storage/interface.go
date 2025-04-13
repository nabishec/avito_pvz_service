package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

type StorageImp interface {
	AddPVZ(city string) (*model.PVZResp, error)
	AddReception(pvzID uuid.UUID) (*model.ReceptionsResp, error)
	AddProduct(pvzID uuid.UUID, productType string) (*model.ProductsResp, error)
	DeleteLastProducts(pvzID uuid.UUID) error
	CloseLastReceptions(pvzID uuid.UUID) error
	CreateUser(email string, password string, role string) (*model.RegisterResp, error)
	Login(email string, password string) (uuid.UUID, string, error)
	GetPVZListWithRecep(startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
	GetPVZList() ([]*model.PVZResp, error)
	GetValuesForMetrics() (pvzs int, receptions int, products int, err error)
}
