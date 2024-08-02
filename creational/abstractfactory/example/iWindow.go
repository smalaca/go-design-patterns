package abstractfactory

type iWindow interface {
	display()
	load(content string)
}