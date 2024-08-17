package proxy

type ImageRepository struct {
	imageStorage ImageStorage
}

func (repository *ImageRepository) findByName(name string) Image {
	return Image{
		fileName: name,
		description: "Description of " + name,
		picture: repository.imageStorage.load(name),
	}
}