package proxy

type ImageRepository struct {
	imageStorage ImageStorage
}

func (repository *ImageRepository) findByName(name string) Image {
	return &ProxyImage{
		fileName: name,
		description: "Description of " + name,
		imageStorage: repository.imageStorage,
	}
}