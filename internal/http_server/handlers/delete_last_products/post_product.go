package deletelastproducts

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

type DeleteLastProducts struct {
	PostDeleteLastProducts PostDeleteLastProducts
}

func NewDeleteProducts(postDeleteLastProducts PostDeleteLastProducts) *DeleteLastProducts {
	return &DeleteLastProducts{
		PostDeleteLastProducts: postDeleteLastProducts,
	}
}

// @Summary Удаление товара
// @Description Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pvzId path string true "ID ПВЗ"
// @Success 200 {object} model.SuccResp "Товар удален"
// @Failure 400 {object} model.ErrorResponse "Неверный запрос, нет активной приемки или нет товаров для удаления"
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /pvz/{pvzId}/delete_last_product  [post]
func (h *DeleteLastProducts) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.hadnlers.deletelastproducts.DeleteLastProducts()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to delete products has been received")

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
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}
	logger.Debug().Msg("pvzID are received")

	pvzID, err := uuid.Parse(pvzIDString)
	if err != nil {
		logger.Error().Msg("Failed to get pvzID")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}

	err = h.PostDeleteLastProducts.DeleteLastProducts(pvzID)
	if err != nil {
		if err == storage.ErrOpenReceptionNotExist {
			logger.Error().Msg("Active reception not exist")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
			return
		}
		if err == storage.ErrProductsInReceptionNotExist {
			logger.Error().Msg("Active reception is empty")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
			return
		}

		logger.Error().Msg("Failed delete product")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	render.JSON(w, r, model.ReturnSuccResp("Товар удален"))
}
