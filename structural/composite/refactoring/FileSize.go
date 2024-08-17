package composite

type FileSize struct {
	value int
}

func (size *FileSize) add(fileSize FileSize) FileSize{
	return FileSize{size.value + fileSize.value};
}