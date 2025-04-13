package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
)

type httpServer struct {
	Router  *chi.Mux
	Storage storage.StorageImp
}

func NewHTTPServer(storage storage.StorageImp) *httpServer {
	s := &httpServer{
		Router:  chi.NewRouter(),
		Storage: storage,
	}
	return s
}

func (s *httpServer) MountHandlers() {
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

func (s *httpServer) Run() error {
	s.Router.Use(middleware.Heartbeat("/"))
	s.MountHandlers()

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

	addr := "localhost:" + os.Getenv("HTTP_SERVER_PORT")

	srv := &http.Server{
		Addr:         addr,
		Handler:      s.Router,
		ReadTimeout:  wrTime,
		WriteTimeout: wrTime,
		IdleTimeout:  idleTime,
	}
	log.Info().Msgf("Starting http server on %s", srv.Addr)

	return srv.ListenAndServe()
}
