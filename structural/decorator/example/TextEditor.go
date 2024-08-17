package decorator

type TextEditor struct {
}

func (editor *TextEditor) write(sentence string) Text {
	return Text{sentence}
}

func (editor *TextEditor) decorate(node Text, decoration string) Text {
	return node
}
