package bridge

import "fmt"

type MongoDbRepository struct {
	cache iCache
}

func createMongoDbRepository(cache iCache) MongoDbRepository {
	repository := MongoDbRepository{}
	repository.cache = cache
	return repository
}

func (repository *MongoDbRepository) save(entity Entity) {
	repository.persist(entity)
	fmt.Println("Saving entity: " + entity.name)
	repository.cache.add(entity)
}

func (repository *MongoDbRepository) persist(entity Entity) {
	fmt.Println("Persistiing entity in Mongo DB: " + entity.name)
}

func (repository *MongoDbRepository) findById(id int) Entity {
	fmt.Println("Find entity by id")
	if (repository.cache.contains(id)) {
		return repository.cache.get(id)
	} else {
		return repository.findOneById(id)
	}
}

func (repository *MongoDbRepository) findOneById(id int) Entity {
	fmt.Println("Finding entity in Mongo DB");
	return Entity{id: id, name: "Found in MongoDB"}
}

func (repository *MongoDbRepository) delete(id int) {
	fmt.Println("Deleting entity")

	if (repository.cache.contains(id)) {
		repository.cache.delete(id);
	}

	repository.deleteOneById(id)
}

func (repository *MongoDbRepository) deleteOneById(id int) {
	fmt.Println("Removing entity from Mongo DB");
}

