package closelastreceptions

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
)

type CloseLastReceptions struct {
	PostCloseLastReceptions PostCloseLastReceptions
}

func NewCloseLastReceptions(postCloseLastReceptions PostCloseLastReceptions) *CloseLastReceptions {
	return &CloseLastReceptions{
		PostCloseLastReceptions: postCloseLastReceptions,
	}
}

// @Summary Закрытие приемки
// @Description Закрытие последней открытой приемки товаров в рамках ПВЗ
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pvzId path string true "ID ПВЗ"
// @Success 200 {object} model.SuccResp "Приемка закрыта"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос или приемка уже закрыта"
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /pvz/{pvzId}/close_last_reception  [post]
func (h *CloseLastReceptions) CloseLastReceptions(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.close_last_receptions.CloseLastReceptions()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to close receptions has been received")

	userRole := r.Context().Value(middleware.RequestUserRoleKey).(string)
	if userRole != "client" {
		logger.Error().Msg("User role doesn't have required permissions")

		w.WriteHeader(http.StatusForbidden) // 403
		render.JSON(w, r, model.ReturnErrResp("Доступ запрещен."))
		return
	}

	pvzIDString := chi.URLParam(r, "pvzId")
	if pvzIDString == "" {
		logger.Error().Msg("PVZ ID is empty")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или приемка уже закрыта"))
		return
	}
	logger.Debug().Msg("pvzID are received")

	pvzID, err := uuid.Parse(pvzIDString)
	if err != nil {
		logger.Error().Msg("Failed to get pvzID")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	err = h.PostCloseLastReceptions.CloseLastReceptions(pvzID)
	if err != nil {
		if err == storage.ErrOpenReceptionNotExist {
			logger.Error().Msg("Active reception not exist")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или приемка уже закрыта"))
			return
		}
		logger.Error().Msg("Failed Close product")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	render.JSON(w, r, model.ReturnSuccResp("Приемка закрыта."))
}
