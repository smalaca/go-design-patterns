package proxy

type Image interface {
	displayShort() string
	displayFull() string
}