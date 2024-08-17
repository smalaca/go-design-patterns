package bridge

import "fmt"

type InMemoryCache struct {
	entities map[int]Entity
}

func (cache *InMemoryCache) contains(id int) bool {
	fmt.Println("Checks if exists in in memory cache")
	_, ok := cache.getEntities()[id]
	return ok
}

func (cache *InMemoryCache) get(id int) Entity {
	fmt.Println("Returns from in memory cache")
	return cache.getEntities()[id]
}

func (cache *InMemoryCache) add(entity Entity) {
	fmt.Println("Adds to in memory cache")
	cache.getEntities()[entity.id] = entity
}

func (cache *InMemoryCache) delete(id int) {
	fmt.Println("Deletes from in memory cache")
	delete(cache.getEntities(), id)
}

func (cache *InMemoryCache) getEntities() map[int]Entity {
	if (cache.entities == nil) {
		cache.entities = make(map[int]Entity)
	}
	
	return cache.entities
}