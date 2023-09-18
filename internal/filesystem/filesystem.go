package filesystem

type file struct {
	name *string
}

func NewFileSystem(name *string) *file {
	return &file{name}
}
