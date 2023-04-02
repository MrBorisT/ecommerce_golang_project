package tracing

import (
	"net/http"

	"route256/libs/logger"
	"route256/libs/reswrapper"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

func Init(serviceName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}

	_, err := cfg.InitGlobalTracer(serviceName)
	if err != nil {
		logger.Fatal("Cannot init tracing", zap.Error(err))
	}
}

func Middleware(next http.Handler, handlerName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		span, ctx := opentracing.StartSpanFromContext(ctx, handlerName)
		defer span.Finish()

		span.SetTag("url", r.URL.String())

		if spanContext, ok := span.Context().(jaeger.SpanContext); ok {
			w.Header().Add("x-trace-id", spanContext.TraceID().String())
		}

		r = r.WithContext(ctx)
		wrapper := reswrapper.NewResponseWrapper(w)
		next.ServeHTTP(wrapper, r)

		if wrapper.StatusCode != http.StatusOK {
			ext.Error.Set(span, true)
		}
		span.SetTag("status_code", http.StatusText(wrapper.StatusCode))
	})
}
