package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	dbconnection "github.com/nabishec/avito_pvz_service/cmd/db_connection"
	"github.com/nabishec/avito_pvz_service/internal/model"
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
	err := LoadEnv()
	if err != nil {
		t.Error("Don't found config")
	}

	dbConnection, err := dbconnection.NewDatabaseConnection()
	if err != nil {
		t.Error("Failed init database")
	}

	storage := db.NewStorage(dbConnection.DB)

	s := NewHTTPServer(storage)
	s.MountHandlers()
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
