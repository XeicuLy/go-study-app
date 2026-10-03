package memory

import (
	"errors"
	"testing"

	"github.com/XeicuLy/go-study-app/internal/model"
)

//  1. シードに入れたコードを FindByCode すると、そのLinkが返る
//     a. model.Link のスライスを作る: []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}}
//     b. NewRepository(seed) で Repository を作る
//     c. repo.FindByCode("godev01") を呼び、err が nil であること、
//     返ってきた Link の OriginalURL が期待どおりであることを確認する
//     (違っていたら t.Errorf / 予期しないerrなら t.Fatalf で報告する)
func TestFindByCodeFound(t *testing.T) {
	seed := []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}}
	repo := NewRepository(seed)
	result, err := repo.FindByCode("godev01")
	if err != nil {
		t.Fatalf("FindByCode returned an error: %v", err)
	}
	want := "https://go.dev"
	if result.OriginalURL != want {
		t.Errorf("result = %v, want = %v", result.OriginalURL, want)
	}
}

//  2. 無いコードだと model.ErrNotFound が返る
//     a. NewRepository(nil) で空の Repository を作る (nil スライスで動く理由は Issue本文の末尾を参照)
//     b. repo.FindByCode("nosuchcode") を呼ぶ
//     c. errors.Is(err, model.ErrNotFound) で判定する。err != model.ErrNotFound とは書かない
//     (#4 でエラーをラップすると != では判定できなくなるため、最初から errors.Is)
func TestFindByCodeNotFound(t *testing.T) {
	repo := NewRepository(nil)
	_, err := repo.FindByCode("nosuchcode")
	if !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("FindByCode returned Link Not Found: %v", err)
	}
}
