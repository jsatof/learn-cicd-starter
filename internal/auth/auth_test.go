package auth;

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        input http.Header
        want  string
    }{
        {input: make(map[string][]string), want: ""},
    };

    for _, tc := range tests {
        got, err := GetAPIKey(tc.input)
        _ = err;
        //if err != nil {
		//	t.Fatalf("GetAPIKey returned error: %s", err);
        //}
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    }
}
