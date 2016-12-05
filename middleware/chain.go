package middleware

import (
	"net/http"
	"context"
)

type ContextFunc func(ctx context.Context, rw http.ResponseWriter, req *http.Request)

type Constructor func(ContextFunc) ContextFunc

func New(constructors ... Constructor) Chain {
	return Chain{append(([]Constructor)(nil), constructors...)}
}

type Chain struct {
	constructors []Constructor
}

func (c Chain) Then(next ContextFunc) http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		// Store the user IP in ctx for use by code in other packages.
		newNext := next
		for i := range c.constructors {
			newNext = c.constructors[len(c.constructors)-1-i](newNext)
		}

		ctx := context.TODO()
		newNext(ctx, rw, req)
	}

	return fn
}