package getpvzlist

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/gojuno/minimock/v3"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

func TestGetPVZList(t *testing.T) {
	mc := minimock.NewController(t)

	getPVZListMock := NewGetPVZMock(mc)
	handler := PVZ{GetPVZ: getPVZListMock}

	startDate, _ := time.Parse(time.DateOnly, "2025-01-01")
	endDate, _ := time.Parse(time.DateOnly, "2025-12-04")

	t.Run("Successful get pvz list", func(t *testing.T) {
		getPVZListMock.GetPVZListMock.Expect(startDate, endDate, 1, 10).Return([]*model.PVZWithRecep{}, nil)

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+startDate.Format(time.DateOnly)+"&endDate="+endDate.Format(time.DateOnly), nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Status Bad Request incorrect Date", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+"yesterday"+"&endDate="+endDate.Format(time.DateOnly), nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect page", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+startDate.Format(time.DateOnly)+
			"&endDate="+endDate.Format(time.DateOnly)+"&page=0", nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect limit", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+startDate.Format(time.DateOnly)+
			"&endDate="+endDate.Format(time.DateOnly)+"&limit=35", nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect date", func(t *testing.T) {
		startDate, endDate := endDate, startDate
		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+startDate.Format(time.DateOnly)+"&endDate="+endDate.Format(time.DateOnly), nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Internal Server Error", func(t *testing.T) {
		getPVZListMock.GetPVZListMock.Expect(startDate, endDate, 1, 10).Return(nil, errors.New("today is Cosmonautics Day"))

		req := httptest.NewRequest(http.MethodGet, "/pvz?startDate="+startDate.Format(time.DateOnly)+"&endDate="+endDate.Format(time.DateOnly), nil)
		w := httptest.NewRecorder()
		handler.GetPVZList(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
