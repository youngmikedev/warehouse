// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/imranzahaev/warehouse/internal/delivery/swagger/gen/models"
)

// NewPutProductProductIDParams creates a new PutProductProductIDParams object
//
// There are no default values defined in the spec.
func NewPutProductProductIDParams() PutProductProductIDParams {

	return PutProductProductIDParams{}
}

// PutProductProductIDParams contains all the bound params for the put product product ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutProductProductID
type PutProductProductIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*update product data
	  Required: true
	  In: body
	*/
	Input *models.ProductUpdateProductRequest
	/*product id
	  Required: true
	  In: path
	*/
	ProductID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutProductProductIDParams() beforehand.
func (o *PutProductProductIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ProductUpdateProductRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("input", "body", ""))
			} else {
				res = append(res, errors.NewParseError("input", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Input = &body
			}
		}
	} else {
		res = append(res, errors.Required("input", "body", ""))
	}

	rProductID, rhkProductID, _ := route.Params.GetOK("productId")
	if err := o.bindProductID(rProductID, rhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProductID binds and validates parameter ProductID from path.
func (o *PutProductProductIDParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("productId", "path", "int64", raw)
	}
	o.ProductID = value

	return nil
}