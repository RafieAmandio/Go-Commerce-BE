package middleware

import (
	"mooi/library/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Log.Info("request completed in ", time.Since(start))
	})
}
