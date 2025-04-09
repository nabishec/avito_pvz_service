package model

import (
	"time"

	"github.com/google/uuid"
)

type SuccResp struct {
	Status string `json:"status"`
}

func ReturnSuccResp(status string) SuccResp {
	return SuccResp{
		Status: status,
	}
}

type ErrorResponse struct {
	Error string `json:"errors"`
}

func ReturnErrResp(errMsg string) ErrorResponse {
	return ErrorResponse{
		Error: errMsg,
	}
}

type RootTokenReq struct {
	UserRole string `json:"role" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type PVZReq struct {
	City string `json:"city" validate:"required"`
}

type PVZResp struct {
	ID               uuid.UUID `json:"id" db:"id"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	City             string    `json:"city" validate:"required" db:"city"`
}

type ReceptionsResp struct {
	ID       uuid.UUID `json:"id" db:"id"`
	DateTime time.Time `json:"date_time" db:"registration_date"`
	PVZID    string    `json:"pvz_id" db:"pvz_id"`
	Status   string    `json:"status" db:"status"`
}

type ReceptionsReq struct {
	PVZID string `json:"pvz_id" validate:"required"`
}

type ProductsReq struct {
	Type  string `json:"type" validate:"required"`
	PVZID string `json:"pvz_id" validate:"required"`
}

type ProductsResp struct {
	ID          uuid.UUID `json:"id" db:"id"`
	DateTime    time.Time `json:"date_time" db:"registration_date"`
	Type        string    `json:"type" db:"type"`
	ReceptionID string    `json:"reception_id" db:"reception_id"`
}

type RegisterReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type RegisterResp struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
	Role  string    `json:"role" db:"role"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
