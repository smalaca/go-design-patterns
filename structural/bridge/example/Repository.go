package bridge

import "fmt"

type RepositoryRefinement interface {
	persist(entity Entity)
	findOneById(id int) Entity
	deleteOneById(id int)
}

type Repository struct {
	cache iCache
	refinement RepositoryRefinement
}

func (repository *Repository) save(entity Entity) {
	repository.refinement.persist(entity)
	fmt.Println("Saving entity: " + entity.name)
	repository.cache.add(entity)
}

func (repository *Repository) findById(id int) Entity {
	fmt.Println("Find entity by id")
	if (repository.cache.contains(id)) {
		return repository.cache.get(id)
	} else {
		return repository.refinement.findOneById(id)
	}
}

func (repository *Repository) delete(id int) {
	fmt.Println("Deleting entity")

	if (repository.cache.contains(id)) {
		repository.cache.delete(id);
	}

	repository.refinement.deleteOneById(id)
}