package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetComponentByReleaseParams creates a new GetComponentByReleaseParams object
// with the default values initialized.
func NewGetComponentByReleaseParams() *GetComponentByReleaseParams {
	var ()
	return &GetComponentByReleaseParams{}
}

/*GetComponentByReleaseParams contains all the parameters to send to the API endpoint
for the get component by release operation typically these are written to a http.Request
*/
type GetComponentByReleaseParams struct {

	/*Component
	  A component is a single deis component, e.g., deis-router

	*/
	Component string
	/*Release
	  The release version of the deis component, eg., 2.0.0-beta2

	*/
	Release string
	/*Train
	  A train is a release cadence type, e.g., "beta" or "stable"

	*/
	Train string
}

// WithComponent adds the component to the get component by release params
func (o *GetComponentByReleaseParams) WithComponent(component string) *GetComponentByReleaseParams {
	o.Component = component
	return o
}

// WithRelease adds the release to the get component by release params
func (o *GetComponentByReleaseParams) WithRelease(release string) *GetComponentByReleaseParams {
	o.Release = release
	return o
}

// WithTrain adds the train to the get component by release params
func (o *GetComponentByReleaseParams) WithTrain(train string) *GetComponentByReleaseParams {
	o.Train = train
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentByReleaseParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param component
	if err := r.SetPathParam("component", o.Component); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	// path param train
	if err := r.SetPathParam("train", o.Train); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
