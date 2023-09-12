package sorting

func Select(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}
		if min != i {
			swap(ar, i, min)
		}
	}
}

func swap(ar []int, i, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}
