package helper

import (
	"context"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

// HTTPToContext returns an http RequestFunc that tries to join with an
// OpenTracing trace found in `req` and starts a new Span called
// `operationName` accordingly. If no trace could be found in `req`, the Span
// will be a trace root. The Span is incorporated in the returned Context and
// can be retrieved with opentracing.SpanFromContext(ctx).
func HTTPToContext(tracer stdopentracing.Tracer, operationName string, logger log.Logger) kithttp.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		// Try to join to a trace propagated in `req`.
		var span stdopentracing.Span
		wireContext, err := tracer.Extract(
			stdopentracing.HTTPHeaders,
			stdopentracing.HTTPHeadersCarrier(req.Header),
		)
		if err != nil && err != stdopentracing.ErrSpanContextNotFound {
			logger.Log("err", err)
		}

		span = tracer.StartSpan(operationName, ext.RPCServerOption(wireContext))
		defer span.Finish()

		ext.HTTPMethod.Set(span, req.Method)
		ext.HTTPUrl.Set(span, req.URL.String())
		return stdopentracing.ContextWithSpan(ctx, span)
	}
}
