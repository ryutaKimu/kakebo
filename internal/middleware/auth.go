package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ryutaKimu/kakebo/internal/common"
	"github.com/ryutaKimu/kakebo/internal/pkg/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

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
