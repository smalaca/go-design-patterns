package composite

import "errors"

type Directory struct {
	path  string
	files []File
	directories []Directory
}

func createDirectory(path string) Directory {
	return Directory{
		path: path,
		files: []File{},
		directories: []Directory{},
	}
}

func (directory *Directory) addDirectory(node Directory) {
	directory.directories = append(directory.directories, node)
}

func (directory *Directory) addFile(node File) {
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

	for _, dir := range directory.directories {
		if (dir.existsWith(path)) {
			return true;
		}
	}

	return false
}

func (directory *Directory) getFile(path string) (*File, error) {
	for _, file := range directory.files {
		if (file.existsWith(path)) {
			return &file, nil;
		}
	}

	return nil, errors.New("No file found in this path")
}

func (directory *Directory) getDirectory(path string) (*Directory, error) {
	if directory.path == path {
		return directory, nil
	} 
	
	for _, dir := range directory.directories {
		if (dir.existsWith(path)) {
			return &dir, nil;
		}
	}

	return nil, errors.New("No directory found in this path")
}

func (directory *Directory) getSize() FileSize {
	size := FileSize{value: 0}

	for _, file := range directory.files {
		size = size.add(file.getSize())
	}

	for _, dir := range directory.directories {
		size = size.add(dir.getSize())
	}

	return size
}
