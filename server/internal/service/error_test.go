package service_test

import (
	"fmt"
	"testing"

	"github.com/mattismoel/konnekt/internal/service"
)

func TestErrorCode(t *testing.T) {
	type test struct {
		error    error
		wantCode string
	}

	tests := map[string]test{
		"Non-empty error": {
			error:    service.Errorf(service.ERRNOTFOUND, "Test"),
			wantCode: service.ERRNOTFOUND,
		},
		"Non-service.error": {
			error:    fmt.Errorf("Test Error"),
			wantCode: service.ERRINTERNAL,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			code := service.ErrorCode(tt.error)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q", code, tt.wantCode)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	type test struct {
		err         error
		wantMessage string
	}

	tests := map[string]test{
		"Non-service.error": {
			err:         fmt.Errorf("Error message"),
			wantMessage: "Internal error",
		},
		"Nil error": {
			err:         nil,
			wantMessage: "",
		},
		"service.Error": {
			err:         service.Errorf(service.ERRNOTFOUND, "not_found"),
			wantMessage: "not_found",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			code := service.ErrorMessage(tt.err)

			if code != tt.wantMessage {
				t.Fatalf("got code %q, want code %q", code, tt.wantMessage)
			}
		})
	}

}
