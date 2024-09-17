package chainofresponsibility

type ProcessStep interface {
	execute()
}
