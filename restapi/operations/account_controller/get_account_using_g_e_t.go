// Code generated by go-swagger; DO NOT EDIT.

package account_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAccountUsingGETHandlerFunc turns a function with the right signature into a get account using g e t handler
type GetAccountUsingGETHandlerFunc func(GetAccountUsingGETParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountUsingGETHandlerFunc) Handle(params GetAccountUsingGETParams) middleware.Responder {
	return fn(params)
}

// GetAccountUsingGETHandler interface for that can handle valid get account using g e t params
type GetAccountUsingGETHandler interface {
	Handle(GetAccountUsingGETParams) middleware.Responder
}

// NewGetAccountUsingGET creates a new http.Handler for the get account using g e t operation
func NewGetAccountUsingGET(ctx *middleware.Context, handler GetAccountUsingGETHandler) *GetAccountUsingGET {
	return &GetAccountUsingGET{Context: ctx, Handler: handler}
}

/* GetAccountUsingGET swagger:route GET /accounts/{addressHash} account-controller getAccountUsingGET

Retrieve an account by its address hash.

*/
type GetAccountUsingGET struct {
	Context *middleware.Context
	Handler GetAccountUsingGETHandler
}

func (o *GetAccountUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAccountUsingGETParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
