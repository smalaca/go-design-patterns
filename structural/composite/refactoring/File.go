package composite

import "errors"

type File struct {
	path string
	size FileSize
}

func (file *File) existsWith(path string) bool {
	return file.path == path
}

func (file *File) get(path string) (*File, error) {
	if file.existsWith(path) {
		return file, nil
	} else {
		return nil, errors.New("No file found in this path")
	}
}

func (file *File) getSize() FileSize {
	return file.size
}
