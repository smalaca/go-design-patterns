package proxy

type RealImage struct {
	fileName    string
	description string
	picture     Picture
}

func (real *RealImage) displayShort() string {
	return "Real image: " + real.fileName + ": " + real.description
}

func (real *RealImage) displayFull() string {
	return "Real image: " + real.fileName + ": " + real.description + ": " + real.picture.display()
}
