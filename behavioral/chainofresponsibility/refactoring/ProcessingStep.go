package chainofresponsibility

type ProcessingStep interface {
	execute()
}