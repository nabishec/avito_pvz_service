package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/nabishec/avito_pvz_service/internal/model"
	"github.com/nabishec/avito_pvz_service/internal/pkg/token"
	"github.com/rs/zerolog/log"
)

type ctxKeyRequestUserID string
type ctxKeyRequestUserRole string

const RequestUserRoleKey ctxKeyRequestUserRole = "user_role"
const RequestUserIDKey ctxKeyRequestUserID = "user_id"

func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		const op = "internal.http_server.middlweare.Authorization()"

		logger := log.With().Str("fn", op).Logger()
		logger.Debug().Msg("auth start")

		authLine := r.Header.Get("Authorization")

		if authLine == "" {
			logger.Error().Msg("Request doesn't specify an authorization header")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неавторизован."))
			return
		}

		auth := strings.Split(authLine, " ")

		if auth[0] != "Bearer" {
			logger.Error().Msg("Invalid authorization scheme")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неавторизован."))
			return
		}

		if len(auth) != 2 {
			logger.Error().Msg("Invalid authorization line")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неавторизован."))
			return
		}

		requestUserID, requestUserRole, err := token.CheckJWT(auth[1])

		if err != nil {
			logger.Error().Err(err)

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неавторизован."))
			return
		}

		if requestUserRole != "client" && requestUserRole != "moderator" {
			logger.Error().Msg("Invalid user role")

			w.WriteHeader(http.StatusUnauthorized) // 401
			render.JSON(w, r, model.ReturnErrResp("Неавторизован."))
			return
		}

		r.Header.Del("Authorization")
		ctx := context.WithValue(r.Context(), RequestUserIDKey, requestUserID)
		ctx = context.WithValue(ctx, RequestUserRoleKey, requestUserRole)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
