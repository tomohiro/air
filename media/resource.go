package media

// Resource is online file
type Resource struct {
	Path string
}

// NewResource creates a new resource
func NewResource(path string) *Resource {
	return &Resource{Path: path}
}

// URL returns a resource's url
func (r *Resource) URL() string {
	return r.Path
}
