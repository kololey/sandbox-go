package sorting

// Bubble sorting O(n^2)
func Bubble(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

func SimpleBubbleSortDown(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				Swap(ar, j-1, j)
			}
		}
	}
}

func SimpleBubbleSortUp(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				Swap(ar, j-1, j)
			}
		}
	}
}

func Swap(ar []int, i, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}
