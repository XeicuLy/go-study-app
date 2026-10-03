package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XeicuLy/go-study-app/internal/model"
	"github.com/XeicuLy/go-study-app/internal/repository/memory"
	"github.com/XeicuLy/go-study-app/internal/service"
)

func TestRedirectFound(t *testing.T) {
	seed := []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}}
	// 1. memory.NewRepository(seed) で repository を作る
	repo := memory.NewRepository(seed)
	// 2. service.NewLinkService(repo) で service を作る
	svc := service.NewLinkService(repo)
	// 3. NewLinkHandler(svc) で handler を作る
	h := NewLinkHandler(svc)
	// 4. NewRouter(h) で router を作る
	router := NewRouter(h)
	// 5. httptest.NewRequest(http.MethodGet, "/godev01", nil) でリクエストを作る
	req := httptest.NewRequest(http.MethodGet, "/godev01", nil)
	// 6. httptest.NewRecorder() でレスポンスの書き込み先を作る
	res := httptest.NewRecorder()
	// 7. router.ServeHTTP(rec, req) で呼び出す (h.Redirect を直接呼ばない)
	router.ServeHTTP(res, req)
	// 8. rec.Code が http.StatusFound (302) か確認する。違っていたら t.Errorf
	if res.Code != http.StatusFound {
		t.Errorf("status:%v, want:%v", res.Code, http.StatusFound)
	}
	// 9. rec.Header().Get("Location") が "https://go.dev" か確認する。違っていたら t.Errorf
	if res.Header().Get("Location") != "https://go.dev" {
		t.Errorf("location:%v, want:%v", res.Header().Get("Location"), "https://go.dev")
	}
}

func TestRedirectNotFound(t *testing.T) {
	// TestRedirectFound と同じ形。違うのは次の3点だけ
	// - repository は memory.NewRepository(nil) で空にする
	// - リクエストは "/nosuchcode"
	// - 確認するのは rec.Code が http.StatusNotFound (404) であること
	repo := memory.NewRepository(nil)
	svc := service.NewLinkService(repo)
	h := NewLinkHandler(svc)
	router := NewRouter(h)
	req := httptest.NewRequest(http.MethodGet, "/nosuchcode", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusNotFound {
		t.Errorf("status:%v, want:%v", res.Code, http.StatusNotFound)
	}
}
