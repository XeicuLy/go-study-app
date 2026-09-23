package handler

import "net/http"

// TODO(human): NewRouter を実装する
//  1. http.NewServeMux() でルーターを作る
//  2. mux.HandleFunc("GET /healthz", Health) で GET /healthz に Health を紐付ける
//     ("GET /healthz" という書き方は Go 1.22 以降のメソッド指定ルーティング)
//  3. 作った mux を返す
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", Health)
	return mux
}
