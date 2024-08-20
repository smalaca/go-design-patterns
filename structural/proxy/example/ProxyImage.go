package proxy

type ProxyImage struct {
	fileName string
	description string
	image Image
	imageStorage ImageStorage
}

func (proxy *ProxyImage) displayShort() string {
	return "Proxy image: " + proxy.fileName + ": " + proxy.description
}

func (proxy *ProxyImage) displayFull() string {
	if (proxy.image == nil) {
		proxy.image = proxy.imageStorage.load(proxy.fileName, proxy.description);
	}
	return proxy.image.displayFull()
}