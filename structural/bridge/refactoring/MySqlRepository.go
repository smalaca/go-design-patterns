package bridge

import "fmt"

type MySqlRepository struct {
	cache iCache
}

func createMySqlRepository(cache iCache) MySqlRepository {
	repository := MySqlRepository{}
	repository.cache = cache
	return repository
}

func (repository *MySqlRepository) save(entity Entity) {
	repository.persist(entity)
	fmt.Println("Saving entity: " + entity.name)
	repository.cache.add(entity)
}

func (repository *MySqlRepository) persist(entity Entity) {
	fmt.Println("Persistiing entity in MySQL: " + entity.name)
}

func (repository *MySqlRepository) findById(id int) Entity {
	fmt.Println("Find entity by id")
	if (repository.cache.contains(id)) {
		return repository.cache.get(id)
	} else {
		return repository.findOneById(id)
	}
}

func (repository *MySqlRepository) findOneById(id int) Entity {
	fmt.Println("Finding entity in MySQL");
	return Entity{id: id, name: "Found in MySQL"}
}

func (repository *MySqlRepository) delete(id int) {
	fmt.Println("Deleting entity")

	if (repository.cache.contains(id)) {
		repository.cache.delete(id);
	}

	repository.deleteOneById(id)
}

func (repository *MySqlRepository) deleteOneById(id int) {
	fmt.Println("Removing entity from MySQL");
}

