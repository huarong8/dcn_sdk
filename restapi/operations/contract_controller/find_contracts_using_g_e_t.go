// Code generated by go-swagger; DO NOT EDIT.

package contract_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindContractsUsingGETHandlerFunc turns a function with the right signature into a find contracts using g e t handler
type FindContractsUsingGETHandlerFunc func(FindContractsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindContractsUsingGETHandlerFunc) Handle(params FindContractsUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindContractsUsingGETHandler interface for that can handle valid find contracts using g e t params
type FindContractsUsingGETHandler interface {
	Handle(FindContractsUsingGETParams) middleware.Responder
}

// NewFindContractsUsingGET creates a new http.Handler for the find contracts using g e t operation
func NewFindContractsUsingGET(ctx *middleware.Context, handler FindContractsUsingGETHandler) *FindContractsUsingGET {
	return &FindContractsUsingGET{Context: ctx, Handler: handler}
}

/* FindContractsUsingGET swagger:route GET /contracts contract-controller findContractsUsingGET

Retrieve contracts.

*/
type FindContractsUsingGET struct {
	Context *middleware.Context
	Handler FindContractsUsingGETHandler
}

func (o *FindContractsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindContractsUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
