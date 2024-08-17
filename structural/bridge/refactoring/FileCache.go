package bridge

import "fmt"

type FileCache struct {
	entities map[int]Entity
}

func (cache *FileCache) contains(id int) bool {
	fmt.Println("Checks if exists in file cache")
	_, ok := cache.getEntities()[id]
	return ok
}

func (cache *FileCache) get(id int) Entity {
	fmt.Println("Returns from file cache")
	return cache.getEntities()[id]
}

func (cache *FileCache) add(entity Entity) {
	fmt.Println("Adds to file cache")
	cache.getEntities()[entity.id] = entity
}

func (cache *FileCache) delete(id int) {
	fmt.Println("Deletes from file cache")
	delete(cache.getEntities(), id)
}

func (cache *FileCache) getEntities() map[int]Entity {
	if (cache.entities == nil) {
		cache.entities = make(map[int]Entity)
	}
	
	return cache.entities
}