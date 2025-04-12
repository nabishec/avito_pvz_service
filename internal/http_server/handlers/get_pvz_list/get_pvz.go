package getpvzlist

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/render"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/rs/zerolog/log"
)

type PVZ struct {
	GetPVZ GetPVZ
}

func NewPVZ(getPVZ GetPVZ) *PVZ {
	return &PVZ{
		GetPVZ: getPVZ,
	}
}

// @Summary Получение списка ПВЗ
// @Description Получение списка ПВЗ с фильтрацией по дате приемки и пагинацией
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param startDate query string false "Начальная дата"
// @Param endDate query string false "Конечная дата"
// @Param page query int false "Номер страницы (начиная с 1)" default(1) minimum(1)
// @Param limit query int false "Количество элементов на странице" default(10) minimum(1) maximum(30)
// @Success 200 {object} []model.PVZWithRecep "Список ПВЗ"
// @Failure 400 {object} model.ErrorResponse "Неверный формат параметров"
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /pvz  [get]
func (h *PVZ) GetPVZList(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.getpvzlist.GetPVZList()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request for get pvz list has been received")

	defaultPage := os.Getenv("DEFAULT_PAGE")
	if defaultPage == "" {
		defaultPage = "1"
	}
	defaultLimit := os.Getenv("DEFAULT_LIMIT")
	if defaultLimit == "" {
		defaultLimit = "10"
	}

	var startDate = r.URL.Query().Get("startDate")
	var endDate = r.URL.Query().Get("endDate")
	var page = r.URL.Query().Get("page")
	var limit = r.URL.Query().Get("limit")

	var startDateTime time.Time
	var endDateTime time.Time
	var err error

	if startDate != "" {
		startDateTime, err = time.Parse(time.DateOnly, startDate)
		if err != nil {
			logger.Error().Err(err).Msg("Incorrect start time parameter")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный формат параметров"))
			return
		}
	}

	if endDate != "" {
		endDateTime, err = time.Parse(time.DateOnly, endDate)
		if err != nil {
			logger.Error().Err(err).Msg("Incorrect end time parameter")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный формат параметров"))
			return
		}
	}

	var pageInt int
	var limitInt int
	if page == "" {
		pageInt, err = strconv.Atoi(defaultPage)
	} else {
		pageInt, err = strconv.Atoi(page)
	}

	if err != nil || pageInt < 1 {
		logger.Error().Err(err).Msg("Incorrect page parameter")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный формат параметров"))
		return
	}

	if limit == "" {
		limitInt, err = strconv.Atoi(defaultLimit)
	} else {
		limitInt, err = strconv.Atoi(limit)
	}

	if err != nil || limitInt < 1 || limitInt > 30 {
		logger.Error().Err(err).Msg("Incorrect limit parameter")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный формат параметров"))
		return
	}

	if endDateTime.Before(startDateTime) {
		logger.Error().Msg("End date is before start date")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный формат параметров"))
		return
	}

	pvzList, err := h.GetPVZ.GetPVZList(startDateTime, endDateTime, pageInt, limitInt)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get pvz list")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	render.JSON(w, r, pvzList)
}
