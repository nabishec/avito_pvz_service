package main

import (
	"net"
	"os"

	grpcgetpvz "github.com/nabishec/avito_pvz_service/internal/grpc_server/handler/grpc_get_pvz"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/nabishec/avito_pvz_service/pvz/pvz_v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type grpcServer struct {
	storage storage.StorageImp
}

func NewGRPCServer(storage storage.StorageImp) *grpcServer {
	return &grpcServer{
		storage: storage,
	}
}

func (s *grpcServer) Run() error {

	addr := ":" + os.Getenv("GRPC_SERVER_PORT")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error().Err(err).Msg("failed to run grpc")
		return err
	}
	grpcServer := grpc.NewServer()
	pvz_v1.RegisterPVZServiceServer(grpcServer, grpcgetpvz.NewPVZ(s.storage))
	log.Info().Msgf("Start grpc server on %s", addr)

	return grpcServer.Serve(lis)

}
