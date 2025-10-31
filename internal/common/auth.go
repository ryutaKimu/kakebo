package common

import (
	"context"
	"errors"
	"net/http"
)

type userIDKey struct{}

var ErrUnauthorized = errors.New("unauthorized")

func SetUserID(req *http.Request, userID int) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userIDKey{}, userID)
	return req.WithContext(ctx)
}

func GetCurrentUserID(ctx context.Context) (int, error) {
userID, ok := ctx.Value(userIDKey{}).(int)
if !ok {
	return 0, ErrUnauthorized
}

	return userID, nil
}
