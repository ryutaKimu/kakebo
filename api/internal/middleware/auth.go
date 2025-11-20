package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ryutaKimu/kakebo/api/internal/common"
	"github.com/ryutaKimu/kakebo/api/internal/pkg/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := cookie.Value

		claims, err := jwt.NewJWT().VerifyToken(token)
		if err != nil {
			log.Printf("failed to verify token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, err := strconv.Atoi(claims.UserID)
		if err != nil {
			log.Printf("failed to parse userID from token: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r = common.SetUserID(r, userID)

		next.ServeHTTP(w, r)
	})
}
