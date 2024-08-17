package flyweight

import "fmt"

type Army struct {
	warriors []Warrior
}

func createArmy() Army {
	return Army{
		warriors: []Warrior{},
	}
}

func (army *Army) addWarrior(warriorName string, typeName string, x int, y int) {
	warrior := Warrior{
		name: warriorName,
		warriorType: WarriorType{typeName: typeName, texture: Texture{}},
		position: Position{x: x, y: y},
	}

	fmt.Println("Adding new warrior:", warriorName)
	army.warriors = append(army.warriors, warrior)
}