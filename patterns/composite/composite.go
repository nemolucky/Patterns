package composite

type Component interface {
	Size() int
}

type File struct {
	Title string
	Bytes int
}

func (f *File) Size() int {
	return f.Bytes
}

type Folder struct {
	Title    string
	children []Component
}

func (f *Folder) Add(component Component) {
	f.children = append(f.children, component)
}

func (f *Folder) Size() int {
	total := 0

	for _, child := range f.children {
		total += child.Size()
	}

	return total
}
