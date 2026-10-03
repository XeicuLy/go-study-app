package main

import (
	"log"
	"net/http"

	"github.com/XeicuLy/go-study-app/internal/handler"
	"github.com/XeicuLy/go-study-app/internal/model"
	"github.com/XeicuLy/go-study-app/internal/repository/memory"
	"github.com/XeicuLy/go-study-app/internal/service"
)

func main() {
	// 1. シード用の []model.Link を作る (2〜3件。例: Code "godev01" → "https://go.dev")
	seed := []model.Link{{Code: "godev01", OriginalURL: "https://go.dev"}, {Code: "godev02", OriginalURL: "https://go.dev/solutions/"}}
	// 2. memory.NewRepository(seed) で repository を作る
	repo := memory.NewRepository(seed)
	// 3. service.NewLinkService(repo) で service を作る
	svc := service.NewLinkService(repo)
	// 4. handler.NewLinkHandler(svc) で handler を作る
	h := handler.NewLinkHandler(svc)
	// 5. handler.NewRouter(h) でルーターを作る (今は引数なしになっているので、ここを直す)
	router := handler.NewRouter(h)
	// 6. http.ListenAndServe(":8080", router) でサーバーを起動する (ここは今のままでよい)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
