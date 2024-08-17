package proxy

import "fmt"

type ImageStorage struct {

}

func (storage *ImageStorage) load(fileName string) Picture {
	fmt.Println("Loading image");
	return Picture{name: fileName}
}