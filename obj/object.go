package obj

// An Object is the toplevel loadable object
type Object struct {
	Name     string
	Vertices []Vertex
	Normals  []Normal
	Textures []TextureCoord
	Faces    []Face

	// Custom types for custom
	Custom map[string][]interface{}
}

// AddCustom adds a custom object by key to the Custom map
func (o *Object) AddCustom(key string, ox interface{}) {
	if o.Custom == nil {
		o.Custom = make(map[string][]interface{})
	}
	l, ok := o.Custom[key]
	if !ok {
		l = make([]interface{}, 0)
	}

	l = append(l, ox)
	o.Custom[key] = l
}

// GetCustom gets the custom fields added to this object
func (o *Object) GetCustom(key string) (ix []interface{}, ok bool) {
	if o.Custom == nil {
		ok = false
		return
	}

	ix, ok = o.Custom[key]
	return
}
