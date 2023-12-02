package main;

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {input: "a/b/c", want: ""},
        {input: "a/b/c", want: ""},
        {input: "abc", want: ""},
        {input: "", want: ""},
    };

    for _, tc := range tests {
        got, err := GetAPIKey(tc.input)
        if err != nil {
			t.Fatalf("GetAPIKey returned error:", err);
        }
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    }
}
