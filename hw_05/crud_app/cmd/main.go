package main

import (
	"github.com/gorilla/mux"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/config"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/handler"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/repository"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/service/user"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

var logger = logrus.NewEntry(logrus.New())

var (
	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: []float64{0.1, 0.5, 1.0, 2.0, 5.0},
	}, []string{"method", "path", "status"})

	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "path", "status"})

	httpResponseErrorsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_response_errors_total",
		Help: "Total number of HTTP 5xx responses",
	}, []string{"method", "path"})
)

func main() {
	db := config.ConnectDB(logger)

	router := getRouter(db)

	router.Use(prometheusMiddleware)
	router.Handle("/metrics", promhttp.Handler())

	logger.Info("Server is started...")

	logger.Fatal(http.ListenAndServe(":8080", router))
}

func getRouter(db *gorm.DB) *mux.Router {
	userRepository := repository.New(
		repository.WithLogger(logger),
		repository.WithDB(db),
	)

	userService := user.New(
		user.WithLogger(logger),
		user.WithRepository(userRepository),
	)

	services := service.New(
		service.WithUserService(userService),
	)

	handlers := handler.New(
		handler.WithLogger(logger),
		handler.WithService(services),
	)

	return handlers.InitRouter()
}

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedResponseWriter := wrapResponseWriter(w)

		next.ServeHTTP(wrappedResponseWriter, r)

		statusCode := strconv.Itoa(wrappedResponseWriter.Status())
		duration := time.Since(start)

		httpRequestDuration.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path, "status": statusCode}).Observe(duration.Seconds())
		httpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path, "status": statusCode}).Inc()
		if statusCode[0] == '5' {
			httpResponseErrorsTotal.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).Inc()
		}
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriterWrapper) Status() int {
	return rw.statusCode
}
