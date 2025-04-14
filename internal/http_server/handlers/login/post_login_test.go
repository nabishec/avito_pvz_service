package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
)

func TestLogin(t *testing.T) {
	mc := minimock.NewController(t)

	postLoginMock := NewPostLoginMock(mc)
	handler := Login{PostLogin: postLoginMock}

	reqBody := model.LoginReq{
		Email:    "bigballs@men.ru",
		Password: "hasbik",
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	t.Run("Successful login", func(t *testing.T) {
		postLoginMock.LoginMock.Expect(minimock.AnyContext, reqBody.Email, reqBody.Password).Return(uuid.New(), "client", nil)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`whats up`))

		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		badReqBody := model.LoginReq{
			Email:    "",
			Password: "",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Unauthorized user not exist", func(t *testing.T) {
		postLoginMock.LoginMock.Expect(minimock.AnyContext, reqBody.Email, reqBody.Password).Return(uuid.Nil, "", storage.ErrUserNotExist)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Status Unauthorized incorrect password of user", func(t *testing.T) {
		postLoginMock.LoginMock.Expect(minimock.AnyContext, reqBody.Email, reqBody.Password).Return(uuid.Nil, "", storage.ErrPasswordIsWrong)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Status Unauthorized incorrect password is empty", func(t *testing.T) {
		postLoginMock.LoginMock.Expect(minimock.AnyContext, reqBody.Email, reqBody.Password).Return(uuid.Nil, "", storage.ErrPasswordIsEmpty)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		postLoginMock.LoginMock.Expect(minimock.AnyContext, reqBody.Email, reqBody.Password).Return(uuid.Nil, "", errors.New("its bad time to work"))

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
