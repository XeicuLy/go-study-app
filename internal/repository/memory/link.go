package memory

import (
	"sync"

	"github.com/XeicuLy/go-study-app/internal/model"
)

// Repository はメモリ上の map で Link を保持する。
//
//  1. type Repository struct { ... } で struct を定義する
//  2. フィールドは次の2つ
//     - mu    sync.RWMutex          map を同時に触らないためのロック (sync パッケージの import が必要)
//     - links map[string]model.Link キーは Code、値は Link (model パッケージの import が必要)
//  3. フィールド名を小文字で始めると、パッケージの外から触れない (非公開) になる。ここは非公開でよい
//
// 注意: sync.RWMutex を含む struct はコピーすると壊れる。
// メソッドのレシーバも、コンストラクタの戻り値も、必ずポインタ (*Repository) にする。
type Repository struct {
	mu    sync.RWMutex
	links map[string]model.Link
}

// NewRepository は seed の Link を登録した Repository を返す。
//
//  1. map[string]model.Link を make で作る (make せずに書き込むと panic する。宣言しただけの map は nil)
//  2. seed を for range で回し、各 Link を Code をキーにして map に入れる
//     (seed が nil でも for range は0回回るだけで、エラーにならない)
//  3. 作った Repository のポインタを返す: &Repository{ ... } の形で struct リテラルのアドレスを返せる
//     (Go に constructor は無く、New で始まる普通の関数でよい)
func NewRepository(seed []model.Link) *Repository {
	m := make(map[string]model.Link)
	for _, v := range seed {
		m[v.Code] = v
	}
	repo := &Repository{links: m}
	return repo
}

// FindByCode は code に対応する Link を返す。見つからなければ model.ErrNotFound を返す。
//
//  1. レシーバは (r *Repository) とポインタにする
//  2. 読み取りなので r.mu.RLock() でロックを取り、その直後に defer r.mu.RUnlock() を書く
//     (defer は「この関数を抜けるときに実行する」予約。途中で return しても解放漏れが起きない)
//  3. v, ok := r.links[code] の2値で受けると、キーが存在したかどうかが ok で分かる
//  4. ok が false なら、ゼロ値の model.Link{} と model.ErrNotFound を返す
//  5. ok が true なら、v と nil を返す
//
// 読み方は次のとおりです。
//
// - (r *Repository): このメソッドは Repository のもの。中では r と呼ぶ
// - FindByCode: メソッドの名前
// - (code string): 引数
// - (model.Link, error): 戻り値が2つ
//
// シグネチャは func (r *Repository) FindByCode(code string) (model.Link, error)
func (r *Repository) FindByCode(code string) (model.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.links[code]
	if !ok {
		return model.Link{}, model.ErrNotFound
	}
	return v, nil
}
