package composite

import (
	"fmt"
	"testing"
)

func Test_Composite(t *testing.T) {
	adapter := createDirectory("/structural/adaper")
	adapter.addDirectory(createDirectory("/structural/adaper/refactoring"))
	adapter.addFile(File{path: "/structural/adaper/refactoring/adapter_test.go", size: FileSize{5}})
	adapter.addDirectory(createDirectory("/structural/adaper/example"))
	adapter.addFile(File{path: "/structural/adaper/example/adapter_test.go", size: FileSize{5}})

	bridge := createDirectory("/structural/bridge")
	bridge.addFile(File{path: "/structural/bridge/bridge_test.go", size: FileSize{13}})
	composite := createDirectory("/structural/composite")
	composite.addFile(File{path: "/structural/composite/composite_test.go", size: FileSize{42}})
	
	root := createDirectory("/structural")
	root.addDirectory(adapter)
	root.addDirectory(bridge)
	root.addDirectory(composite)

	fmt.Println(root.existsWith("/structural/adaper/refactoring/adapter_test.go"))
	fmt.Println(root.existsWith("/structural/composite"))
	fmt.Println(root.existsWith("/structural"))
	fmt.Println(root.existsWith("/creational"))
	fmt.Println(root.getSize())
	
	shouldExist, noError := root.getDirectory("/structural/composite")
	fmt.Println(noError)
	fmt.Println(shouldExist)
	
	shouldNotExist, error := root.getDirectory("/creational")
	fmt.Println(error)
	fmt.Println(shouldNotExist)
}
