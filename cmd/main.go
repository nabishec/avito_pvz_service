package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	dbconnection "github.com/nabishec/avito_pvz_service/cmd/db_connection"
	_ "github.com/nabishec/avito_pvz_service/docs"
	addproducts "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/add_products"
	addpvz "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/add_pvz"
	"github.com/nabishec/avito_pvz_service/internal/http_server/handlers/auth"
	closelastreceptions "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/close_last_receptions"
	deletelastproducts "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/delete_last_products"
	getpvzlist "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/get_pvz_list"
	"github.com/nabishec/avito_pvz_service/internal/http_server/handlers/login"
	opennewreceptions "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/open_new_receptions"
	roottoken "github.com/nabishec/avito_pvz_service/internal/http_server/handlers/root_token"
	custommiddleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/nabishec/avito_pvz_service/internal/storage/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
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
	//for easy reading
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

	s := CreateNewServer(storage)
	s.Router.Use(middleware.Heartbeat("/"))
	s.MountHandlers()

	//  run server
	wrTime, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil || wrTime == 0 {
		log.Warn().Err(err).Msg("timeout not received from env")
		wrTime = 4 * time.Second
	}
	idleTime, err := time.ParseDuration(os.Getenv("IDLE_TIMEOUT"))
	if err != nil || idleTime == 0 {
		log.Warn().Err(err).Msg("idle timeout not received from env")
		idleTime = 60 * time.Second
	}

	srv := &http.Server{
		Addr:         "localhost:" + os.Getenv("SERVER_PORT"),
		Handler:      s.Router,
		ReadTimeout:  wrTime,
		WriteTimeout: wrTime,
		IdleTimeout:  idleTime,
	}
	log.Info().Msgf("Starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("failed to start server")
		os.Exit(1)
	}

	log.Error().Msg("Program ended")
}

func (s *Server) MountHandlers() {
	rootToken := roottoken.NewRootToken()
	pvz := addpvz.NewPVZ(s.Storage)
	receptions := opennewreceptions.NewReceptions(s.Storage)
	products := addproducts.NewProducts(s.Storage)
	deleteLastProduct := deletelastproducts.NewDeleteProducts(s.Storage)
	closeLastReceptions := closelastreceptions.NewCloseLastReceptions(s.Storage)
	auth := auth.NewAuth(s.Storage)
	login := login.NewLogin(s.Storage)
	pvzList := getpvzlist.NewPVZ(s.Storage)

	s.Router.Group(func(r chi.Router) {
		r.Post("/dummyLogin", rootToken.ReturnRootToken)
		r.Get("/swagger/*", httpSwagger.WrapHandler)
		r.Post("/register", auth.Register)
		r.Post("/login", login.Login)
	})

	// Require Authentication
	s.Router.Group(func(r chi.Router) {
		r.Use(custommiddleware.Auth)
		r.Post("/pvz", pvz.AddPVZ)
		r.Post("/receptions", receptions.AddReceptions)
		r.Post("/products", products.AddProducts)
		r.Post("/pvz/{pvzId}/delete_last_product", deleteLastProduct.DeleteProducts)
		r.Post("/pvz/{pvzId}/close_last_reception", closeLastReceptions.CloseLastReceptions)
		r.Get("/pvz", pvzList.GetPVZList)
	})

}

type Server struct {
	Router  *chi.Mux
	Storage storage.StorageImp
}

func CreateNewServer(storage storage.StorageImp) *Server {
	s := &Server{
		Router:  chi.NewRouter(),
		Storage: storage,
	}
	return s
}

func LoadEnv() error {
	const op = "cmd.loadEnv()"
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("%s:%s", op, "failed load env file")
	}
	return nil
}
