package grpcgetpvz

import (
	"context"

	"github.com/nabishec/avito_pvz_service/pvz/pvz_v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PVZ struct {
	GetPVZ GetPVZ
	pvz_v1.UnimplementedPVZServiceServer
}

func NewPVZ(getPVZ GetPVZ) *PVZ {
	return &PVZ{
		GetPVZ: getPVZ,
	}
}

func (h *PVZ) GetPVZList(ctx context.Context, req *pvz_v1.GetPVZListRequest) (*pvz_v1.GetPVZListResponse, error) {
	const op = "internal.grpc_server.handlers.grpcgetpvz.GetPVZList()"
	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to get pvz list from grpc has been received")

	pvzList, err := h.GetPVZ.GetPVZList()
	if err != nil {
		return nil, err
	}

	responceList := make([]*pvz_v1.PVZ, len(pvzList))
	for i, pvz := range pvzList {
		responceList[i] = &pvz_v1.PVZ{
			Id:               pvz.ID.String(),
			City:             pvz.City,
			RegistrationDate: timestamppb.New(pvz.RegistrationDate),
		}
	}

	return &pvz_v1.GetPVZListResponse{
		Pvzs: responceList,
	}, nil
}
