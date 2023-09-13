package Algorithms

import "fmt"

// Bir sonraki elemanın, şu ankinden küçük olduğu ilk eşleşmeyi bul ve yer değiştir.
func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		//  i < len(array)-1
		// ÇÜNKÜ: İlk dönüş, en büyük elemanı tespit edip, en sona atıyor zaten.
		for j := 0; j < len(array)-i-1; j++ {
			//  j < len(array)-i-1
			// Daha önceden sıralanmış (ilerideki) index'lere bakmamalı.
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				fmt.Printf("%d ve %d yer değiştiriyor...\n", array[j], array[j+1])
			}
		}
		fmt.Printf("%d kez dönülecek. Mevcut slice: %d  - Sonraki dönüşe geçiliyor!\n", len(array)-i-1, array)
	}
	return array
}

func BubbleSort() {
	var SliceToSort = []int{3, 6, 8, 2, 9, 5, 1, 7}
	var sortedSlice = bubbleSort(SliceToSort)
	fmt.Println(sortedSlice)
}
