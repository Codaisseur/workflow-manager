package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetComponentsByLatestReleaseForV2Params creates a new GetComponentsByLatestReleaseForV2Params object
// with the default values initialized.
func NewGetComponentsByLatestReleaseForV2Params() *GetComponentsByLatestReleaseForV2Params {
	var ()
	return &GetComponentsByLatestReleaseForV2Params{}
}

/*GetComponentsByLatestReleaseForV2Params contains all the parameters to send to the API endpoint
for the get components by latest release for v2 operation typically these are written to a http.Request
*/
type GetComponentsByLatestReleaseForV2Params struct {

	/*Body*/
	Body GetComponentsByLatestReleaseForV2Body
}

// WithBody adds the body to the get components by latest release for v2 params
func (o *GetComponentsByLatestReleaseForV2Params) WithBody(body GetComponentsByLatestReleaseForV2Body) *GetComponentsByLatestReleaseForV2Params {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentsByLatestReleaseForV2Params) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}