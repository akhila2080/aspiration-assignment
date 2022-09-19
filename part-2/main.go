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

func NewEntity(l int, s string) Input {
	return Input{
		Entity:  strings.ToLower(s),
		Intervl: l,
		Counter: 1,
	}
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func (s *Input) TransformRune(pos int) {
	if isSpecialCharacter(string(s.Entity[pos])) {
		return
	}

	if s.Counter%s.Intervl == 0 {
		s.Entity = s.Entity[:pos] + strings.ToUpper(string(s.Entity[pos])) + s.Entity[pos+1:]
	}

	s.Counter++
}

func (s *Input) GetValueAsRuneSlice() []rune {
	return []rune(s.Entity)
}

func isSpecialCharacter(s string) bool {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.ReplaceAllString(s, " ") == " "
}
