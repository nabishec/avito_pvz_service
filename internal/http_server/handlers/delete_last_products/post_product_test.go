package deletelastproducts

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

func TestDeleteLastProduct(t *testing.T) {
	mc := minimock.NewController(t)

	postDeleteLastProducts := NewPostDeleteLastProductsMock(mc)
	handler := DeleteLastProducts{PostDeleteLastProducts: postDeleteLastProducts}

	r := chi.NewRouter()
	r.Post("/pvz/{pvzId}/delete_last_product", handler.DeleteProducts)

	t.Run("Successful delete last product", func(t *testing.T) {
		pvzID := uuid.New().String()

		postDeleteLastProducts.DeleteLastProductsMock.Expect(minimock.AnyContext, uuid.MustParse(pvzID)).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Status Bad Request empty pvzID", func(t *testing.T) {
		pvzID := ""

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect pvzID", func(t *testing.T) {
		pvzID := "blabliblo"

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request open reception not exist", func(t *testing.T) {
		pvzID := uuid.New().String()

		postDeleteLastProducts.DeleteLastProductsMock.Expect(minimock.AnyContext, uuid.MustParse(pvzID)).Return(storage.ErrOpenReceptionNotExist)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request not product in request", func(t *testing.T) {
		pvzID := uuid.New().String()

		postDeleteLastProducts.DeleteLastProductsMock.Expect(minimock.AnyContext, uuid.MustParse(pvzID)).Return(storage.ErrProductsInReceptionNotExist)

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status forbidden", func(t *testing.T) {
		pvzID := uuid.New().String()

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "moderator"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("Status Internal Server Error", func(t *testing.T) {
		pvzID := uuid.New().String()

		postDeleteLastProducts.DeleteLastProductsMock.Expect(minimock.AnyContext, uuid.MustParse(pvzID)).Return(errors.New("good var love bad func"))

		req := httptest.NewRequest(http.MethodPost, "/pvz/"+pvzID+"/delete_last_product", nil)
		req = req.WithContext(context.WithValue(req.Context(), middleware.RequestUserRoleKey, "client"))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
