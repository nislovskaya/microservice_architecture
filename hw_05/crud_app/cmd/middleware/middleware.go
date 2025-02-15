package middleware

//
//import (
//	"github.com/gorilla/mux"
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/promauto"
//	"net/http"
//	"strconv"
//)
//
//var (
//	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
//		Name:    "http_request_duration_seconds",
//		Help:    "Duration of HTTP requests.",
//		Buckets: []float64{0.1, 0.5, 1.0, 2.0, 5.0},
//	}, []string{"method", "path", "status"})
//
//	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Name: "http_requests_total",
//		Help: "Total number of HTTP requests",
//	}, []string{"method", "path", "status"})
//
//	httpResponseErrorsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Name: "http_response_errors_total",
//		Help: "Total number of HTTP 5xx responses",
//	}, []string{"method", "path"})
//)
//
//func init() {
//	prometheus.MustRegister(httpRequestDuration)
//	prometheus.MustRegister(httpRequestsTotal)
//	prometheus.MustRegister(httpResponseErrorsTotal)
//
//	prometheus.Register(totalRequests)
//	prometheus.Register(responseStatus)
//	prometheus.Register(httpDuration)
//}
//
//type responseWriter struct {
//	http.ResponseWriter
//	statusCode int
//}
//
//func NewResponseWriter(w http.ResponseWriter) *responseWriter {
//	return &responseWriter{w, http.StatusOK}
//}
//
//func (rw *responseWriter) WriteHeader(code int) {
//	rw.statusCode = code
//	rw.ResponseWriter.WriteHeader(code)
//}
//
//var totalRequests = prometheus.NewCounterVec(
//	prometheus.CounterOpts{
//		Name: "http_requests_total",
//		Help: "Number of get requests.",
//	},
//	[]string{"path"},
//)
//
//var responseStatus = prometheus.NewCounterVec(
//	prometheus.CounterOpts{
//		Name: "response_status",
//		Help: "Status of HTTP response",
//	},
//	[]string{"status"},
//)
//
//var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
//	Name: "http_response_time_seconds",
//	Help: "Duration of HTTP requests.",
//}, []string{"path"})
//
//func PrometheusMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		route := mux.CurrentRoute(r)
//		path, _ := route.GetPathTemplate()
//
//		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
//		rw := NewResponseWriter(w)
//		next.ServeHTTP(rw, r)
//
//		statusCode := rw.statusCode
//
//		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
//		totalRequests.WithLabelValues(path).Inc()
//
//		timer.ObserveDuration()
//	})
//}
