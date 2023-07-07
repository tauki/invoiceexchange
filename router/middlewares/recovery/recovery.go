package recovery

import (
	"github.com/tauki/invoiceexchange/handlers"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/router/middlewares/logging"
	"go.uber.org/zap"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				ctx := r.Context()
				log := logging.GetZap(ctx)

				log.Error("recovered", zap.Any("error", err))

				handlers.HTTPErrorResponse(w,
					errors.NewInfrastructureError("server encountered a problem", nil))
				return
			}
		}()

		next.ServeHTTP(w, r)

	})
}
