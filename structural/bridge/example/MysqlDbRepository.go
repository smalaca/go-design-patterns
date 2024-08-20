package bridge

import "fmt"

type MysqlDbRepository struct {
	Repository
}

func createMysqlDbRepository(cache Cache) MysqlDbRepository {
	repository := MysqlDbRepository{};
	repository.cache = cache
	repository.refinement = &repository
	return repository
}

func (repository *MysqlDbRepository) saveOne(entity Entity) {
	fmt.Println("Saving in MySQL")
}