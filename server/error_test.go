package konnekt_test

import (
	"fmt"
	"testing"

	"github.com/mattismoel/konnekt"
)

func TestErrorCode(t *testing.T) {

	type test struct {
		error    error
		wantCode string
	}

	tests := map[string]test{
		"Non-empty error": {
			error:    konnekt.Errorf(konnekt.ERRNOTFOUND, "Test"),
			wantCode: konnekt.ERRNOTFOUND,
		},
		"Non-konnekt error": {
			error:    fmt.Errorf("Test Error"),
			wantCode: konnekt.ERRINTERNAL,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			code := konnekt.ErrorCode(tt.error)

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
		"Non-Konnekt error": {
			err:         fmt.Errorf("Error message"),
			wantMessage: "Internal error",
		},
		"Nil error": {
			err:         nil,
			wantMessage: "",
		},
		"Konnekt Error": {
			err:         konnekt.Errorf(konnekt.ERRNOTFOUND, "not_found"),
			wantMessage: "not_found",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			code := konnekt.ErrorMessage(tt.err)

			if code != tt.wantMessage {
				t.Fatalf("got code %q, want code %q", code, tt.wantMessage)
			}
		})
	}

}
