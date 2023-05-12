package middles

import (
	"net/http"

	sqlite "github.com/FreezeSnail/urlShortener/src/db"
	"golang.org/x/exp/slog"
)

type Middleware struct {
	DB  *sqlite.SQLite
	Log *slog.Logger
}

func (m Middleware) ValidateAPIKey() func(next http.Handler) http.Handler {
	m.Log.Info("validating api key")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			valid, err := m.DB.ValidateAPIKey(r.Context(), apiKey)
			if err != nil {
				return
			}

			if !valid {
				m.Log.Info("apikey invalid")
				apiKeyAuthFailed(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func apiKeyAuthFailed(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", "invalid api key")
	w.WriteHeader(http.StatusUnauthorized)
}
