package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggerResponseWriter(w http.ResponseWriter) *loggerResponseWriter {
	return &loggerResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (l *loggerResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

func Logger(n http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		l := newLoggerResponseWriter(w)

		n.ServeHTTP(l, r)
		since := time.Since(start)

		log.Printf("%s %s %d %s", r.Method, r.URL.Path, l.statusCode, since)
	}
}
