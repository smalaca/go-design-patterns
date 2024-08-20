package bridge

import "fmt"

type MySqlRepository struct {
	Repository
}

func createMySqlRepository(cache iCache) MySqlRepository {
	repository := MySqlRepository{}
	repository.cache = cache
	repository.refinement = &repository
	return repository
}

func (repository *MySqlRepository) persist(entity Entity) {
	fmt.Println("Persistiing entity in MySQL: " + entity.name)
}

func (repository *MySqlRepository) findOneById(id int) Entity {
	fmt.Println("Finding entity in MySQL");
	return Entity{id: id, name: "Found in MySQL"}
}

func (repository *MySqlRepository) deleteOneById(id int) {
	fmt.Println("Removing entity from MySQL");
}

