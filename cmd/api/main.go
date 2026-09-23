package main

import (
	"log"
	"net/http"

	"github.com/XeicuLy/go-study-app/internal/handler"
)

//  1. handler.NewRouter() でルーターを作る
//  2. http.ListenAndServe(":8080", router) でポート8080でサーバーを起動する
//     (ListenAndServe はエラーを返すので、log.Fatal でラップしてプロセスを終了させる)
func main() {
	router := handler.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
