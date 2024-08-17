package decorator

type Text struct {
	value string
}

func (node *Text) getValue() string {
	return node.value
}