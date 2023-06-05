package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alexcesaro/statsd"
	"github.com/building-microservices-with-go/chapter7/server/entities"
	"github.com/building-microservices-with-go/chapter7/server/httputil"
	"github.com/sirupsen/logrus"
)

type validationHandler struct {
	next   http.Handler
	statsd *statsd.Client
	logger *logrus.Logger
}

// NewValidationHandler creates a new Validation handler with the given statsD client and logger
func NewValidationHandler(statsd *statsd.Client, logger *logrus.Logger, next http.Handler) http.Handler {
	return &validationHandler{next: next, statsd: statsd, logger: logger}
}

func (h *validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request entities.HelloWorldRequest
	// read request param
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		// metric data
		h.statsd.Increment(validationFailed)

		// display request as a json
		message := httputil.SerialzableRequest{r}
		h.logger.WithFields(logrus.Fields{
			"handler": "Validation",
			"status":  http.StatusBadRequest,
			"method":  r.Method,
		}).Info(message.ToJSON())

		http.Error(rw, "Bad request", http.StatusBadRequest)

		return
	}

	// send arg as context to helloworld
	c := context.WithValue(r.Context(), "name", request.Name)
	r = r.WithContext(c)

	// metric data
	h.statsd.Increment(validationSuccess)

	//next chain handler
	h.next.ServeHTTP(rw, r)
}
