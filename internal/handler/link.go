package handler

import (
	"errors"
	"net/http"

	"github.com/XeicuLy/go-study-app/internal/model"
	"github.com/XeicuLy/go-study-app/internal/service"
)

// LinkHandler は Link に関する HTTP リクエストを受け取る。
//
// type LinkHandler struct { ... }
// - フィールドは svc 1つ。型は *service.LinkService (service パッケージを import する)
type LinkHandler struct {
	svc *service.LinkService
}

// NewLinkHandler は svc を持つ LinkHandler を返す。
//
// func NewLinkHandler(svc *service.LinkService) *LinkHandler
// - NewLinkService と同じ形。&LinkHandler{ svc: svc } を返す
func NewLinkHandler(svc *service.LinkService) *LinkHandler {
	h := &LinkHandler{svc: svc}
	return h
}

// Redirect は GET /{code} を処理する。
//
// func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
//
//  1. r.PathValue("code") で URL の {code} 部分を取り出す
//  2. h.svc.Get(code) を呼んで、link と err を受け取る
//  3. err が model.ErrNotFound なら、http.Error(w, "not found", http.StatusNotFound) を呼んで return する
//     (判定は errors.Is(err, model.ErrNotFound)。「ErrNotFound だったら」 なので ! は付けない)
//  4. err が nil でも ErrNotFound でもない場合は、http.Error(w, "internal server error", http.StatusInternalServerError) で 500 を返して return する
//  5. 成功したら http.Redirect(w, r, link.OriginalURL, http.StatusFound) を呼ぶ
//
// }
func (h *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	link, err := h.svc.Get(code)
	if errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}
