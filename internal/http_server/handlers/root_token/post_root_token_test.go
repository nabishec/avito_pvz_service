package roottoken

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

func TestPostRootToken(t *testing.T) {

	handler := RootToken{}

	reqBody := model.RootTokenReq{
		UserRole: "client",
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	t.Run("Successful get root roken", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/dummyLogin", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.ReturnRootToken(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/dummyLogin", bytes.NewBufferString(`im tired`))
		w := httptest.NewRecorder()
		handler.ReturnRootToken(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		badReqBody := model.RootTokenReq{
			UserRole: "",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/dummyLogin", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.ReturnRootToken(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect role", func(t *testing.T) {
		badReqBody := model.RootTokenReq{
			UserRole: "slave",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/dummyLogin", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.ReturnRootToken(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
