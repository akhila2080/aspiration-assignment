package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type Input struct {
	Entity  string
	Intervl int
	Counter int
}

func main() {
	s := NewEntity(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s.Entity)
}

// NewEntity is used to build initial payload of an Input model.
func NewEntity(l int, s string) Input {
	return Input{
		Entity:  strings.ToLower(s),
		Intervl: l,
		Counter: 1,
	}
}

// MapString is used to process the runes using for loop.
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

// TransformRune is used to capitalize the rune if it meets all the requirements.
func (s *Input) TransformRune(pos int) {
	// check if the rune is special character.
	if isSpecialCharacter(string(s.Entity[pos])) {
		return
	}

	// if rune is at required interval, capitalize the rune.
	if s.Counter%s.Intervl == 0 {
		s.Entity = s.Entity[:pos] + strings.ToUpper(string(s.Entity[pos])) + s.Entity[pos+1:]
	}

	// increase the counter value.
	s.Counter++
}

// GetValueAsRuneSlice will return the slice of rune values from input string.
func (s *Input) GetValueAsRuneSlice() []rune {
	return []rune(s.Entity)
}

// specialCharacter is used to find if there is any special character in given string.
func isSpecialCharacter(s string) bool {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(s, " ") == " "
}
