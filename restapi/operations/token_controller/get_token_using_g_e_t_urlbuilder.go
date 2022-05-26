// Code generated by go-swagger; DO NOT EDIT.

package token_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetTokenUsingGETURL generates an URL for the get token using g e t operation
type GetTokenUsingGETURL struct {
	AddressHash string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetTokenUsingGETURL) WithBasePath(bp string) *GetTokenUsingGETURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetTokenUsingGETURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetTokenUsingGETURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/tokens/{addressHash}"

	addressHash := o.AddressHash
	if addressHash != "" {
		_path = strings.Replace(_path, "{addressHash}", addressHash, -1)
	} else {
		return nil, errors.New("addressHash is required on GetTokenUsingGETURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetTokenUsingGETURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetTokenUsingGETURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetTokenUsingGETURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetTokenUsingGETURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetTokenUsingGETURL")
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
func (o *GetTokenUsingGETURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
