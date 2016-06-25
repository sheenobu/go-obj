package obj

// An Object is the toplevel loadable object
type Object struct {
	Name     string
	Vertices []Vertex
	Normals  []Normal
	Textures []TextureCoord
	Faces    []Face
}
