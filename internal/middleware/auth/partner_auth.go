package auth

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/amnestia/xyz-multifinance/internal/lib/crypto/aes"
	"github.com/amnestia/xyz-multifinance/internal/lib/paseto"
	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/amnestia/xyz-multifinance/pkg/response"
)

// AuthorizeAPIKey authorize by API Key for partner
func (a *AuthorizationModule) AuthorizeAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		clientID := strings.TrimSpace(r.Header.Get("X-Client-ID"))
		if len(clientID) < 1 {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Unknown client").WriteJSON(w)
			return
		}
		apiKey := strings.TrimSpace(r.Header.Get("X-API-Key"))
		if len(apiKey) < 1 {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Unknown client").WriteJSON(w)
			return
		}

		partner, err := a.AuthRepo.GetPartner(ctx, clientID)
		if err != nil {
			if err == sql.ErrNoRows {
				response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Unknown client").WriteJSON(w)
				return
			}
			response.NewResponse(ctx).SetError(logger.ErrorWrap(err, "AuthorizeAPIKey.GetPartner").Error()).SetResponse(http.StatusInternalServerError, "", "").WriteJSON(w)
			return
		}
		dec, err := aes.Decrypt(partner.APIKey, a.Config.Crypt.AESKey)
		if err != nil {
			response.NewResponse(ctx).SetError(logger.ErrorWrap(err, "AuthorizeAPIKey.DecryptKey").Error()).SetResponse(http.StatusInternalServerError, "", "").WriteJSON(w)
			return
		}
		if apiKey != dec {
			response.NewResponse(ctx).SetResponse(http.StatusUnauthorized, "", "Unknown client").WriteJSON(w)
			return
		}
		ctx = context.WithValue(ctx, paseto.AuthData, partner)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
