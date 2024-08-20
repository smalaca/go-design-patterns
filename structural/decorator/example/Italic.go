package decorator

type Italic struct {
	node TextNode
}

func (node *Italic) getValue() string {
	return "<I>" + node.node.getValue() + "</I>"
}