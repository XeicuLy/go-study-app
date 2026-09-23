package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	// 1. httptest.NewRequest で GET /healthz のリクエストを作る
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	// 2. httptest.NewRecorder でレスポンスの書き込み先を作る
	rec := httptest.NewRecorder()
	// 3. NewRouter().ServeHTTP(rec, req) でルーター経由で呼び出す
	//    (Health を直接呼ばないこと。理由はIssue本文の「詰まりやすいポイント」参照)
	NewRouter().ServeHTTP(rec, req)
	// 4. rec.Code が 200 か、rec.Body.String() が "ok" かを確認し、
	//    違っていたら t.Errorf で報告する
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	expectBody := "ok"

	if rec.Body.String() != expectBody {
		t.Errorf("body = %s, want %s", rec.Body.String(), expectBody)
	}
}
