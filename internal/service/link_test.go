package service

import (
	"errors"
	"testing"

	"github.com/XeicuLy/go-study-app/internal/model"
	"github.com/XeicuLy/go-study-app/internal/repository/memory"
)

//  1. シードに入れたコードを Get すると、そのLinkが返る
//     a. seed を作る: []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}}
//     b. memory.NewRepository(seed) で repository を作る
//     c. NewLinkService(repo) で service を作る
//     d. svc.Get("godev01") を呼び、err が nil、OriginalURL が期待どおりであることを確認する
//     (必要な import: model と repository/memory。memory のパッケージ名は memory)
func TestGetFound(t *testing.T) {
	seed := []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}}
	repo := memory.NewRepository(seed)
	svc := NewLinkService(repo)
	result, err := svc.Get("godev01")
	if err != nil {
		t.Fatalf("Unknown Error: %v", err)
	}
	want := "https://go.dev"
	if result.OriginalURL != want {
		t.Errorf("result = %v, want = %v", result.OriginalURL, want)
	}
}

//  2. 無いコードだと model.ErrNotFound が返る
//     a. memory.NewRepository(nil) で空の repository を作り、NewLinkService に渡す
//     b. svc.Get("nosuchcode") を呼ぶ
//     c. errors.Is(err, model.ErrNotFound) で判定する。ErrNotFound「じゃなかったら」失敗
func TestGetNotFound(t *testing.T) {
	repo := memory.NewRepository(nil)
	svc := NewLinkService(repo)
	_, err := svc.Get("nosuchcode")
	if !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("NewLinkService.Get returned unknown error: %v", err)
	}

}
