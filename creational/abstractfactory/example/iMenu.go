package abstractfactory

type iMenu interface {
	display()
	load(items []string)
}
