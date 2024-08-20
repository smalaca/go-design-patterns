package composite

type FileNode interface {
	existsWith(path string) bool
    get(path string) (FileNode, error)
    getSize() FileSize
}