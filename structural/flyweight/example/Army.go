package flyweight

type Army struct {
	warriors []Warrior
}

func createArmy() Army {
	return Army{
		warriors: []Warrior{},
	}
}

func (army *Army) addWarrior(warriorName string, typeName string, x int, y int) {
	
}