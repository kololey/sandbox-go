package basics

import "fmt"

func changePointer(p *int) {
	v := 3
	p = &v
}

func ChangePointerFirst() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	fmt.Println(*p)
}

// Output:
// 5
// 5
