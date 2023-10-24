package basics

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Olga",
	}
}

func ChangePointersSecond() {
	person := &Person{
		Name: "Eugene",
	}
	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

// Output:
// Eugene
// Eugene
