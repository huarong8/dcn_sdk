// Code generated by go-swagger; DO NOT EDIT.

package block_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindBlocksUsingGETHandlerFunc turns a function with the right signature into a find blocks using g e t handler
type FindBlocksUsingGETHandlerFunc func(FindBlocksUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindBlocksUsingGETHandlerFunc) Handle(params FindBlocksUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindBlocksUsingGETHandler interface for that can handle valid find blocks using g e t params
type FindBlocksUsingGETHandler interface {
	Handle(FindBlocksUsingGETParams) middleware.Responder
}

// NewFindBlocksUsingGET creates a new http.Handler for the find blocks using g e t operation
func NewFindBlocksUsingGET(ctx *middleware.Context, handler FindBlocksUsingGETHandler) *FindBlocksUsingGET {
	return &FindBlocksUsingGET{Context: ctx, Handler: handler}
}

/* FindBlocksUsingGET swagger:route GET /blocks block-controller findBlocksUsingGET

Retrieve blocks.

*/
type FindBlocksUsingGET struct {
	Context *middleware.Context
	Handler FindBlocksUsingGETHandler
}

func (o *FindBlocksUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindBlocksUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
