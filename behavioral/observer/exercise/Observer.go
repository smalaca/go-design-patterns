package observer

type Observer interface {
	notify(product Product)
	name() string
}
