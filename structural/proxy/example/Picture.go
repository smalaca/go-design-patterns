package proxy

type Picture struct {
	name string
}

func (picture *Picture) display() string {
	return "Picture: <" + picture.name + ">"
}