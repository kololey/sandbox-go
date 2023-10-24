package basics

import "fmt"

func LoopVarFirst() {
	a := []int{1, 2, 3, 4, 5}
	b := []*int{}
	for _, i := range a {
		b = append(b, &i)
	}
	for _, j := range b {
		fmt.Printf("%d ", *j)
	}
}

// Output:
// 5 5 5 5 5

func LoopVarSecond() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Print(i)
		}()
	}
}

// Output:
//

func LoopVarThird() {
	numbers := []*int{}
	for i := 0; i < 5; i++ {
		numbers = append(numbers, &i)
	}
	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}

// Output:
// 5 5 5 5 5
