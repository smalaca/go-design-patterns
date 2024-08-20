package decorator

type TextEditor struct {
}

func (editor *TextEditor) write(sentence string) TextNode {
	return &Text{sentence}
}

func (editor *TextEditor) decorate(node TextNode, decoration string) TextNode {
	switch decoration {
		case "B": return &Bold{node: node}
		case "U": return &Underline{node: node}
		case "I": return &Italic{node: node}
	}
	
	return node
}
