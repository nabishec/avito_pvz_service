package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
)

type Auth struct {
	PostAuth PostAuth
}

func NewAuth(postAuth PostAuth) *Auth {
	return &Auth{
		PostAuth: postAuth,
	}
}

// @Summary Регистрация
// @Description Регистрация пользователя
// @Accept json
// @Produce json
// @Param request body model.RegisterReq true "Данные нового пользователя"
// @Success 201 {object} model.RegisterResp "Пользователь создан"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /register  [post]
func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.hadnlers.auth.Register()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to register user has been received")

	var userReq model.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(userReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}
	if userReq.Role != "client" && userReq.Role != "moderator" {
		logger.Error().Msg("Invalid user role in request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	user, err := h.PostAuth.CreateUser(userReq.Email, userReq.Password, userReq.Role)
	if err != nil {
		if err == storage.ErrPasswordIsEmpty {
			logger.Error().Err(err).Msg("Password is empty")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
			return
		}
		if err == storage.ErrUserAlreadyExist {
			logger.Error().Err(err).Msg("User already exist")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
			return
		}
		logger.Error().Msg("Failed to add user to database")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	render.JSON(w, r, user)
}
