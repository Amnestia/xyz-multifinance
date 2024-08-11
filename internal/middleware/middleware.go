package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/amnestia/xyz-multifinance/pkg/logger"
	"github.com/amnestia/xyz-multifinance/pkg/response"
	"github.com/go-chi/chi/v5/middleware"
)

// Logger log api request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()
		next.ServeHTTP(ww, r)
		elapsed := time.Since(start)
		reqURI := r.URL.RequestURI()
		logger.Logger.Info().Fields(map[string]interface{}{
			"latency":     elapsed,
			"request-id":  r.Header.Get("X-Request-ID"),
			"request-uri": reqURI,
			"source":      r.RemoteAddr,
			"status-code": ww.Status(),
		}).Msgf("%v %v: %v", r.Method, reqURI, ww.Status())

	})
}

// PanicRecovery panic recovery handler
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				e := logger.ErrorWrap(fmt.Errorf("panic occurred on :%v, %v\r\n", r.URL.Path, err), "", "")
				logger.Logger.Error().Err(e).Msg(string(debug.Stack()))
				response.NewResponse(r.Context()).SetResponse(http.StatusInternalServerError, nil, "").WriteJSON(w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
