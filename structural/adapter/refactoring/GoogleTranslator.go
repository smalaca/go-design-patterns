package adapter

type Request struct {
	from string
	to string
	toTranslate string
}

type Response struct {
	translation string
}

type GoogleTranslator struct {
}

func (translator *GoogleTranslator) translate(request Request) Response {
	if (request.from == "PL") {
		return Response{"Hello World"}
	} else {
		return Response{"Witaj Świecie"}
	}
}