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

//  1. httptest.NewRequest で POST /healthz のリクエストを作る
//  2. httptest.NewRecorder でレスポンスの書き込み先を作る
//  3. NewRouter().ServeHTTP(rec, req) でルーター経由で呼び出す
//  4. rec.Code が http.StatusMethodNotAllowed (405) かを確認し、
//     違っていたら t.Errorf で報告する
//     (TestHealth と同じ形の関数になるはず)
func TestHealthRejectsPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rec := httptest.NewRecorder()
	NewRouter().ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
