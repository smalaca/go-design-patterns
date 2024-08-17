package decorator

type TextEditor struct {
}

func (editor *TextEditor) write(sentence string) Text {
	return Text{value: sentence}
}

func (editor *TextEditor) decorate(node Text, decoration string) Text {
	switch decoration {
		case "B": node.markAsBold()
		case "U": node.markAsUnderline()
		case "I": node.markAsItalic()
	}
	
	return node
}
