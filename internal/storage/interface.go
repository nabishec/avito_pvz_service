package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

type StorageImp interface {
	AddPVZ(city string) (*model.PVZResp, error)
	AddReception(ctx context.Context, pvzID uuid.UUID) (*model.ReceptionsResp, error)
	AddProduct(pvzID uuid.UUID, productType string) (*model.ProductsResp, error)
	DeleteLastProducts(pvzID uuid.UUID) error
	CloseLastReceptions(pvzID uuid.UUID) error
	CreateUser(email string, password string, role string) (*model.RegisterResp, error)
	Login(ctx context.Context, email string, password string) (userID uuid.UUID, role string, err error)
	GetPVZListWithRecep(startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
	GetPVZList() ([]*model.PVZResp, error)
	GetValuesForMetrics() (pvzs int, receptions int, products int, err error)
}
