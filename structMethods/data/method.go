package data

import "fmt"

func (i Instuctor) PrintInfo() string {
	return fmt.Sprintf("%v , %v (%v)", i.FirstName, i.LastName, i.Score)
}
