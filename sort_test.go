package sort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T){
	numList := []int {1, 10, 2, 20, 3, 2, 0, -3, 3, 3, -10, 4, 20, -20, -10}
	bubbleList := Bubble(numList)
	fmt.Println("bubble sort :" , bubbleList)

	SelectList := Select(numList)
	fmt.Println("selection sort :" , SelectList)

	InsertList := Insert(numList)
	fmt.Println("Insertion sort :" , InsertList)

	MergeList := Merge(numList)
	fmt.Println("Merge sort :", MergeList)

	QuickList := Quick(numList)
	fmt.Println("Quick sort :", QuickList)
}