package roottoken

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/pkg/token"
	"github.com/rs/zerolog/log"
)

type RootToken struct {
}

func NewRootToken() *RootToken {
	return &RootToken{}
}

// @Summary Получение тестового токена
// @Description Возвращает JWT-токен с указанной ролью (employee или moderator)
// @Accept json
// @Produce json
// @Param request body model.RootTokenReq true "Роль для тестового токена"
// @Success 200 {object} model.AuthResponse "Успешная авторизация."
// @Failure 400 {object} model.ErrorResponse "Неверный запрос."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /dummyLogin  [post]
func (h *RootToken) ReturnRootToken(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.hadnlers.root_roken.ReturnRootToken()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request for get jwt token has been received")

	var rootToken model.RootTokenReq

	err := json.NewDecoder(r.Body).Decode(&rootToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(rootToken)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	if rootToken.UserRole != "client" && rootToken.UserRole != "moderator" {
		logger.Error().Msg("Invalid user role in request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	token, err := token.CreateJWT(uuid.New(), rootToken.UserRole)
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
