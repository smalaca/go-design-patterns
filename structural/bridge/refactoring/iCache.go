package bridge

type iCache interface {
	contains(id int) bool
	
	get(id int) Entity

	add(entity Entity)

	delete(id int)
}