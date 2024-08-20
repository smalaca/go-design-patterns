package bridge

type Cache interface {
	// contains(id int) bool
	add(entity Entity)
	// remove(entity Entity)
	// get(id int) Entity
}