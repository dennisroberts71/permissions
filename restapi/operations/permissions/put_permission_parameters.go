// Code generated by go-swagger; DO NOT EDIT.

package permissions

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
	"github.com/go-openapi/validate"

	"github.com/cyverse-de/permissions/models"
)

// NewPutPermissionParams creates a new PutPermissionParams object
//
// There are no default values defined in the spec.
func NewPutPermissionParams() PutPermissionParams {

	return PutPermissionParams{}
}

// PutPermissionParams contains all the bound params for the put permission operation
// typically these are obtained from a http.Request
//
// swagger:parameters putPermission
type PutPermissionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The permission level to assign.
	  Required: true
	  In: body
	*/
	Permission *models.PermissionPutRequest
	/*The resource name.
	  Required: true
	  In: path
	*/
	ResourceName string
	/*The resource type name.
	  Required: true
	  In: path
	*/
	ResourceType string
	/*The external subject identifier.
	  Required: true
	  In: path
	*/
	SubjectID string
	/*The subject type name.
	  Required: true
	  In: path
	*/
	SubjectType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutPermissionParams() beforehand.
func (o *PutPermissionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PermissionPutRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("permission", "body", ""))
			} else {
				res = append(res, errors.NewParseError("permission", "body", "", err))
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
				o.Permission = &body
			}
		}
	} else {
		res = append(res, errors.Required("permission", "body", ""))
	}

	rResourceName, rhkResourceName, _ := route.Params.GetOK("resource_name")
	if err := o.bindResourceName(rResourceName, rhkResourceName, route.Formats); err != nil {
		res = append(res, err)
	}

	rResourceType, rhkResourceType, _ := route.Params.GetOK("resource_type")
	if err := o.bindResourceType(rResourceType, rhkResourceType, route.Formats); err != nil {
		res = append(res, err)
	}

	rSubjectID, rhkSubjectID, _ := route.Params.GetOK("subject_id")
	if err := o.bindSubjectID(rSubjectID, rhkSubjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSubjectType, rhkSubjectType, _ := route.Params.GetOK("subject_type")
	if err := o.bindSubjectType(rSubjectType, rhkSubjectType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindResourceName binds and validates parameter ResourceName from path.
func (o *PutPermissionParams) bindResourceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ResourceName = raw

	return nil
}

// bindResourceType binds and validates parameter ResourceType from path.
func (o *PutPermissionParams) bindResourceType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ResourceType = raw

	return nil
}

// bindSubjectID binds and validates parameter SubjectID from path.
func (o *PutPermissionParams) bindSubjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SubjectID = raw

	return nil
}

// bindSubjectType binds and validates parameter SubjectType from path.
func (o *PutPermissionParams) bindSubjectType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.SubjectType = raw

	if err := o.validateSubjectType(formats); err != nil {
		return err
	}

	return nil
}

// validateSubjectType carries on validations for parameter SubjectType
func (o *PutPermissionParams) validateSubjectType(formats strfmt.Registry) error {

	if err := validate.EnumCase("subject_type", "path", o.SubjectType, []interface{}{"user", "group"}, true); err != nil {
		return err
	}

	return nil
}
