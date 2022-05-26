// Code generated by go-swagger; DO NOT EDIT.

package token_events_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTokenEventsUsingGETHandlerFunc turns a function with the right signature into a get token events using g e t handler
type GetTokenEventsUsingGETHandlerFunc func(GetTokenEventsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTokenEventsUsingGETHandlerFunc) Handle(params GetTokenEventsUsingGETParams) middleware.Responder {
	return fn(params)
}

// GetTokenEventsUsingGETHandler interface for that can handle valid get token events using g e t params
type GetTokenEventsUsingGETHandler interface {
	Handle(GetTokenEventsUsingGETParams) middleware.Responder
}

// NewGetTokenEventsUsingGET creates a new http.Handler for the get token events using g e t operation
func NewGetTokenEventsUsingGET(ctx *middleware.Context, handler GetTokenEventsUsingGETHandler) *GetTokenEventsUsingGET {
	return &GetTokenEventsUsingGET{Context: ctx, Handler: handler}
}

/* GetTokenEventsUsingGET swagger:route GET /tokens/{addressHash}/events token-events-controller getTokenEventsUsingGET

Retrieve all token events using token address.

*/
type GetTokenEventsUsingGET struct {
	Context *middleware.Context
	Handler GetTokenEventsUsingGETHandler
}

func (o *GetTokenEventsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTokenEventsUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
