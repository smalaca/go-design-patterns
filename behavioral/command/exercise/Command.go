package command

type Command interface {
	execute(product Product)
}