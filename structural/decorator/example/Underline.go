package decorator

type Underline struct {
	node TextNode
}

func (node *Underline) getValue() string {
	return "<U>" + node.node.getValue() + "</U>"
}