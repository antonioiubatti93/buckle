package person

type Person struct {
	name string
	age  int
}

func (p Person) Name() string {
	return p.name
}

func (p Person) Age() int {
	return p.age
}

func NewPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}
