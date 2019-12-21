// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteProfileFavouriteHandlerFunc turns a function with the right signature into a delete profile favourite handler
type DeleteProfileFavouriteHandlerFunc func(DeleteProfileFavouriteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProfileFavouriteHandlerFunc) Handle(params DeleteProfileFavouriteParams) middleware.Responder {
	return fn(params)
}

// DeleteProfileFavouriteHandler interface for that can handle valid delete profile favourite params
type DeleteProfileFavouriteHandler interface {
	Handle(DeleteProfileFavouriteParams) middleware.Responder
}

// NewDeleteProfileFavourite creates a new http.Handler for the delete profile favourite operation
func NewDeleteProfileFavourite(ctx *middleware.Context, handler DeleteProfileFavouriteHandler) *DeleteProfileFavourite {
	return &DeleteProfileFavourite{Context: ctx, Handler: handler}
}

/*DeleteProfileFavourite swagger:route DELETE /profiles/{id}/favourite profile deleteProfileFavourite

delete profile from favourite list

*/
type DeleteProfileFavourite struct {
	Context *middleware.Context
	Handler DeleteProfileFavouriteHandler
}

func (o *DeleteProfileFavourite) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProfileFavouriteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}