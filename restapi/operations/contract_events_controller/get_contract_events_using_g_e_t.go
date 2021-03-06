// Code generated by go-swagger; DO NOT EDIT.

package contract_events_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetContractEventsUsingGETHandlerFunc turns a function with the right signature into a get contract events using g e t handler
type GetContractEventsUsingGETHandlerFunc func(GetContractEventsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContractEventsUsingGETHandlerFunc) Handle(params GetContractEventsUsingGETParams) middleware.Responder {
	return fn(params)
}

// GetContractEventsUsingGETHandler interface for that can handle valid get contract events using g e t params
type GetContractEventsUsingGETHandler interface {
	Handle(GetContractEventsUsingGETParams) middleware.Responder
}

// NewGetContractEventsUsingGET creates a new http.Handler for the get contract events using g e t operation
func NewGetContractEventsUsingGET(ctx *middleware.Context, handler GetContractEventsUsingGETHandler) *GetContractEventsUsingGET {
	return &GetContractEventsUsingGET{Context: ctx, Handler: handler}
}

/* GetContractEventsUsingGET swagger:route GET /contracts/{addressHash}/events contract-events-controller getContractEventsUsingGET

Retrieve all events for this contract.

*/
type GetContractEventsUsingGET struct {
	Context *middleware.Context
	Handler GetContractEventsUsingGETHandler
}

func (o *GetContractEventsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetContractEventsUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
