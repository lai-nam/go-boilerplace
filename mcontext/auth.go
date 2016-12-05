package mcontext

import (
	"net/http"
	"context"
)

type key int

const userKey key = 0

func NewContext(ctx context.Context, req *http.Request) context.Context {
	token := req.Header.Get("X-Auth-Token")
	if token == "" {
		return ctx
	}

	payload, err := ParseToken(token)
	if err != nil {
		return ctx
	}

	return context.WithValue(ctx, userKey, payload)
}

func FromContext(ctx context.Context) (*UserClaim, bool) {
	userClaim, ok := ctx.Value(userKey).(*UserClaim)
	return userClaim, ok
}


