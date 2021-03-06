package storage

import (
	"io/ioutil"
)

// Card : Main struct for the memory card
type Card struct {
	Path  string
	Slots []Slot
}

// LoadCard : Loads a memory card from file on disc
func LoadCard(path string) Card {
	f, err := ioutil.ReadFile(path)
	var s []Slot
	if err == nil {
		for i := 0; i <= 15; i++ {
			b := i * 0x2000
			t := b + 0x1FFF
			s = append(s, Slot{Head: b, Data: f[b:t]})
		}
	}
	return Card{
		Path:  path,
		Slots: s,
	}
}

// SaveCard : Saves the active memory card back to disc
func (c Card) SaveCard() {
	b := make([]byte, 0x20000)
	for _, slot := range c.Slots {
		for i, byte := range slot.Data {
			b[slot.Head+i] = byte
		}
	}
	err := ioutil.WriteFile(c.Path, b, 0644)
	if err != nil {
		panic(err)
	}
}
