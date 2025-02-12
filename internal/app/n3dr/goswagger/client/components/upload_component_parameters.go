// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUploadComponentParams creates a new UploadComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadComponentParams() *UploadComponentParams {
	return &UploadComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadComponentParamsWithTimeout creates a new UploadComponentParams object
// with the ability to set a timeout on a request.
func NewUploadComponentParamsWithTimeout(timeout time.Duration) *UploadComponentParams {
	return &UploadComponentParams{
		timeout: timeout,
	}
}

// NewUploadComponentParamsWithContext creates a new UploadComponentParams object
// with the ability to set a context for a request.
func NewUploadComponentParamsWithContext(ctx context.Context) *UploadComponentParams {
	return &UploadComponentParams{
		Context: ctx,
	}
}

// NewUploadComponentParamsWithHTTPClient creates a new UploadComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadComponentParamsWithHTTPClient(client *http.Client) *UploadComponentParams {
	return &UploadComponentParams{
		HTTPClient: client,
	}
}

/* UploadComponentParams contains all the parameters to send to the API endpoint
   for the upload component operation.

   Typically these are written to a http.Request.
*/
type UploadComponentParams struct {

	/* AptAsset.

	   apt Asset
	*/
	AptAsset runtime.NamedReadCloser

	/* DockerAsset.

	   docker Asset
	*/
	DockerAsset runtime.NamedReadCloser

	/* HelmAsset.

	   helm Asset
	*/
	HelmAsset runtime.NamedReadCloser

	/* Maven2ArtifactID.

	   maven2 Artifact ID
	*/
	Maven2ArtifactID *string

	/* Maven2Asset1.

	   maven2 Asset 1
	*/
	Maven2Asset1 runtime.NamedReadCloser

	/* Maven2Asset1Classifier.

	   maven2 Asset 1 Classifier
	*/
	Maven2Asset1Classifier *string

	/* Maven2Asset1Extension.

	   maven2 Asset 1 Extension
	*/
	Maven2Asset1Extension *string

	/* Maven2Asset2.

	   maven2 Asset 2
	*/
	Maven2Asset2 runtime.NamedReadCloser

	/* Maven2Asset2Classifier.

	   maven2 Asset 2 Classifier
	*/
	Maven2Asset2Classifier *string

	/* Maven2Asset2Extension.

	   maven2 Asset 2 Extension
	*/
	Maven2Asset2Extension *string

	/* Maven2Asset3.

	   maven2 Asset 3
	*/
	Maven2Asset3 runtime.NamedReadCloser

	/* Maven2Asset3Classifier.

	   maven2 Asset 3 Classifier
	*/
	Maven2Asset3Classifier *string

	/* Maven2Asset3Extension.

	   maven2 Asset 3 Extension
	*/
	Maven2Asset3Extension *string

	/* Maven2GeneratePom.

	   maven2 Generate a POM file with these coordinates
	*/
	Maven2GeneratePom *bool

	/* Maven2GroupID.

	   maven2 Group ID
	*/
	Maven2GroupID *string

	/* Maven2Packaging.

	   maven2 Packaging
	*/
	Maven2Packaging *string

	/* Maven2Version.

	   maven2 Version
	*/
	Maven2Version *string

	/* NpmAsset.

	   npm Asset
	*/
	NpmAsset runtime.NamedReadCloser

	/* NugetAsset.

	   nuget Asset
	*/
	NugetAsset runtime.NamedReadCloser

	/* PypiAsset.

	   pypi Asset
	*/
	PypiAsset runtime.NamedReadCloser

	/* RAsset.

	   r Asset
	*/
	RAsset runtime.NamedReadCloser

	/* RAssetPathID.

	   r Asset  Package Path
	*/
	RAssetPathID *string

	/* RawAsset1.

	   raw Asset 1
	*/
	RawAsset1 runtime.NamedReadCloser

	/* RawAsset1Filename.

	   raw Asset 1 Filename
	*/
	RawAsset1Filename *string

	/* RawAsset2.

	   raw Asset 2
	*/
	RawAsset2 runtime.NamedReadCloser

	/* RawAsset2Filename.

	   raw Asset 2 Filename
	*/
	RawAsset2Filename *string

	/* RawAsset3.

	   raw Asset 3
	*/
	RawAsset3 runtime.NamedReadCloser

	/* RawAsset3Filename.

	   raw Asset 3 Filename
	*/
	RawAsset3Filename *string

	/* RawDirectory.

	   raw Directory
	*/
	RawDirectory *string

	/* Repository.

	   Name of the repository to which you would like to upload the component
	*/
	Repository string

	/* RubygemsAsset.

	   rubygems Asset
	*/
	RubygemsAsset runtime.NamedReadCloser

	/* YumAsset.

	   yum Asset
	*/
	YumAsset runtime.NamedReadCloser

	/* YumAssetFilename.

	   yum Asset  Filename
	*/
	YumAssetFilename *string

	/* YumDirectory.

	   yum Directory
	*/
	YumDirectory *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadComponentParams) WithDefaults() *UploadComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload component params
func (o *UploadComponentParams) WithTimeout(timeout time.Duration) *UploadComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload component params
func (o *UploadComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload component params
func (o *UploadComponentParams) WithContext(ctx context.Context) *UploadComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload component params
func (o *UploadComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload component params
func (o *UploadComponentParams) WithHTTPClient(client *http.Client) *UploadComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload component params
func (o *UploadComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAptAsset adds the aptAsset to the upload component params
func (o *UploadComponentParams) WithAptAsset(aptAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetAptAsset(aptAsset)
	return o
}

// SetAptAsset adds the aptAsset to the upload component params
func (o *UploadComponentParams) SetAptAsset(aptAsset runtime.NamedReadCloser) {
	o.AptAsset = aptAsset
}

// WithDockerAsset adds the dockerAsset to the upload component params
func (o *UploadComponentParams) WithDockerAsset(dockerAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetDockerAsset(dockerAsset)
	return o
}

// SetDockerAsset adds the dockerAsset to the upload component params
func (o *UploadComponentParams) SetDockerAsset(dockerAsset runtime.NamedReadCloser) {
	o.DockerAsset = dockerAsset
}

// WithHelmAsset adds the helmAsset to the upload component params
func (o *UploadComponentParams) WithHelmAsset(helmAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetHelmAsset(helmAsset)
	return o
}

// SetHelmAsset adds the helmAsset to the upload component params
func (o *UploadComponentParams) SetHelmAsset(helmAsset runtime.NamedReadCloser) {
	o.HelmAsset = helmAsset
}

// WithMaven2ArtifactID adds the maven2ArtifactID to the upload component params
func (o *UploadComponentParams) WithMaven2ArtifactID(maven2ArtifactID *string) *UploadComponentParams {
	o.SetMaven2ArtifactID(maven2ArtifactID)
	return o
}

// SetMaven2ArtifactID adds the maven2ArtifactId to the upload component params
func (o *UploadComponentParams) SetMaven2ArtifactID(maven2ArtifactID *string) {
	o.Maven2ArtifactID = maven2ArtifactID
}

// WithMaven2Asset1 adds the maven2Asset1 to the upload component params
func (o *UploadComponentParams) WithMaven2Asset1(maven2Asset1 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetMaven2Asset1(maven2Asset1)
	return o
}

// SetMaven2Asset1 adds the maven2Asset1 to the upload component params
func (o *UploadComponentParams) SetMaven2Asset1(maven2Asset1 runtime.NamedReadCloser) {
	o.Maven2Asset1 = maven2Asset1
}

// WithMaven2Asset1Classifier adds the maven2Asset1Classifier to the upload component params
func (o *UploadComponentParams) WithMaven2Asset1Classifier(maven2Asset1Classifier *string) *UploadComponentParams {
	o.SetMaven2Asset1Classifier(maven2Asset1Classifier)
	return o
}

// SetMaven2Asset1Classifier adds the maven2Asset1Classifier to the upload component params
func (o *UploadComponentParams) SetMaven2Asset1Classifier(maven2Asset1Classifier *string) {
	o.Maven2Asset1Classifier = maven2Asset1Classifier
}

// WithMaven2Asset1Extension adds the maven2Asset1Extension to the upload component params
func (o *UploadComponentParams) WithMaven2Asset1Extension(maven2Asset1Extension *string) *UploadComponentParams {
	o.SetMaven2Asset1Extension(maven2Asset1Extension)
	return o
}

// SetMaven2Asset1Extension adds the maven2Asset1Extension to the upload component params
func (o *UploadComponentParams) SetMaven2Asset1Extension(maven2Asset1Extension *string) {
	o.Maven2Asset1Extension = maven2Asset1Extension
}

// WithMaven2Asset2 adds the maven2Asset2 to the upload component params
func (o *UploadComponentParams) WithMaven2Asset2(maven2Asset2 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetMaven2Asset2(maven2Asset2)
	return o
}

// SetMaven2Asset2 adds the maven2Asset2 to the upload component params
func (o *UploadComponentParams) SetMaven2Asset2(maven2Asset2 runtime.NamedReadCloser) {
	o.Maven2Asset2 = maven2Asset2
}

// WithMaven2Asset2Classifier adds the maven2Asset2Classifier to the upload component params
func (o *UploadComponentParams) WithMaven2Asset2Classifier(maven2Asset2Classifier *string) *UploadComponentParams {
	o.SetMaven2Asset2Classifier(maven2Asset2Classifier)
	return o
}

// SetMaven2Asset2Classifier adds the maven2Asset2Classifier to the upload component params
func (o *UploadComponentParams) SetMaven2Asset2Classifier(maven2Asset2Classifier *string) {
	o.Maven2Asset2Classifier = maven2Asset2Classifier
}

// WithMaven2Asset2Extension adds the maven2Asset2Extension to the upload component params
func (o *UploadComponentParams) WithMaven2Asset2Extension(maven2Asset2Extension *string) *UploadComponentParams {
	o.SetMaven2Asset2Extension(maven2Asset2Extension)
	return o
}

// SetMaven2Asset2Extension adds the maven2Asset2Extension to the upload component params
func (o *UploadComponentParams) SetMaven2Asset2Extension(maven2Asset2Extension *string) {
	o.Maven2Asset2Extension = maven2Asset2Extension
}

// WithMaven2Asset3 adds the maven2Asset3 to the upload component params
func (o *UploadComponentParams) WithMaven2Asset3(maven2Asset3 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetMaven2Asset3(maven2Asset3)
	return o
}

// SetMaven2Asset3 adds the maven2Asset3 to the upload component params
func (o *UploadComponentParams) SetMaven2Asset3(maven2Asset3 runtime.NamedReadCloser) {
	o.Maven2Asset3 = maven2Asset3
}

// WithMaven2Asset3Classifier adds the maven2Asset3Classifier to the upload component params
func (o *UploadComponentParams) WithMaven2Asset3Classifier(maven2Asset3Classifier *string) *UploadComponentParams {
	o.SetMaven2Asset3Classifier(maven2Asset3Classifier)
	return o
}

// SetMaven2Asset3Classifier adds the maven2Asset3Classifier to the upload component params
func (o *UploadComponentParams) SetMaven2Asset3Classifier(maven2Asset3Classifier *string) {
	o.Maven2Asset3Classifier = maven2Asset3Classifier
}

// WithMaven2Asset3Extension adds the maven2Asset3Extension to the upload component params
func (o *UploadComponentParams) WithMaven2Asset3Extension(maven2Asset3Extension *string) *UploadComponentParams {
	o.SetMaven2Asset3Extension(maven2Asset3Extension)
	return o
}

// SetMaven2Asset3Extension adds the maven2Asset3Extension to the upload component params
func (o *UploadComponentParams) SetMaven2Asset3Extension(maven2Asset3Extension *string) {
	o.Maven2Asset3Extension = maven2Asset3Extension
}

// WithMaven2GeneratePom adds the maven2GeneratePom to the upload component params
func (o *UploadComponentParams) WithMaven2GeneratePom(maven2GeneratePom *bool) *UploadComponentParams {
	o.SetMaven2GeneratePom(maven2GeneratePom)
	return o
}

// SetMaven2GeneratePom adds the maven2GeneratePom to the upload component params
func (o *UploadComponentParams) SetMaven2GeneratePom(maven2GeneratePom *bool) {
	o.Maven2GeneratePom = maven2GeneratePom
}

// WithMaven2GroupID adds the maven2GroupID to the upload component params
func (o *UploadComponentParams) WithMaven2GroupID(maven2GroupID *string) *UploadComponentParams {
	o.SetMaven2GroupID(maven2GroupID)
	return o
}

// SetMaven2GroupID adds the maven2GroupId to the upload component params
func (o *UploadComponentParams) SetMaven2GroupID(maven2GroupID *string) {
	o.Maven2GroupID = maven2GroupID
}

// WithMaven2Packaging adds the maven2Packaging to the upload component params
func (o *UploadComponentParams) WithMaven2Packaging(maven2Packaging *string) *UploadComponentParams {
	o.SetMaven2Packaging(maven2Packaging)
	return o
}

// SetMaven2Packaging adds the maven2Packaging to the upload component params
func (o *UploadComponentParams) SetMaven2Packaging(maven2Packaging *string) {
	o.Maven2Packaging = maven2Packaging
}

// WithMaven2Version adds the maven2Version to the upload component params
func (o *UploadComponentParams) WithMaven2Version(maven2Version *string) *UploadComponentParams {
	o.SetMaven2Version(maven2Version)
	return o
}

// SetMaven2Version adds the maven2Version to the upload component params
func (o *UploadComponentParams) SetMaven2Version(maven2Version *string) {
	o.Maven2Version = maven2Version
}

// WithNpmAsset adds the npmAsset to the upload component params
func (o *UploadComponentParams) WithNpmAsset(npmAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetNpmAsset(npmAsset)
	return o
}

// SetNpmAsset adds the npmAsset to the upload component params
func (o *UploadComponentParams) SetNpmAsset(npmAsset runtime.NamedReadCloser) {
	o.NpmAsset = npmAsset
}

// WithNugetAsset adds the nugetAsset to the upload component params
func (o *UploadComponentParams) WithNugetAsset(nugetAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetNugetAsset(nugetAsset)
	return o
}

// SetNugetAsset adds the nugetAsset to the upload component params
func (o *UploadComponentParams) SetNugetAsset(nugetAsset runtime.NamedReadCloser) {
	o.NugetAsset = nugetAsset
}

// WithPypiAsset adds the pypiAsset to the upload component params
func (o *UploadComponentParams) WithPypiAsset(pypiAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetPypiAsset(pypiAsset)
	return o
}

// SetPypiAsset adds the pypiAsset to the upload component params
func (o *UploadComponentParams) SetPypiAsset(pypiAsset runtime.NamedReadCloser) {
	o.PypiAsset = pypiAsset
}

// WithRAsset adds the rAsset to the upload component params
func (o *UploadComponentParams) WithRAsset(rAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetRAsset(rAsset)
	return o
}

// SetRAsset adds the rAsset to the upload component params
func (o *UploadComponentParams) SetRAsset(rAsset runtime.NamedReadCloser) {
	o.RAsset = rAsset
}

// WithRAssetPathID adds the rAssetPathID to the upload component params
func (o *UploadComponentParams) WithRAssetPathID(rAssetPathID *string) *UploadComponentParams {
	o.SetRAssetPathID(rAssetPathID)
	return o
}

// SetRAssetPathID adds the rAssetPathId to the upload component params
func (o *UploadComponentParams) SetRAssetPathID(rAssetPathID *string) {
	o.RAssetPathID = rAssetPathID
}

// WithRawAsset1 adds the rawAsset1 to the upload component params
func (o *UploadComponentParams) WithRawAsset1(rawAsset1 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetRawAsset1(rawAsset1)
	return o
}

// SetRawAsset1 adds the rawAsset1 to the upload component params
func (o *UploadComponentParams) SetRawAsset1(rawAsset1 runtime.NamedReadCloser) {
	o.RawAsset1 = rawAsset1
}

// WithRawAsset1Filename adds the rawAsset1Filename to the upload component params
func (o *UploadComponentParams) WithRawAsset1Filename(rawAsset1Filename *string) *UploadComponentParams {
	o.SetRawAsset1Filename(rawAsset1Filename)
	return o
}

// SetRawAsset1Filename adds the rawAsset1Filename to the upload component params
func (o *UploadComponentParams) SetRawAsset1Filename(rawAsset1Filename *string) {
	o.RawAsset1Filename = rawAsset1Filename
}

// WithRawAsset2 adds the rawAsset2 to the upload component params
func (o *UploadComponentParams) WithRawAsset2(rawAsset2 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetRawAsset2(rawAsset2)
	return o
}

// SetRawAsset2 adds the rawAsset2 to the upload component params
func (o *UploadComponentParams) SetRawAsset2(rawAsset2 runtime.NamedReadCloser) {
	o.RawAsset2 = rawAsset2
}

// WithRawAsset2Filename adds the rawAsset2Filename to the upload component params
func (o *UploadComponentParams) WithRawAsset2Filename(rawAsset2Filename *string) *UploadComponentParams {
	o.SetRawAsset2Filename(rawAsset2Filename)
	return o
}

// SetRawAsset2Filename adds the rawAsset2Filename to the upload component params
func (o *UploadComponentParams) SetRawAsset2Filename(rawAsset2Filename *string) {
	o.RawAsset2Filename = rawAsset2Filename
}

// WithRawAsset3 adds the rawAsset3 to the upload component params
func (o *UploadComponentParams) WithRawAsset3(rawAsset3 runtime.NamedReadCloser) *UploadComponentParams {
	o.SetRawAsset3(rawAsset3)
	return o
}

// SetRawAsset3 adds the rawAsset3 to the upload component params
func (o *UploadComponentParams) SetRawAsset3(rawAsset3 runtime.NamedReadCloser) {
	o.RawAsset3 = rawAsset3
}

// WithRawAsset3Filename adds the rawAsset3Filename to the upload component params
func (o *UploadComponentParams) WithRawAsset3Filename(rawAsset3Filename *string) *UploadComponentParams {
	o.SetRawAsset3Filename(rawAsset3Filename)
	return o
}

// SetRawAsset3Filename adds the rawAsset3Filename to the upload component params
func (o *UploadComponentParams) SetRawAsset3Filename(rawAsset3Filename *string) {
	o.RawAsset3Filename = rawAsset3Filename
}

// WithRawDirectory adds the rawDirectory to the upload component params
func (o *UploadComponentParams) WithRawDirectory(rawDirectory *string) *UploadComponentParams {
	o.SetRawDirectory(rawDirectory)
	return o
}

// SetRawDirectory adds the rawDirectory to the upload component params
func (o *UploadComponentParams) SetRawDirectory(rawDirectory *string) {
	o.RawDirectory = rawDirectory
}

// WithRepository adds the repository to the upload component params
func (o *UploadComponentParams) WithRepository(repository string) *UploadComponentParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the upload component params
func (o *UploadComponentParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithRubygemsAsset adds the rubygemsAsset to the upload component params
func (o *UploadComponentParams) WithRubygemsAsset(rubygemsAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetRubygemsAsset(rubygemsAsset)
	return o
}

// SetRubygemsAsset adds the rubygemsAsset to the upload component params
func (o *UploadComponentParams) SetRubygemsAsset(rubygemsAsset runtime.NamedReadCloser) {
	o.RubygemsAsset = rubygemsAsset
}

// WithYumAsset adds the yumAsset to the upload component params
func (o *UploadComponentParams) WithYumAsset(yumAsset runtime.NamedReadCloser) *UploadComponentParams {
	o.SetYumAsset(yumAsset)
	return o
}

// SetYumAsset adds the yumAsset to the upload component params
func (o *UploadComponentParams) SetYumAsset(yumAsset runtime.NamedReadCloser) {
	o.YumAsset = yumAsset
}

// WithYumAssetFilename adds the yumAssetFilename to the upload component params
func (o *UploadComponentParams) WithYumAssetFilename(yumAssetFilename *string) *UploadComponentParams {
	o.SetYumAssetFilename(yumAssetFilename)
	return o
}

// SetYumAssetFilename adds the yumAssetFilename to the upload component params
func (o *UploadComponentParams) SetYumAssetFilename(yumAssetFilename *string) {
	o.YumAssetFilename = yumAssetFilename
}

// WithYumDirectory adds the yumDirectory to the upload component params
func (o *UploadComponentParams) WithYumDirectory(yumDirectory *string) *UploadComponentParams {
	o.SetYumDirectory(yumDirectory)
	return o
}

// SetYumDirectory adds the yumDirectory to the upload component params
func (o *UploadComponentParams) SetYumDirectory(yumDirectory *string) {
	o.YumDirectory = yumDirectory
}

// WriteToRequest writes these params to a swagger request
func (o *UploadComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AptAsset != nil {

		if o.AptAsset != nil {
			// form file param apt.asset
			if err := r.SetFileParam("apt.asset", o.AptAsset); err != nil {
				return err
			}
		}
	}

	if o.DockerAsset != nil {

		if o.DockerAsset != nil {
			// form file param docker.asset
			if err := r.SetFileParam("docker.asset", o.DockerAsset); err != nil {
				return err
			}
		}
	}

	if o.HelmAsset != nil {

		if o.HelmAsset != nil {
			// form file param helm.asset
			if err := r.SetFileParam("helm.asset", o.HelmAsset); err != nil {
				return err
			}
		}
	}

	if o.Maven2ArtifactID != nil {

		// form param maven2.artifactId
		var frMaven2ArtifactID string
		if o.Maven2ArtifactID != nil {
			frMaven2ArtifactID = *o.Maven2ArtifactID
		}
		fMaven2ArtifactID := frMaven2ArtifactID
		if fMaven2ArtifactID != "" {
			if err := r.SetFormParam("maven2.artifactId", fMaven2ArtifactID); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset1 != nil {

		if o.Maven2Asset1 != nil {
			// form file param maven2.asset1
			if err := r.SetFileParam("maven2.asset1", o.Maven2Asset1); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset1Classifier != nil {

		// form param maven2.asset1.classifier
		var frMaven2Asset1Classifier string
		if o.Maven2Asset1Classifier != nil {
			frMaven2Asset1Classifier = *o.Maven2Asset1Classifier
		}
		fMaven2Asset1Classifier := frMaven2Asset1Classifier
		if fMaven2Asset1Classifier != "" {
			if err := r.SetFormParam("maven2.asset1.classifier", fMaven2Asset1Classifier); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset1Extension != nil {

		// form param maven2.asset1.extension
		var frMaven2Asset1Extension string
		if o.Maven2Asset1Extension != nil {
			frMaven2Asset1Extension = *o.Maven2Asset1Extension
		}
		fMaven2Asset1Extension := frMaven2Asset1Extension
		if fMaven2Asset1Extension != "" {
			if err := r.SetFormParam("maven2.asset1.extension", fMaven2Asset1Extension); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset2 != nil {

		if o.Maven2Asset2 != nil {
			// form file param maven2.asset2
			if err := r.SetFileParam("maven2.asset2", o.Maven2Asset2); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset2Classifier != nil {

		// form param maven2.asset2.classifier
		var frMaven2Asset2Classifier string
		if o.Maven2Asset2Classifier != nil {
			frMaven2Asset2Classifier = *o.Maven2Asset2Classifier
		}
		fMaven2Asset2Classifier := frMaven2Asset2Classifier
		if fMaven2Asset2Classifier != "" {
			if err := r.SetFormParam("maven2.asset2.classifier", fMaven2Asset2Classifier); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset2Extension != nil {

		// form param maven2.asset2.extension
		var frMaven2Asset2Extension string
		if o.Maven2Asset2Extension != nil {
			frMaven2Asset2Extension = *o.Maven2Asset2Extension
		}
		fMaven2Asset2Extension := frMaven2Asset2Extension
		if fMaven2Asset2Extension != "" {
			if err := r.SetFormParam("maven2.asset2.extension", fMaven2Asset2Extension); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset3 != nil {

		if o.Maven2Asset3 != nil {
			// form file param maven2.asset3
			if err := r.SetFileParam("maven2.asset3", o.Maven2Asset3); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset3Classifier != nil {

		// form param maven2.asset3.classifier
		var frMaven2Asset3Classifier string
		if o.Maven2Asset3Classifier != nil {
			frMaven2Asset3Classifier = *o.Maven2Asset3Classifier
		}
		fMaven2Asset3Classifier := frMaven2Asset3Classifier
		if fMaven2Asset3Classifier != "" {
			if err := r.SetFormParam("maven2.asset3.classifier", fMaven2Asset3Classifier); err != nil {
				return err
			}
		}
	}

	if o.Maven2Asset3Extension != nil {

		// form param maven2.asset3.extension
		var frMaven2Asset3Extension string
		if o.Maven2Asset3Extension != nil {
			frMaven2Asset3Extension = *o.Maven2Asset3Extension
		}
		fMaven2Asset3Extension := frMaven2Asset3Extension
		if fMaven2Asset3Extension != "" {
			if err := r.SetFormParam("maven2.asset3.extension", fMaven2Asset3Extension); err != nil {
				return err
			}
		}
	}

	if o.Maven2GeneratePom != nil {

		// form param maven2.generate-pom
		var frMaven2GeneratePom bool
		if o.Maven2GeneratePom != nil {
			frMaven2GeneratePom = *o.Maven2GeneratePom
		}
		fMaven2GeneratePom := swag.FormatBool(frMaven2GeneratePom)
		if fMaven2GeneratePom != "" {
			if err := r.SetFormParam("maven2.generate-pom", fMaven2GeneratePom); err != nil {
				return err
			}
		}
	}

	if o.Maven2GroupID != nil {

		// form param maven2.groupId
		var frMaven2GroupID string
		if o.Maven2GroupID != nil {
			frMaven2GroupID = *o.Maven2GroupID
		}
		fMaven2GroupID := frMaven2GroupID
		if fMaven2GroupID != "" {
			if err := r.SetFormParam("maven2.groupId", fMaven2GroupID); err != nil {
				return err
			}
		}
	}

	if o.Maven2Packaging != nil {

		// form param maven2.packaging
		var frMaven2Packaging string
		if o.Maven2Packaging != nil {
			frMaven2Packaging = *o.Maven2Packaging
		}
		fMaven2Packaging := frMaven2Packaging
		if fMaven2Packaging != "" {
			if err := r.SetFormParam("maven2.packaging", fMaven2Packaging); err != nil {
				return err
			}
		}
	}

	if o.Maven2Version != nil {

		// form param maven2.version
		var frMaven2Version string
		if o.Maven2Version != nil {
			frMaven2Version = *o.Maven2Version
		}
		fMaven2Version := frMaven2Version
		if fMaven2Version != "" {
			if err := r.SetFormParam("maven2.version", fMaven2Version); err != nil {
				return err
			}
		}
	}

	if o.NpmAsset != nil {

		if o.NpmAsset != nil {
			// form file param npm.asset
			if err := r.SetFileParam("npm.asset", o.NpmAsset); err != nil {
				return err
			}
		}
	}

	if o.NugetAsset != nil {

		if o.NugetAsset != nil {
			// form file param nuget.asset
			if err := r.SetFileParam("nuget.asset", o.NugetAsset); err != nil {
				return err
			}
		}
	}

	if o.PypiAsset != nil {

		if o.PypiAsset != nil {
			// form file param pypi.asset
			if err := r.SetFileParam("pypi.asset", o.PypiAsset); err != nil {
				return err
			}
		}
	}

	if o.RAsset != nil {

		if o.RAsset != nil {
			// form file param r.asset
			if err := r.SetFileParam("r.asset", o.RAsset); err != nil {
				return err
			}
		}
	}

	if o.RAssetPathID != nil {

		// form param r.asset.pathId
		var frRAssetPathID string
		if o.RAssetPathID != nil {
			frRAssetPathID = *o.RAssetPathID
		}
		fRAssetPathID := frRAssetPathID
		if fRAssetPathID != "" {
			if err := r.SetFormParam("r.asset.pathId", fRAssetPathID); err != nil {
				return err
			}
		}
	}

	if o.RawAsset1 != nil {

		if o.RawAsset1 != nil {
			// form file param raw.asset1
			if err := r.SetFileParam("raw.asset1", o.RawAsset1); err != nil {
				return err
			}
		}
	}

	if o.RawAsset1Filename != nil {

		// form param raw.asset1.filename
		var frRawAsset1Filename string
		if o.RawAsset1Filename != nil {
			frRawAsset1Filename = *o.RawAsset1Filename
		}
		fRawAsset1Filename := frRawAsset1Filename
		if fRawAsset1Filename != "" {
			if err := r.SetFormParam("raw.asset1.filename", fRawAsset1Filename); err != nil {
				return err
			}
		}
	}

	if o.RawAsset2 != nil {

		if o.RawAsset2 != nil {
			// form file param raw.asset2
			if err := r.SetFileParam("raw.asset2", o.RawAsset2); err != nil {
				return err
			}
		}
	}

	if o.RawAsset2Filename != nil {

		// form param raw.asset2.filename
		var frRawAsset2Filename string
		if o.RawAsset2Filename != nil {
			frRawAsset2Filename = *o.RawAsset2Filename
		}
		fRawAsset2Filename := frRawAsset2Filename
		if fRawAsset2Filename != "" {
			if err := r.SetFormParam("raw.asset2.filename", fRawAsset2Filename); err != nil {
				return err
			}
		}
	}

	if o.RawAsset3 != nil {

		if o.RawAsset3 != nil {
			// form file param raw.asset3
			if err := r.SetFileParam("raw.asset3", o.RawAsset3); err != nil {
				return err
			}
		}
	}

	if o.RawAsset3Filename != nil {

		// form param raw.asset3.filename
		var frRawAsset3Filename string
		if o.RawAsset3Filename != nil {
			frRawAsset3Filename = *o.RawAsset3Filename
		}
		fRawAsset3Filename := frRawAsset3Filename
		if fRawAsset3Filename != "" {
			if err := r.SetFormParam("raw.asset3.filename", fRawAsset3Filename); err != nil {
				return err
			}
		}
	}

	if o.RawDirectory != nil {

		// form param raw.directory
		var frRawDirectory string
		if o.RawDirectory != nil {
			frRawDirectory = *o.RawDirectory
		}
		fRawDirectory := frRawDirectory
		if fRawDirectory != "" {
			if err := r.SetFormParam("raw.directory", fRawDirectory); err != nil {
				return err
			}
		}
	}

	// query param repository
	qrRepository := o.Repository
	qRepository := qrRepository
	if qRepository != "" {

		if err := r.SetQueryParam("repository", qRepository); err != nil {
			return err
		}
	}

	if o.RubygemsAsset != nil {

		if o.RubygemsAsset != nil {
			// form file param rubygems.asset
			if err := r.SetFileParam("rubygems.asset", o.RubygemsAsset); err != nil {
				return err
			}
		}
	}

	if o.YumAsset != nil {

		if o.YumAsset != nil {
			// form file param yum.asset
			if err := r.SetFileParam("yum.asset", o.YumAsset); err != nil {
				return err
			}
		}
	}

	if o.YumAssetFilename != nil {

		// form param yum.asset.filename
		var frYumAssetFilename string
		if o.YumAssetFilename != nil {
			frYumAssetFilename = *o.YumAssetFilename
		}
		fYumAssetFilename := frYumAssetFilename
		if fYumAssetFilename != "" {
			if err := r.SetFormParam("yum.asset.filename", fYumAssetFilename); err != nil {
				return err
			}
		}
	}

	if o.YumDirectory != nil {

		// form param yum.directory
		var frYumDirectory string
		if o.YumDirectory != nil {
			frYumDirectory = *o.YumDirectory
		}
		fYumDirectory := frYumDirectory
		if fYumDirectory != "" {
			if err := r.SetFormParam("yum.directory", fYumDirectory); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
