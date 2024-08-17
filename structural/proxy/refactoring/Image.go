package proxy

type Image struct {
	fileName    string
	description string
	picture     Picture
}

func (real *Image) displayShort() string {
	return "Real image: " + real.fileName + ": " + real.description
}

func (real *Image) displayFull() string {
	return "Real image: " + real.fileName + ": " + real.description + ": " + real.picture.display()
}
