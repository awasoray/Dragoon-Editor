package characters

import (
	. "LODeditor/internal/app/storage"
)

// Character : struct with available game attributes
type Character struct {
	ID        int
	Name      string
	XP        Attribute
	HP        Attribute
	MP        Attribute
	SP        Attribute
	SPMax     Attribute
	Level     Attribute
	DLevel    Attribute
	Weapon    Attribute
	Helmet    Attribute
	Chest     Attribute
	Boots     Attribute
	Accessory Attribute
}

// CreateCharacter : based off a root position in the save file
func CreateCharacter(ID int, name string, root int) Character {
	return Character{
		ID:        ID,
		Name:      name,
		XP:        Attribute{root, 4, true},
		HP:        Attribute{root + 8, 2, true},
		MP:        Attribute{root + 10, 2, true},
		SP:        Attribute{root + 12, 2, true},
		SPMax:     Attribute{root + 14, 2, true},
		Level:     Attribute{root + 18, 1, false},
		DLevel:    Attribute{root + 19, 1, false},
		Weapon:    Attribute{root + 20, 1, false},
		Helmet:    Attribute{root + 21, 1, false},
		Chest:     Attribute{root + 22, 1, false},
		Boots:     Attribute{root + 23, 1, false},
		Accessory: Attribute{root + 24, 1, false},
	}
}

// GetCharacters : returns splice of initiated characters
func GetCharacters() []Character {
	return []Character{
		Dart(), Shana(), Lavitz(), Rose(), Haschel(), Albert(), Miranda(), Meru(), Kongol(),
	}
}

// GetCharacterNames : returns splice of string for character names from initiated characters
func GetCharacterNames() []string {
	var r []string
	for _, c := range GetCharacters() {
		r = append(r, c.Name)
	}
	return r
}

// GetNameByID : Get character name by id from initiated characters
func GetNameByID(i int) string {
	r := ""
	for _, c := range GetCharacters() {
		if i == c.ID {
			r = c.Name
		}
	}
	return r
}

// GetIDByName : Get character id from name from initiated characters
func GetIDByName(n string) int {
	r := 0
	for _, c := range GetCharacters() {
		if n == c.Name {
			r = c.ID
		}
	}
	return r
}
