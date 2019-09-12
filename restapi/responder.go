// Copyright (C) 2017 ScyllaDB

package restapi

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
)

func responder(w http.ResponseWriter, r *http.Request, v interface{}) {
	err, ok := v.(error)

	// If not an error use DefaultResponder
	if !ok {
		render.DefaultResponder(w, r, v)
		return
	}

	herr, _ := v.(*httpError)
	if herr == nil {
		herr = &httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    errors.Wrap(err, "unexpected error, consult logs").Error(),
			TraceID:    log.TraceID(r.Context()),
		}
	}

	if le, _ := middleware.GetLogEntry(r).(*httpLogEntry); le != nil {
		le.AddError(err)
	}

	render.Status(r, herr.StatusCode)
	render.DefaultResponder(w, r, herr)
}
