package middleware

import (
	"log"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 許可するオリジン
		log.Println("CORS middleware hit for:", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // React dev server
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // 認証付きの場合

		// プリフライトリクエスト（OPTIONS）はここで終わらせる
		if r.Method == "OPTIONS" {
			log.Println("OPTIONS preflight detected")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
