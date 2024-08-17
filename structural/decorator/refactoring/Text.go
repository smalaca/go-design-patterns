package decorator

type Text struct {
	value string
	isBold bool
	isItalic bool
	isUnderline bool
}

func (node *Text) getValue() string {
	value := node.value

	if node.isBold {
		value = "<B>" + value + "</B>"
	}

	if node.isItalic {
		value = "<I>" + value + "</I>"
	}

	if node.isUnderline {
		value = "<U>" + value + "</U>"
	}

	return value
}

func (node *Text) markAsBold() {
	node.isBold = true
}

func (node *Text) markAsItalic() {
	node.isItalic = true
}

func (node *Text) markAsUnderline() {
	node.isUnderline = true
}