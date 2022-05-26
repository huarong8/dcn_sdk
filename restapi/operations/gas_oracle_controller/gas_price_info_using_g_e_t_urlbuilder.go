// Code generated by go-swagger; DO NOT EDIT.

package gas_oracle_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GasPriceInfoUsingGETURL generates an URL for the gas price info using g e t operation
type GasPriceInfoUsingGETURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GasPriceInfoUsingGETURL) WithBasePath(bp string) *GasPriceInfoUsingGETURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GasPriceInfoUsingGETURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GasPriceInfoUsingGETURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/gas/price"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GasPriceInfoUsingGETURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GasPriceInfoUsingGETURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GasPriceInfoUsingGETURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GasPriceInfoUsingGETURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GasPriceInfoUsingGETURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GasPriceInfoUsingGETURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
