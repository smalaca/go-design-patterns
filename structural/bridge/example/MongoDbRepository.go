package bridge

import "fmt"

type MongoDbRepository struct {
	Repository
}

func createMongoDbRepository(cache iCache) MongoDbRepository {
	repository := MongoDbRepository{}
	repository.cache = cache
	repository.refinement = &repository
	return repository
}

func (repository *MongoDbRepository) persist(entity Entity) {
	fmt.Println("Persistiing entity in Mongo DB: " + entity.name)
}

func (repository *MongoDbRepository) findOneById(id int) Entity {
	fmt.Println("Finding entity in Mongo DB");
	return Entity{id: id, name: "Found in MongoDB"}
}

func (repository *MongoDbRepository) deleteOneById(id int) {
	fmt.Println("Removing entity from Mongo DB");
}

