package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "no auth header",
			args: args{
				headers: http.Header{},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed auth header",
			args: args{
				headers: http.Header{
					"Authorization": []string{"ApiKey"},
				},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "valid auth header",
			args: args{
				headers: http.Header{
					"Authorization": []string{"ApiKey abc123"},
				},
			},
			want:    "abc123",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
