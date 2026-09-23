package handler

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	//  1. レスポンスの Content-Type やステータスコードは w.WriteHeader(http.StatusOK) で 200 を明示する
	//     (何もしなければデフォルトで200になるが、明示した方が意図が伝わる)
	w.WriteHeader(http.StatusOK)
	//  2. レスポンスボディとして "ok" という文字列を w.Write で書き込む
	//     (w.Write は []byte を受け取るので、文字列は []byte("ok") に変換する)
	w.Write([]byte("ok"))
}
