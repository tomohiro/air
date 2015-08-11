package main

// Resource is online file
type resource struct {
	Path string
}

// NewResource creates a new resource
func NewResource(path string) *resource {
	return &resource{Path: path}
}

// URL returns resource's url
func (r *resource) URL() string {
	return r.Path
}
