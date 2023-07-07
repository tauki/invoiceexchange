package logging

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type zapKey struct{}

func WithContext(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, zapKey{}, log)
}

func FromContext(ctx context.Context) (*zap.Logger, bool) {
	log, ok := ctx.Value(zapKey{}).(*zap.Logger)
	return log, ok
}

func GetZap(ctx context.Context) *zap.Logger {
	log, ok := FromContext(ctx)
	if !ok {
		log = zap.L()
		log.Warn("zap logger not found on Context, using zap.L()")
	}
	return log
}

//func ZapMiddleware(log *zap.Logger, next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//
//		log2 := new(zap.Logger)
//		*log2 = *log
//
//		ctx = WithContext(ctx, log2)
//		r = r.WithContext(ctx)
//		next.ServeHTTP(w, r)
//	})
//}

func ZapMiddleware(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			log2 := new(zap.Logger)
			*log2 = *log

			ctx = WithContext(ctx, log2)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
