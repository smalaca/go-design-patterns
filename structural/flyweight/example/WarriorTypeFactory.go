package flyweight

import "fmt"

type WarriorTypeFactory struct {
	warriorTypes map[string]WarriorType
}

func createWarriorTypeFactory() WarriorTypeFactory {
	return WarriorTypeFactory{warriorTypes: make(map[string]WarriorType)}
}

func (factory *WarriorTypeFactory) create(typeName string) WarriorType {
	warriorType, exists := factory.warriorTypes[typeName]
	
	if (!exists) {
		fmt.Println("Create new warrior type:", typeName)
		warriorType = WarriorType{typeName: typeName, texture: Texture{}}
		factory.warriorTypes[typeName] = warriorType
	}

	return warriorType
}