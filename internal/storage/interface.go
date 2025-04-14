package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

type StorageImp interface {
	AddPVZ(ctx context.Context, city string) (*model.PVZResp, error)
	AddReception(ctx context.Context, pvzID uuid.UUID) (*model.ReceptionsResp, error)
	AddProduct(pvzID uuid.UUID, productType string) (*model.ProductsResp, error)
	DeleteLastProducts(ctx context.Context, pvzID uuid.UUID) error
	CloseLastReceptions(ctx context.Context, pvzID uuid.UUID) error
	CreateUser(ctx context.Context, email string, password string, role string) (*model.RegisterResp, error)
	Login(ctx context.Context, email string, password string) (userID uuid.UUID, role string, err error)
	GetPVZListWithRecep(ctx context.Context, startDate, endDate time.Time, page, limit int) ([]*model.PVZWithRecep, error)
	GetPVZList() ([]*model.PVZResp, error)
	GetValuesForMetrics() (pvzs int, receptions int, products int, err error)
}
