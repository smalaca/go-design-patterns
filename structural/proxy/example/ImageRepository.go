package proxy

type ImageRepository struct {
	imageStorage ImageStorage
}

func (repository *ImageRepository) findByName(name string) string {
	return name
}