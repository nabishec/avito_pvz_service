package addproducts

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	middleware "github.com/nabishec/avito_pvz_service/internal/http_server/middleware"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/storage"
	"github.com/rs/zerolog/log"
)

type Products struct {
	PostProducts PostProducts
}

func NewProducts(postProducts PostProducts) *Products {
	return &Products{
		PostProducts: postProducts,
	}
}

// @Summary Добавление товара
// @Description Добавление товара в текущую приемку (только для сотрудников ПВЗ)
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.ProductsReq true "Данные продукта"
// @Success 201 {object} model.ProductsResp "Товар добавлен."
// @Failure 400 {object} model.ErrorResponse "Неверный запрос или нет активной приемки."
// @Failure 401 {object} model.ErrorResponse "Неавторизован."
// @Failure 403 {object} model.ErrorResponse "Доступ запрещен."
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера."
// @Router /products  [post]
func (h *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http_server.handlers.addproducts.AddProducts()"

	logger := log.With().Str("fn", op).Logger()
	logger.Debug().Msg("Request to add products has been received")

	userRole := r.Context().Value(middleware.RequestUserRoleKey).(string)
	if userRole != "client" {
		logger.Error().Msg("User role doesn't have required permissions")

		w.WriteHeader(http.StatusForbidden) // 403
		render.JSON(w, r, model.ReturnErrResp("Доступ запрещен."))
		return
	}

	var productsReq model.ProductsReq

	err := json.NewDecoder(r.Body).Decode(&productsReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to decode request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}
	logger.Debug().Msg("Request body decoded")

	err = validator.New().Struct(productsReq)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to validate request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}

	types := []string{"электроника", "одежда", "обувь"}
	typesCorrect := slices.Contains(types, productsReq.Type)
	if !typesCorrect {
		logger.Error().Err(err).Msg("Incorrect product in request body")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}

	pvzID, err := uuid.Parse(productsReq.PVZID)
	if err != nil {
		logger.Error().Msg("Failed to get pvzID")

		w.WriteHeader(http.StatusBadRequest) // 400
		render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
		return
	}

	product, err := h.PostProducts.AddProduct(r.Context(), pvzID, productsReq.Type)
	if err != nil {
		if err == storage.ErrOpenReceptionNotExist {
			logger.Error().Msg("Active reception not exist")

			w.WriteHeader(http.StatusBadRequest) // 400
			render.JSON(w, r, model.ReturnErrResp("Неверный запрос или нет активной приемки."))
			return
		}
		logger.Error().Msg("Failed to add product")

		w.WriteHeader(http.StatusInternalServerError) // 500
		render.JSON(w, r, model.ReturnErrResp("Внутренняя ошибка сервера."))
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
	render.JSON(w, r, product)
}
