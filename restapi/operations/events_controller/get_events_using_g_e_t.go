// Code generated by go-swagger; DO NOT EDIT.

package events_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEventsUsingGETHandlerFunc turns a function with the right signature into a get events using g e t handler
type GetEventsUsingGETHandlerFunc func(GetEventsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventsUsingGETHandlerFunc) Handle(params GetEventsUsingGETParams) middleware.Responder {
	return fn(params)
}

// GetEventsUsingGETHandler interface for that can handle valid get events using g e t params
type GetEventsUsingGETHandler interface {
	Handle(GetEventsUsingGETParams) middleware.Responder
}

// NewGetEventsUsingGET creates a new http.Handler for the get events using g e t operation
func NewGetEventsUsingGET(ctx *middleware.Context, handler GetEventsUsingGETHandler) *GetEventsUsingGET {
	return &GetEventsUsingGET{Context: ctx, Handler: handler}
}

/* GetEventsUsingGET swagger:route GET /events events-controller getEventsUsingGET

Retrieve all events.

*/
type GetEventsUsingGET struct {
	Context *middleware.Context
	Handler GetEventsUsingGETHandler
}

func (o *GetEventsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetEventsUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
