package composite

import "errors"

type Directory struct {
	path  string
	files []FileNode
}

func createDirectory(path string) *Directory {
	return &Directory{
		path: path,
		files: []FileNode{},
	}
}

func (directory *Directory) add(node FileNode) {
	directory.files = append(directory.files, node)
}

func (directory *Directory) existsWith(path string) bool {
	if (directory.path == path) {
		return true;
	}

	for _, file := range directory.files {
		if (file.existsWith(path)) {
			return true;
		}
	}

	return false
}

func (directory *Directory) get(path string) (FileNode, error) {
	if directory.path == path {
		return directory, nil
	} 
	
	for _, file := range directory.files {
		if (file.existsWith(path)) {
			return file, nil;
		}
	}

	return nil, errors.New("No file found in this path")
}

func (directory *Directory) getSize() FileSize {
	size := FileSize{value: 0}

	for _, file := range directory.files {
		size = size.add(file.getSize())
	}

	return size
}
