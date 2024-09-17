package visitor

type Issue interface {
	accept(visitor Visitor)
}
