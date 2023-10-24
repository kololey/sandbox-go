package basics

import "fmt"

func AppendFirst() {
	var x []int
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}

// Output:
// [0 1 2 4] [0 1 2 4]

func AppendSecond() {
	a := []int{1, 2, 3}
	b := a
	b = append(b, 4)
	c := b
	b[0] = 0
	e := append(c, 5)
	b[2] = 7

	fmt.Println(a, b, c, e)
}

// Output:
// [1 2 3] [0 2 7 4] [0 2 7 4] [0 2 7 4 5]

func AppendThird() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"

	fmt.Printf("%s\n", a)
}

// Output:
// [a q c]

func AppendFourth() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd')
	b[0] = 'z'

	fmt.Printf("%s\n", a)
}

// Output:
// [a z d]
