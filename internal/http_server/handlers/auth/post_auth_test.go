package auth

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

func TestAuth(t *testing.T) {
	mc := minimock.NewController(t)

	postAuthMock := NewPostAuthMock(mc)
	handler := Auth{PostAuth: postAuthMock}

	reqBody := model.RegisterReq{
		Email:    "bigballs@men.ru",
		Password: "hasbik",
		Role:     "moderator",
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("failed to marshal request body")
	}

	t.Run("Successful registration", func(t *testing.T) {

		postAuthMock.CreateUserMock.Expect(reqBody.Email, reqBody.Password, reqBody.Role).Return(&model.RegisterResp{
			ID:    uuid.New(),
			Email: reqBody.Email,
			Role:  reqBody.Role,
		}, nil)
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Status Bad Request incorrect role", func(t *testing.T) {

		badReqBody := model.RegisterReq{
			Email:    "bigballs@men.ru",
			Password: "hasbik",
			Role:     "slave",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request incorrect body", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`badreq`))

		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status Bad Request is no required value", func(t *testing.T) {
		badReqBody := model.RegisterReq{
			Email:    "",
			Password: "",
			Role:     "",
		}

		jsonReq, err := json.Marshal(badReqBody)
		if err != nil {
			t.Fatal("failed to marshal request body")
		}

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Status bad request user exist", func(t *testing.T) {

		postAuthMock.CreateUserMock.Expect(reqBody.Email, reqBody.Password, reqBody.Role).Return(nil, storage.ErrUserAlreadyExist)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Internal Server Error", func(t *testing.T) {

		postAuthMock.CreateUserMock.Expect(reqBody.Email, reqBody.Password, reqBody.Role).Return(nil, errors.New("not for u"))

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		handler.Register(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

}
