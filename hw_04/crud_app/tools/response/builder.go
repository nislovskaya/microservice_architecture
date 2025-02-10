package response

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nislovskaya/microservice_architecture/hw_04/crud_app/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

type response struct {
	writer http.ResponseWriter
	logger *logrus.Entry
}

type Response interface {
	Ok(data interface{})
	Created(data interface{})
	NoContent()

	BadRequest(message string)
	NotFound(message string)
	InternalServerError(message string)
}

func New(w http.ResponseWriter, logger *logrus.Entry) Response {
	return &response{
		writer: w,
		logger: logger,
	}
}

func (r *response) Ok(data interface{}) {
	r.json(http.StatusOK, data)
}

func (r *response) Created(data interface{}) {
	r.json(http.StatusCreated, data)
}

func (r *response) NoContent() {
	r.json(http.StatusNoContent, nil)
}

func (r *response) BadRequest(message string) {
	r.json(http.StatusBadRequest, &model.Error{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func (r *response) NotFound(message string) {
	r.json(http.StatusNotFound, &model.Error{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func (r *response) InternalServerError(message string) {
	r.json(http.StatusInternalServerError, &model.Error{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func (r *response) json(statusCode int, data interface{}) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(statusCode)
	if err := jsoniter.NewEncoder(r.writer).Encode(data); err != nil {
		r.logger.WithError(err).Error("failed to encode json")
	}
}
