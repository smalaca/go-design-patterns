package bridge

import "fmt"

type MongoDbRepository struct {
	Repository
}

func createMongoDbRepository(cache Cache) MongoDbRepository {
	repository := MongoDbRepository{};
	repository.cache = cache
	repository.refinement = &repository
	return repository
}

func (repository *MongoDbRepository) saveOne(entity Entity) {
	fmt.Println("Saving in MongoDB")
}