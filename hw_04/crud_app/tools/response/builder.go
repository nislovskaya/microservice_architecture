package response

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"net/http"
)

type response struct {
	writer     http.ResponseWriter
	logger     *logrus.Entry
	statusCode int
}

type Builder interface {
	Ok() Builder
	Created() Builder
	BadRequest() Builder
	InternalServerError() Builder

	JSON(data interface{})
}

func New(w http.ResponseWriter, logger *logrus.Entry) Builder {
	return &response{
		writer: w,
		logger: logger,
	}
}

func (r *response) Ok() Builder {
	r.statusCode = http.StatusOK
	return r
}

func (r *response) Created() Builder {
	r.statusCode = http.StatusCreated
	return r
}

func (r *response) BadRequest() Builder {
	r.statusCode = http.StatusBadRequest
	return r
}

func (r *response) InternalServerError() Builder {
	r.statusCode = http.StatusInternalServerError
	return r
}

func (r *response) JSON(data interface{}) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(r.statusCode)
	err := jsoniter.NewEncoder(r.writer).Encode(data)
	if err != nil {
		r.logger.WithError(err).Error("failed to encode json")
	}
	r.logger.Info("JSON data sent successfully")
}
