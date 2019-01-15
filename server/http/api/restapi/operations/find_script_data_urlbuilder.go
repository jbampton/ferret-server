// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// FindScriptDataURL generates an URL for the find script data operation
type FindScriptDataURL struct {
	ProjectID string
	ScriptID  string

	Page *int32
	Size *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindScriptDataURL) WithBasePath(bp string) *FindScriptDataURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindScriptDataURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindScriptDataURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/projects/{projectID}/data/{scriptID}"

	projectID := o.ProjectID
	if projectID != "" {
		_path = strings.Replace(_path, "{projectID}", projectID, -1)
	} else {
		return nil, errors.New("ProjectID is required on FindScriptDataURL")
	}

	scriptID := o.ScriptID
	if scriptID != "" {
		_path = strings.Replace(_path, "{scriptID}", scriptID, -1)
	} else {
		return nil, errors.New("ScriptID is required on FindScriptDataURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var page string
	if o.Page != nil {
		page = swag.FormatInt32(*o.Page)
	}
	if page != "" {
		qs.Set("page", page)
	}

	var size string
	if o.Size != nil {
		size = swag.FormatInt32(*o.Size)
	}
	if size != "" {
		qs.Set("size", size)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindScriptDataURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindScriptDataURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindScriptDataURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindScriptDataURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindScriptDataURL")
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
func (o *FindScriptDataURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}