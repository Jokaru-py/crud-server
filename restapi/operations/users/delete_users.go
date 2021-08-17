// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteUsersHandlerFunc turns a function with the right signature into a delete users handler
type DeleteUsersHandlerFunc func(DeleteUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUsersHandlerFunc) Handle(params DeleteUsersParams) middleware.Responder {
	return fn(params)
}

// DeleteUsersHandler interface for that can handle valid delete users params
type DeleteUsersHandler interface {
	Handle(DeleteUsersParams) middleware.Responder
}

// NewDeleteUsers creates a new http.Handler for the delete users operation
func NewDeleteUsers(ctx *middleware.Context, handler DeleteUsersHandler) *DeleteUsers {
	return &DeleteUsers{Context: ctx, Handler: handler}
}

/* DeleteUsers swagger:route DELETE /users Users deleteUsers

Удалить пользователя.

*/
type DeleteUsers struct {
	Context *middleware.Context
	Handler DeleteUsersHandler
}

func (o *DeleteUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteUsersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
