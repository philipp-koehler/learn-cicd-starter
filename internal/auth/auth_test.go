package auth

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"success": {
			input: http.Header{"Authorization": {"ApiKey test"}},
			want:  "test"},
		"noApi": {
			input: http.Header{"Authorization": {"ApiKey"}},
			want:  ""},
		"missing": {
			input: http.Header{"Authorizatio": {"ApiKey test"}},
			want:  ""},
		"spelling": {
			input: http.Header{"Authorization": {"ApiKeey test"}},
			want:  ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
