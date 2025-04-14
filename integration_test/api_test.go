package integrationtest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	dbconnection "github.com/nabishec/avito_pvz_service/internal/app/db_connection"
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
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/nabishec/avito_pvz_service/internal/storage/db"
)

func executeRequest(req *http.Request, s *httpServer) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func testGetTokenForClient(t *testing.T, s *httpServer, role string) string {
	reqBody := map[string]string{
		"role": role,
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	req, _ := http.NewRequest("POST", "/dummyLogin", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)

	var authResp model.AuthResponse
	err = json.Unmarshal(response.Body.Bytes(), &authResp)
	if err != nil {
		t.Fatal("failed to unmarshal response body")
	}
	return authResp.Token
}

func TestService(t *testing.T) {
	err := loadEnv()
	if err != nil {
		t.Error("Don't found config")
	}

	dbConnection, err := dbconnection.NewDatabaseConnection()
	if err != nil {
		t.Error("Failed init database")
	}

	storage := db.NewStorage(dbConnection.DB)

	s := newHTTPServer(storage)
	s.mountHandlers()
	tokenClient := testGetTokenForClient(t, s, "client")
	tokenModerator := testGetTokenForClient(t, s, "moderator")
	var pvzResp model.PVZResp

	t.Run("Add PVZ", func(t *testing.T) {
		reqBody := map[string]string{
			"city": "Санкт-Петербург",
		}

		jsonReq, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer(jsonReq))
		req.Header.Set("Authorization", "Bearer "+tokenModerator)
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req, s)
		checkResponseCode(t, http.StatusCreated, response.Code)

		err = json.Unmarshal(response.Body.Bytes(), &pvzResp)
		if err != nil {
			t.Fatal("failed to unmarshal response body")
		}
	})

	t.Run("Add Reception", func(t *testing.T) {
		reqBody := map[string]string{
			"pvzId": pvzResp.ID.String(),
		}

		jsonReq, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req, _ := http.NewRequest("POST", "/receptions", bytes.NewBuffer(jsonReq))
		req.Header.Set("Authorization", "Bearer "+tokenClient)
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req, s)
		checkResponseCode(t, http.StatusCreated, response.Code)
	})

	t.Run("Add Product", func(t *testing.T) {
		var wg sync.WaitGroup

		reqBody := map[string]string{
			"pvzId": pvzResp.ID.String(),
			"type":  "обувь",
		}

		jsonReq, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		for range 50 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonReq))
				req.Header.Set("Authorization", "Bearer "+tokenClient)
				req.Header.Set("Content-Type", "application/json")

				response := executeRequest(req, s)
				checkResponseCode(t, http.StatusCreated, response.Code)
			}()
		}
		wg.Wait()
	})

	t.Run("Close Reception", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/pvz/"+pvzResp.ID.String()+"/close_last_reception", nil)
		req.Header.Set("Authorization", "Bearer "+tokenClient)
		req.Header.Set("Content-Type", "application/json")

		response := executeRequest(req, s)
		checkResponseCode(t, http.StatusOK, response.Code)
	})
}

type httpServer struct {
	Storage storage.StorageImp
	Router  *chi.Mux
}

func newHTTPServer(storage storage.StorageImp) *httpServer {
	return &httpServer{
		Storage: storage,
		Router:  chi.NewRouter(),
	}
}

func (s *httpServer) mountHandlers() {
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
		r.Post("/register", auth.Register)
		r.Post("/login", login.Login)
	})

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

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}
