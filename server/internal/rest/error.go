package rest

import (
	"log/slog"
	"net/http"

	"github.com/mattismoel/konnekt/internal/service"
)

var codes = map[string]int{
	service.ERRCONFLICT:       http.StatusConflict,
	service.ERRINVALID:        http.StatusBadRequest,
	service.ERRNOTFOUND:       http.StatusNotFound,
	service.ERRNOTIMPLEMENTED: http.StatusNotImplemented,
	service.ERRUNAUTHORIZED:   http.StatusUnauthorized,
	service.ERRINTERNAL:       http.StatusInternalServerError,
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	code, message := service.ErrorCode(err), service.ErrorMessage(err)

	if code == service.ERRINTERNAL {
		logError(r, err)
	}

	httpCode := errorStatusCode(code)

	writeJSON(w, httpCode, apiError{
		Status:  httpCode,
		Message: message,
	})
}

func errorStatusCode(code string) int {
	if v, ok := codes[code]; ok {
		return v
	}

	return http.StatusInternalServerError
}

func logError(r *http.Request, err error) {
	slog.Error(err.Error(), "method", r.Method, "path", r.URL.Path)
}
