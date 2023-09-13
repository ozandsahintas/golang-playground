package Algorithms

import "fmt"

func TestSort() {
	var SliceToSort = []int32{2, 9, 6, 8, 3}
	_testSort(5, SliceToSort)
}

func _testSort(n int32, arr []int32) {
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}
		fmt.Println(arr)
	}
}

//var lastItem = arr[n-1]
//for i := n - 2; i >= 0; i-- {
//if lastItem < arr[i] {
//lastItem, arr[i] = arr[i], lastItem
//}
//fmt.Println(arr)
//}

//for i := n - 2; i >= 0; i-- {
//if arr[n-1] < arr[i] {
//arr[n-1], arr[i] = arr[i], arr[n-1]
//}
//fmt.Println(arr)
//}
