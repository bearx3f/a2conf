package a2conf

// Tree .
type Tree struct {
	Name   string // VirtualHost
	Value  string // *:80
	Blank  bool
	Parent *Tree
	Child  []*Tree
}

// Siblings .
func (t *Tree) Siblings() []*Tree {
	return t.Parent.Child
}
