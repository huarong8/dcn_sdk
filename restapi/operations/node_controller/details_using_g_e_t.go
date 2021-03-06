// Code generated by go-swagger; DO NOT EDIT.

package node_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DetailsUsingGETHandlerFunc turns a function with the right signature into a details using g e t handler
type DetailsUsingGETHandlerFunc func(DetailsUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DetailsUsingGETHandlerFunc) Handle(params DetailsUsingGETParams) middleware.Responder {
	return fn(params)
}

// DetailsUsingGETHandler interface for that can handle valid details using g e t params
type DetailsUsingGETHandler interface {
	Handle(DetailsUsingGETParams) middleware.Responder
}

// NewDetailsUsingGET creates a new http.Handler for the details using g e t operation
func NewDetailsUsingGET(ctx *middleware.Context, handler DetailsUsingGETHandler) *DetailsUsingGET {
	return &DetailsUsingGET{Context: ctx, Handler: handler}
}

/* DetailsUsingGET swagger:route GET /node/details node-controller detailsUsingGET

Node details.

*/
type DetailsUsingGET struct {
	Context *middleware.Context
	Handler DetailsUsingGETHandler
}

func (o *DetailsUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDetailsUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
