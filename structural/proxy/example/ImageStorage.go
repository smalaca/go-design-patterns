package proxy

import "fmt"

type ImageStorage struct {

}

func (storage *ImageStorage) load(fileName string, description string) Image {
	fmt.Println("Loading image");
	
	return &RealImage{
		fileName: fileName,
		description: description,
		picture: Picture{name: fileName},
	}
}