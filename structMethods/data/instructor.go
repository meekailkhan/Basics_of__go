package data

import "fmt"

type Instuctor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstuctor(name, lastname string) Instuctor {
	return Instuctor{FirstName: name, LastName: lastname}
}

func (c Course) String() string {
	return fmt.Sprintf("----%v----%v----", c.Name, c.Instuctor.FirstName)
}
