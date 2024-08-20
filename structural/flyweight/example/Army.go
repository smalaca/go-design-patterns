package flyweight

import "fmt"

type Army struct {
	warriorTypeFactory WarriorTypeFactory
	warriors []Warrior
}

func createArmy() Army {
	return Army{
		warriorTypeFactory: createWarriorTypeFactory(),
		warriors: []Warrior{},
	}
}

func (army *Army) addWarrior(warriorName string, typeName string, x int, y int) {
	warriorType := army.warriorTypeFactory.create(typeName)
	warrior := Warrior{
		name: warriorName,
		warriorType: warriorType,
		position: Position{x: x, y: y},
	}

	fmt.Println("Adding new warrior:", warriorName)
	army.warriors = append(army.warriors, warrior)
}