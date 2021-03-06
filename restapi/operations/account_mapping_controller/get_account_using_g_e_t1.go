// Code generated by go-swagger; DO NOT EDIT.

package account_mapping_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAccountUsingGET1HandlerFunc turns a function with the right signature into a get account using g e t 1 handler
type GetAccountUsingGET1HandlerFunc func(GetAccountUsingGET1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountUsingGET1HandlerFunc) Handle(params GetAccountUsingGET1Params) middleware.Responder {
	return fn(params)
}

// GetAccountUsingGET1Handler interface for that can handle valid get account using g e t 1 params
type GetAccountUsingGET1Handler interface {
	Handle(GetAccountUsingGET1Params) middleware.Responder
}

// NewGetAccountUsingGET1 creates a new http.Handler for the get account using g e t 1 operation
func NewGetAccountUsingGET1(ctx *middleware.Context, handler GetAccountUsingGET1Handler) *GetAccountUsingGET1 {
	return &GetAccountUsingGET1{Context: ctx, Handler: handler}
}

/* GetAccountUsingGET1 swagger:route GET /mappings/{id} account-mapping-controller getAccountUsingGET1

getAccount

*/
type GetAccountUsingGET1 struct {
	Context *middleware.Context
	Handler GetAccountUsingGET1Handler
}

func (o *GetAccountUsingGET1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountUsingGET1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
