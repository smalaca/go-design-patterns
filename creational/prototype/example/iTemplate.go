package prototype

type iTemplate interface {
	print()
	clone() iTemplate
}
