package common

import (
	"context"
	"errors"
	"net/http"
)

type userIDKey struct{}

func SetUserID(req *http.Request, userID int) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userIDKey{}, userID)
	return req.WithContext(ctx)
}

func GetCurrentUserID(ctx context.Context) (int, error) {
	userIDInterface := ctx.Value(userIDKey{})
	if userIDInterface == nil {
		return 0, errors.New("Unauthorized")
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		return 0, errors.New("Unauthorized")
	}

	return userID, nil
}
