package model

import "errors"

// ErrNotFound は「指定したコードのLinkが存在しない」ことを表すセンチネルエラー。
//
//  1. errors パッケージを import する
//  2. パッケージレベルの変数として var ErrNotFound = errors.New("link not found") と宣言する
//     (関数の中ではなくパッケージ直下に1つだけ作る。呼び出し側は errors.Is(err, model.ErrNotFound) で比較する)
var ErrNotFound = errors.New("link not found")
