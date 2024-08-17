package flyweight

import (
	"testing"
)

func Test_Flyweight(t *testing.T) {
	army := createArmy()

	army.addWarrior("Hawkeye", "Archer", 13, 42)
	army.addWarrior("Doctor Strange", "Magic", 23, 12)
	army.addWarrior("Reed Richards", "Scientist", 65, 34)
	army.addWarrior("Scarlet Witch", "Magic", 93, 88)
	army.addWarrior("Doctor Doom", "Magic", 12, 76)
	army.addWarrior("Captain America", "Soldier", 12, 45)
	army.addWarrior("Hulk", "Soldier", 12, 45)
	army.addWarrior("Spider Man", "Soldier", 12, 45)
	army.addWarrior("Daredevil", "Soldier", 12, 45)
}
