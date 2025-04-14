package closelastreceptions

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-playground/assert/v2"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/storage"
)

func TestCloseLastReception(t *testing.T) {
	mc := minimock.NewController(t)

	postCloseLastReceptionsMock := NewPostCloseLastReceptionsMock(mc)
	handler := CloseLastReceptions{PostCloseLastReceptions: postCloseLastReceptionsMock}

	r := chi.NewRouter()
	r.Post("/pvz/{pvzId}/close_last_reception", handler.CloseLastReceptions)

	t.Run("Successful close last reception", func(t *testing.T) {
		pvzID := uuid.New().String()

		postCloseLastReceptionsMock.CloseLastReceptionsMock.Expect(context.Background(), uuid.MustParse(pvzID)).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Status Bad Request empty pvzID", func(t *testing.T) {
		pvzID := ""

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect pvzID", func(t *testing.T) {
		pvzID := "badpvzID"

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request open reception not exist", func(t *testing.T) {
		pvzID := uuid.New().String()

		postCloseLastReceptionsMock.CloseLastReceptionsMock.Expect(context.Background(), uuid.MustParse(pvzID)).Return(storage.ErrOpenReceptionNotExist)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status forbidden", func(t *testing.T) {
		pvzID := uuid.New().String()

		postCloseLastReceptionsMock.CloseLastReceptionsMock.Expect(context.Background(), uuid.MustParse(pvzID)).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Status Internal Server Error", func(t *testing.T) {
		pvzID := uuid.New().String()

		postCloseLastReceptionsMock.CloseLastReceptionsMock.Expect(context.Background(), uuid.MustParse(pvzID)).Return(errors.New("not today baby"))

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/close_last_reception", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
