// Code generated by go-swagger; DO NOT EDIT.

package dashboard_transactions_totals_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTotalsUsingGET1HandlerFunc turns a function with the right signature into a get totals using g e t 1 handler
type GetTotalsUsingGET1HandlerFunc func(GetTotalsUsingGET1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTotalsUsingGET1HandlerFunc) Handle(params GetTotalsUsingGET1Params) middleware.Responder {
	return fn(params)
}

// GetTotalsUsingGET1Handler interface for that can handle valid get totals using g e t 1 params
type GetTotalsUsingGET1Handler interface {
	Handle(GetTotalsUsingGET1Params) middleware.Responder
}

// NewGetTotalsUsingGET1 creates a new http.Handler for the get totals using g e t 1 operation
func NewGetTotalsUsingGET1(ctx *middleware.Context, handler GetTotalsUsingGET1Handler) *GetTotalsUsingGET1 {
	return &GetTotalsUsingGET1{Context: ctx, Handler: handler}
}

/* GetTotalsUsingGET1 swagger:route GET /dashboard/transactions/totals dashboard-transactions-totals-controller getTotalsUsingGET1

Retrieve total number of different transactions types

*/
type GetTotalsUsingGET1 struct {
	Context *middleware.Context
	Handler GetTotalsUsingGET1Handler
}

func (o *GetTotalsUsingGET1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTotalsUsingGET1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
