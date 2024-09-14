package observer

import "strconv"

type ProductController struct {
}

func createProductController() ProductController {
	return ProductController{}
}

func (controller *ProductController) buy(productId int) {
	if controller.exists(productId) {
		product := controller.findById(productId)
		product.bought()
		controller.notifyAll(product)
	}
}

func (controller *ProductController) exists(productId int) bool {
	return true
}

func (controller *ProductController) findById(productId int) Product {
	return Product{
		productId: productId,
		name:      "Product No: " + strconv.Itoa(productId),
	}
}

func (controller *ProductController) notifyAll(product Product) {

}

// func (controller *ProductController) unregisterObserver(observers []Observer, toRemove Observer) []Observer {
// 	length := len(observers)

// 	for i, observer := range observers {
// 		if toRemove.name() == observer.name() {
// 			observers[length-1], observers[i] = observers[i], observers[length-1]
// 			return observers[:length-1]
// 		}
// 	}

// 	return observers
// }
