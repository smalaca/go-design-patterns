package iterator

type Iterable interface {
	iterator() Iterator
}
