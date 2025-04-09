package login

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/pkg/token"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
)

type Login struct {
	PostLogin PostLogin
}

func NewLogin(postLogin PostLogin) *Login {
	return &Login{
		PostLogin: postLogin,
	}
}

// @Summary Авторизация
// @Description Авторизация пользователя
// @Accept json
// @Produce json
// @Param request body model.LoginReq true "Данные авторизации пользователя"
// @Success 200 {object} model.AuthResponse "Успешная авторизация"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос."
// @Failure 401 {object} model.ErrorResponse "Неверные учетные данные"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /login  [post]
func (h *Login) Login(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.login.Login()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to login user has been received")

	var loginReq model.LoginReq
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(loginReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	userID, role, err := h.PostLogin.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		if err == storage.ErrPasswordIsEmpty {
			logger.Error().Err(err).Msg("Password is empty")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
			return
		}
		if err == storage.ErrPasswordIsWrong {
			logger.Error().Err(err).Msg("Password is wrong")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
			return
		}
		if err == storage.ErrUserNotExist {
			logger.Error().Err(err).Msg("User not exist")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
			return
		}
		logger.Error().Msg("Failed to add user to database")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	token, err := token.CreateJWT(userID, role)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create JWT token")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}
	logger.Debug().Msg("JWT token created")

	w.WriteHeader(http.StatusOK) // 200
	render.JSON(w, r, model.AuthResponse{Token: token})
}
