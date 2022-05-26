// Code generated by go-swagger; DO NOT EDIT.

package token_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindTokensUsingGETHandlerFunc turns a function with the right signature into a find tokens using g e t handler
type FindTokensUsingGETHandlerFunc func(FindTokensUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindTokensUsingGETHandlerFunc) Handle(params FindTokensUsingGETParams) middleware.Responder {
	return fn(params)
}

// FindTokensUsingGETHandler interface for that can handle valid find tokens using g e t params
type FindTokensUsingGETHandler interface {
	Handle(FindTokensUsingGETParams) middleware.Responder
}

// NewFindTokensUsingGET creates a new http.Handler for the find tokens using g e t operation
func NewFindTokensUsingGET(ctx *middleware.Context, handler FindTokensUsingGETHandler) *FindTokensUsingGET {
	return &FindTokensUsingGET{Context: ctx, Handler: handler}
}

/* FindTokensUsingGET swagger:route GET /tokens token-controller findTokensUsingGET

Retrieve token contracts.

*/
type FindTokensUsingGET struct {
	Context *middleware.Context
	Handler FindTokensUsingGETHandler
}

func (o *FindTokensUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindTokensUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
