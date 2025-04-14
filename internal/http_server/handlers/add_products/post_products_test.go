package addproducts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestAddProducts(t *testing.T) {
	mc := minimock.NewController(t)

	postProductsMock := NewPostProductsMock(mc)
	handler := Products{PostProducts: postProductsMock}

	reqBody := model.ProductsReq{
		PVZID: uuid.New().String(),
		Type:  "обувь",
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	t.Run("Successful addition of product", func(t *testing.T) {
		postProductsMock.AddProductMock.Expect(uuid.MustParse(reqBody.PVZID), reqBody.Type).Return(&model.ProductsResp{
			ID: uuid.New(),
		}, nil)

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Status Bad Request incorrect product", func(t *testing.T) {
		badReqBody := model.ProductsReq{
			PVZID: uuid.New().String(),
			Type:  "харамные шорты",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect pvzID", func(t *testing.T) {
		badReqBody := model.ProductsReq{
			PVZID: "dodo",
			Type:  "обувь",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(`badreq`))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		badReqBody := model.ProductsReq{
			PVZID: "",
			Type:  "",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status bad request when active reception not exist", func(t *testing.T) {
		postProductsMock.AddProductMock.Expect(uuid.MustParse(reqBody.PVZID), reqBody.Type).Return(nil, storage.ErrOpenReceptionNotExist)

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Fobidden", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		postProductsMock.AddProductMock.Expect(uuid.MustParse(reqBody.PVZID), reqBody.Type).Return(nil, errors.New("lets tomorrow"))

		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddProducts(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
