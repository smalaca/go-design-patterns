package bridge

import (
	"testing"
)

func Test_Bridge(t *testing.T) {
	repositoryMongo := createMongoDbRepository(&InMemoryCache{})
	entity := Entity{id: 13, name: "Lucky"}

	repositoryMongo.save(entity)
	
	repositoryMysql := createMysqlDbRepository(&InMemoryCache{})
	repositoryMysql.save(entity)

	// repository.findById(13)
	// repository.findById(42)
	
	// repository.delete(13)
	// repository.delete(42)
}
