package data

import "time"

type WorkShop struct {
	Course
	Instuctor
	Date time.Time
}

func NewWorkShop(name string, instructor Instuctor) WorkShop {
	w := WorkShop{}
	w.Name = name
	w.Instuctor = instructor
	w.Course.Instuctor.FirstName = w.Instuctor.FirstName
	return w
}

func (c WorkShop) Signup() bool {
	return true
}
