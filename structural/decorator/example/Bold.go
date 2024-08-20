package decorator

type Bold struct {
	node TextNode
}

func (node *Bold) getValue() string {
	return "<B>" + node.node.getValue() + "</B>"
}