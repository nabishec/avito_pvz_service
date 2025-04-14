package addpvz

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

const defaultCity = "Москва"

func TestAddPVZ(t *testing.T) {
	mc := minimock.NewController(t)

	postPVZMock := NewPostPVZMock(mc)
	handler := PVZ{PostPVZ: postPVZMock}

	t.Run("Successful addition to PVZ", func(t *testing.T) {
		city := defaultCity

		postPVZMock.AddPVZMock.Expect(minimock.AnyContext, city).Return(&model.PVZResp{
			ID:               uuid.New(),
			RegistrationDate: time.Now(),
			City:             city,
		}, nil)

		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`{"city":"`+city+`"}`))

		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Status Bad Request incorrect city", func(t *testing.T) {
		city := "Махачкала"

		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`{"city":"`+city+`"}`))

		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`badreq`))

		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`{"city":""}`))

		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status forbidden", func(t *testing.T) {
		city := defaultCity

		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`{"city":"`+city+`"}`))

		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Status Internal Server Error", func(t *testing.T) {
		city := defaultCity

		req := httptest.NewRequest(http.MethodPost, "/pvz", bytes.NewBufferString(`{"city":"`+city+`"}`))
		postPVZMock.AddPVZMock.Expect(minimock.AnyContext, city).Return(nil, errors.New("lazy func won't do anything"))
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		handler.AddPVZ(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
