package bridge

type RepositoryRefinement interface {
	saveOne(entity Entity)
}

type Repository struct {
	cache Cache
	refinement RepositoryRefinement
}

func (repository *Repository) save(entity Entity) {
	repository.refinement.saveOne(entity)
	repository.cache.add(entity)
}