// Code generated by go-swagger; DO NOT EDIT.

package contract_transactions_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// FindContractTransactionsUsingGETURL generates an URL for the find contract transactions using g e t operation
type FindContractTransactionsUsingGETURL struct {
	AddressHash string

	Direction *string
	Filter    *string
	Page      *int32
	Size      *int32
	Sort      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindContractTransactionsUsingGETURL) WithBasePath(bp string) *FindContractTransactionsUsingGETURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindContractTransactionsUsingGETURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindContractTransactionsUsingGETURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/contracts/{addressHash}/transactions"

	addressHash := o.AddressHash
	if addressHash != "" {
		_path = strings.Replace(_path, "{addressHash}", addressHash, -1)
	} else {
		return nil, errors.New("addressHash is required on FindContractTransactionsUsingGETURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var directionQ string
	if o.Direction != nil {
		directionQ = *o.Direction
	}
	if directionQ != "" {
		qs.Set("direction", directionQ)
	}

	var filterQ string
	if o.Filter != nil {
		filterQ = *o.Filter
	}
	if filterQ != "" {
		qs.Set("filter", filterQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = swag.FormatInt32(*o.Page)
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	var sizeQ string
	if o.Size != nil {
		sizeQ = swag.FormatInt32(*o.Size)
	}
	if sizeQ != "" {
		qs.Set("size", sizeQ)
	}

	var sortQ string
	if o.Sort != nil {
		sortQ = *o.Sort
	}
	if sortQ != "" {
		qs.Set("sort", sortQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindContractTransactionsUsingGETURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindContractTransactionsUsingGETURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindContractTransactionsUsingGETURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindContractTransactionsUsingGETURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindContractTransactionsUsingGETURL")
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
func (o *FindContractTransactionsUsingGETURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
