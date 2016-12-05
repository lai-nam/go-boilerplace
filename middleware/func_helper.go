package middleware

import (
	"net/http"
	"github.com/meo/mcontext"
	"golang.org/x/blog/content/context/userip"
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/meo/util"
	"errors"
)

func Auth(next ContextFunc) ContextFunc {
	fn := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		newContext := mcontext.NewContext(ctx, req)

		if _, ok := mcontext.FromContext(newContext); !ok {
			ip, _ := userip.FromRequest(req)
			logrus.WithField("Ip", ip).Error("Access Forbiden")
			rw.WriteHeader(http.StatusForbidden)
			rw.Write([]byte("Access Forbiden"))
			return
		}

		next(newContext, rw, req)
	}

	return fn
}

func Recover(next ContextFunc) ContextFunc {
	fn := func(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error(err)
				util.WriteErrorHTTP(rw, errors.New("StatusInternalServerError"), http.StatusInternalServerError)
			}
		}()

		next(ctx, rw, req)
	}


	return fn
}