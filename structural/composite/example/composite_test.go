package composite

import (
	"testing"
)

func Test_Composite(t *testing.T) {
	// adapter := createDirectory("/structural/adaper")
	// adapter.add(createDirectory("/structural/adaper/refactoring"))
	// adapter.add(&File{path: "/structural/adaper/refactoring/adapter_test.go", size: FileSize{5}})
	// adapter.add(createDirectory("/structural/adaper/example"))
	// adapter.add(&File{path: "/structural/adaper/example/adapter_test.go", size: FileSize{5}})

	// bridge := createDirectory("/structural/bridge")
	// bridge.add(&File{path: "/structural/bridge/bridge_test.go", size: FileSize{13}})
	// composite := createDirectory("/structural/composite")
	// composite.add(&File{path: "/structural/composite/composite_test.go", size: FileSize{42}})
	
	// root := createDirectory("/structural")
	// root.add(adapter)
	// root.add(bridge)
	// root.add(composite)

	// fmt.Println(root.existsWith("/structural/adaper/refactoring/adapter_test.go"))
	// fmt.Println(root.existsWith("/structural/composite"))
	// fmt.Println(root.existsWith("/structural"))
	// fmt.Println(root.existsWith("/creational"))
	// fmt.Println(root.getSize())
	
	// shouldExist, noError := root.get("/structural/composite")
	// fmt.Println(noError)
	// fmt.Println(shouldExist)
	
	// shouldNotExist, error := root.get("/creational")
	// fmt.Println(error)
	// fmt.Println(shouldNotExist)
}
