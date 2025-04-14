package opennewreceptions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
)

func TestOpenReception(t *testing.T) {
	mc := minimock.NewController(t)

	postReceptionsMock := NewPostReceptionsMock(mc)
	handler := Receptions{PostReceptions: postReceptionsMock}

	reqBody := model.ReceptionsReq{
		PVZID: uuid.New().String(),
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	t.Run("Successful opening of reception", func(t *testing.T) {
		postReceptionsMock.AddReceptionMock.Expect(uuid.MustParse(reqBody.PVZID)).Return(&model.ReceptionsResp{}, nil)
		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Status Bad Request incorrct pvzID", func(t *testing.T) {
		badReqBody := model.ReceptionsReq{
			PVZID: "hi babys",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBufferString(`aloha`))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		badReqBody := model.ReceptionsReq{
			PVZID: "",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status bad request when pvz not exist", func(t *testing.T) {
		postReceptionsMock.AddReceptionMock.Expect(uuid.MustParse(reqBody.PVZID)).Return(nil, storage.ErrPVZNotExist)

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status bad request when previous reception not closed", func(t *testing.T) {
		postReceptionsMock.AddReceptionMock.Expect(uuid.MustParse(reqBody.PVZID)).Return(nil, storage.ErrPreviousReceptionNotClosed)

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Fobidden", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		postReceptionsMock.AddReceptionMock.Expect(uuid.MustParse(reqBody.PVZID)).Return(nil, errors.New("no baby not today"))

		req := httptest.NewRequest(http.MethodPost, "/receptions", bytes.NewBuffer(jsonReq))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddReceptions(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
