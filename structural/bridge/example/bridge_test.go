package bridge

import (
	"testing"
)

func Test_Bridge(t *testing.T) {
	repository := createMySqlRepository(&FileCache{})
	entity := Entity{id: 13, name: "Lucky"}

	repository.save(entity)

	repository.findById(13)
	repository.findById(42)
	
	repository.delete(13)
	repository.delete(42)
}
