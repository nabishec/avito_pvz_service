package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"

	_ "github.com/nabishec/avito_pvz_service/docs"
	dbconnection "github.com/nabishec/avito_pvz_service/internal/app/db_connection"
	grpcserver "github.com/nabishec/avito_pvz_service/internal/app/grpc_server"
	httpserver "github.com/nabishec/avito_pvz_service/internal/app/http_server"
	"github.com/nabishec/avito_pvz_service/internal/metrics"
	"github.com/nabishec/avito_pvz_service/internal/storage/db"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title API Avito PVZ Service
// @version 1.0.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("d", false, "set log level to debug")

	easyReading := flag.Bool("r", false, "set console writer")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// for easy reading
	if *easyReading {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	err := LoadEnv()
	if err != nil {
		log.Error().Err(err).Msg("Configuration not found")
		os.Exit(1)
	}

	// init storage postgresql
	log.Info().Msg("Init storage")
	dbConnection, err := dbconnection.NewDatabaseConnection()
	if err != nil {
		log.Error().Err(err).Msg("Failed init database")
		os.Exit(1)
	}
	log.Info().Msg("Database init successful")

	storage := db.NewStorage(dbConnection.DB)

	grpcServer := grpcserver.NewGRPCServer(storage)
	httpServer := httpserver.NewHTTPServer(storage)
	serverPrometheus := httpserver.NewHTTPServer(storage)

	//  run server
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := grpcServer.Run()
		if err != nil {
			log.Error().Err(err).Msg("grpc server stopped")
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		serverPrometheus.Router.Handle("/metrics", promhttp.Handler())
		metricsRecorder := metrics.NewMetrics(storage)
		go metricsRecorder.CheckMetricsFromDBCircle()

		addr := ":" + os.Getenv("INTERNAL_PROMETHEUS_PORT")
		log.Info().Msgf("Starting prometheus connection to app on %s", addr)
		err := http.ListenAndServe(addr, serverPrometheus.Router)
		if err != nil {
			log.Error().Err(err).Msg("prometheus connection stopped")
			return
		}
	}()

	err = httpServer.Run()
	if err != nil {
		log.Error().Err(err).Msg("http server stop")
	}

	wg.Wait()
	log.Error().Msg("Program ended")
}

func LoadEnv() error {
	const op = "cmd.loadEnv()"
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("%s:%s", op, "failed load env file")
	}
	return nil
}
