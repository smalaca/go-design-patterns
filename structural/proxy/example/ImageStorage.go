package proxy

import "fmt"

type ImageStorage struct {

}

func (storage *ImageStorage) load(fileName string, description string) string {
	fmt.Println("Loading image");
	return fileName
}