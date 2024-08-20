package bridge

import "fmt"

type InMemoryCache struct {
	entities map[int]Entity
}

func (cache *InMemoryCache) add(entity Entity) {
	fmt.Println("Adding entity to cache")
	cache.getEntities()[entity.id] = entity
}

func (cache *InMemoryCache) getEntities() map[int]Entity {
	if (cache.entities == nil) {
		cache.entities = make(map[int]Entity)
	}

	return cache.entities
}