package rest

import (
	"log/slog"
	"net/http"

	"github.com/mattismoel/konnekt"
)

var codes = map[string]int{
	konnekt.ERRCONFLICT:       http.StatusConflict,
	konnekt.ERRINVALID:        http.StatusBadRequest,
	konnekt.ERRNOTFOUND:       http.StatusNotFound,
	konnekt.ERRNOTIMPLEMENTED: http.StatusNotImplemented,
	konnekt.ERRUNAUTHORIZED:   http.StatusUnauthorized,
	konnekt.ERRINTERNAL:       http.StatusInternalServerError,
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	code, message := konnekt.ErrorCode(err), konnekt.ErrorMessage(err)

	if code == konnekt.ERRINTERNAL {
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
