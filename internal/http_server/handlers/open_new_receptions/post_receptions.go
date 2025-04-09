package opennewreceptions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
)

type Receptions struct {
	PostReceptions PostReceptions
}

func NewReceptions(postReceptions PostReceptions) *Receptions {
	return &Receptions{
		PostReceptions: postReceptions,
	}
}

// @Summary Добавление новой приемки
// @Description Создание новой приемки товаров (только для сотрудников ПВЗ)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.ReceptionsReq true "ID ПВЗ"
// @Success 201 {object} model.ReceptionsResp "Приемка создана."
// @Failure 400 {object} model.ErrorResponse "Неверный запрос или есть незакрытая приемка."
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /receptions  [post]
func (h *Receptions) AddReceptions(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.opennewreceptions.AddReceptions()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to add receptions has been received")

	userRole := r.Context().Value(middleware.RequestUserRoleKey).(string)
	if userRole != "client" {
		logger.Error().Msg("User role doesn't have required permissions")

		w.WriteHeader(http.StatusForbidden) // 403
		render.JSON(w, r, model.ReturnErrResp("Доступ запрещен."))
		return
	}

	var receptionReq model.ReceptionsReq

	err := json.NewDecoder(r.Body).Decode(&receptionReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или есть незакрытая приемка."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(receptionReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или есть незакрытая приемка."))
		return
	}

	pvzID, err := uuid.Parse(receptionReq.PVZID)
	if err != nil {
		logger.Error().Msg("Failed to get pvzID")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	receptions, err := h.PostReceptions.AddReception(pvzID)
	if err != nil {
		if err == storage.ErrPVZNotExist {
			logger.Error().Msg("PVZ not exist yet")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или есть незакрытая приемка."))
			return
		}
		if err == storage.ErrPreviousReceptionNotClosed {
			logger.Error().Msg("Previous reception not closed")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или есть незакрытая приемка."))
			return
		}
		logger.Error().Msg("Failed to add reception")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	render.JSON(w, r, receptions)
}
