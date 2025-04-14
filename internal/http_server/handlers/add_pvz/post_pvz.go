package addpvz

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/rs/zerolog/log"
)

type PVZ struct {
	PostPVZ PostPVZ
}

func NewPVZ(postPVZ PostPVZ) *PVZ {
	return &PVZ{
		PostPVZ: postPVZ,
	}
}

// @Summary Добавление нового ПВЗ
// @Description Добавляет новый ПВЗ
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.PVZReq true "Город нового ПВЗ"
// @Success 201 {object} model.PVZResp "ПВЗ создан"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос."
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /pvz  [post]
func (h *PVZ) AddPVZ(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.addpvz.AddPVZ()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to add pvz been received")

	userRole := r.Context().Value(middleware.RequestUserRoleKey).(string)
	if userRole != "moderator" {
		logger.Error().Msg("User role doesn't have required permissions")

		w.WriteHeader(http.StatusForbidden) // 403
		render.JSON(w, r, model.ReturnErrResp("Доступ запрещен."))
		return
	}

	var pvzReq model.PVZReq

	err := json.NewDecoder(r.Body).Decode(&pvzReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(pvzReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	cityes := []string{"Москва", "Санкт-Петербург", "Казань"}
	cityCorrect := slices.Contains(cityes, pvzReq.City)
	if !cityCorrect {
		logger.Error().Err(err).Msg("Incorrect city in request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос."))
		return
	}

	pvz, err := h.PostPVZ.AddPVZ(pvzReq.City)
	if err != nil {
		logger.Error().Msg("Failed to add pvz")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	render.JSON(w, r, pvz)
}
