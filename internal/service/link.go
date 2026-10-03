package service

import (
	"github.com/XeicuLy/go-study-app/internal/model"
	"github.com/XeicuLy/go-study-app/internal/repository/memory"
)

// LinkService は Link に関するアプリのルールを担う。今は repository を呼ぶだけでよい。
//
//  1. type LinkService struct { ... } で struct を定義する
//  2. フィールドは repo 1つ。型は *memory.Repository (ポインタ)
//     (repository/memory を import する。interface にするのは #8 なので、今は具象型のまま)
//  3. net/http は import しない (NM4)。HTTP の話は handler だけが知っている
type LinkService struct {
	repo *memory.Repository
}

// NewLinkService は repo を持つ LinkService を返す。
//
//  1. シグネチャは func NewLinkService(repo *memory.Repository) *LinkService
//  2. &LinkService{ ... } で struct リテラルのポインタを作って返す
//     (NewRepository と同じ形。フィールド名を付けて repo を渡す)
func NewLinkService(repo *memory.Repository) *LinkService {
	svc := &LinkService{repo: repo}
	return svc
}

// Get は code に対応する Link を返す。見つからなければ model.ErrNotFound を返す。
//
//  1. レシーバは (s *LinkService)
//  2. シグネチャは func (s *LinkService) Get(code string) (model.Link, error)
//  3. 中身は s.repo.FindByCode(code) の結果をそのまま返すだけ
//     (戻り値が2つなので、そのまま return s.repo.FindByCode(code) と書ける)
func (s *LinkService) Get(code string) (model.Link, error) {
	return s.repo.FindByCode(code)
}
