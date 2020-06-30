package _struct

import "time"

type Person struct {
	Name string
	Age  int
	Time time.Time
}

func (p *Person) NewPerson(name string, age int) {
	p.Age = age
	p.Name = name
	p.Time = time.Now()
}
