package server

import (
	"fmt"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	bytesWritten int
	statusCode   int
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytesWritten += n
	return n, err
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)
		timeStr := start.Format(time.RFC3339)

		// Format [<METHOD>|<status>] <bytes>b <time>: <path>
		fmt.Printf("[%s|%d] %db %s: %s\n", r.Method, lrw.statusCode, lrw.bytesWritten, timeStr, r.URL.Path)
	})
}
