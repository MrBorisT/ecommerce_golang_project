package metrics

import (
	"net/http"
	"route256/libs/reswrapper"
	"time"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RequestsCounter.Inc()

		timeStart := time.Now()

		wrapper := reswrapper.NewResponseWrapper(w)
		next.ServeHTTP(wrapper, r)

		elapsed := time.Since(timeStart)

		HistogramResponseTime.WithLabelValues(http.StatusText(wrapper.StatusCode)).Observe(elapsed.Seconds())
		ResponseCounter.WithLabelValues(http.StatusText(wrapper.StatusCode)).Inc()
	})
}
