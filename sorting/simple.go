package sorting

func Simple(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
