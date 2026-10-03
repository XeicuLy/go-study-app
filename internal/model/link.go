package model

import "time"

// Link は短縮URL1件ぶんのデータ。
//
//  1. type Link struct { ... } で struct を定義する
//  2. フィールドは次の3つ (Issue本文の「決めてあること」参照)
//     - Code        string    短縮コード
//     - OriginalURL string    リダイレクト先の元のURL
//     - CreatedAt   time.Time 作成日時 (time パッケージの import が必要)
//  3. struct にメソッドは書かない。データの定義だけでよい
type Link struct {
	Code        string
	OriginalURL string
	CreatedAt   time.Time
}
