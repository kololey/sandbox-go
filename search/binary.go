package search

import (
	"errors"
)

func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, errors.New("not found")
	}
	mid := lowIndex + (highIndex-lowIndex)/2
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}
