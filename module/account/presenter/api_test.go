package presenter

import (
	"context"
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/moeghifar/mariana/config"
	"github.com/moeghifar/mariana/module/account/entity"
	"github.com/moeghifar/mariana/module/account/usecase"
)

type stubUsecase struct{ err error }

func (s stubUsecase) GetByID(_ context.Context, id int64) (entity.User, error) {
	return entity.User{ID: id, Name: "Ada", Email: "ada@example.com"}, s.err
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		err    error
		status int
		want   string
		hidden string
	}{
		{name: "found", path: "/v1/users/7", status: 200, want: `"name":"Ada"`},
		{name: "bad id", path: "/v1/users/abc", status: 400, want: "ACC-400"},
		{name: "not found", path: "/v1/users/7", err: entity.ErrUserNotFound, status: 404, want: "ACC-404"},
		{name: "internal", path: "/v1/users/7", err: errors.New("dsn secret"), status: 500, hidden: "dsn secret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucs := usecase.Usecases{MyUsecase: stubUsecase{err: tt.err}}
			resp, err := newAPI(config.Config{}, ucs).App().Test(httptest.NewRequest("GET", tt.path, nil))
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			raw, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			body := string(raw)

			if resp.StatusCode != tt.status {
				t.Fatalf("status = %d, want %d, body %s", resp.StatusCode, tt.status, body)
			}
			if !strings.Contains(body, tt.want) {
				t.Errorf("body %s does not contain %q", body, tt.want)
			}
			if tt.hidden != "" && strings.Contains(body, tt.hidden) {
				t.Errorf("body leaks %q: %s", tt.hidden, body)
			}
		})
	}
}
