package Algorithms

import "fmt"

// İlk for döngüsü sırasıyla her elemanı döner.
// İkinci for döngüsü, ilk for döngüsündeki index'te tutulan sayıdan daha küçük
// olan bir sayıyı tespit etmeye çalışır.
// Eğer kendisinden küçük bir sayıya denk gelirse ikisinin yerini değiştirir.

func selectionSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		var minimumId = i
		for j := i; j < len(s); j++ {
			// j := i
			// Daha önceden sıralanmış (gerideki) index'lere bakma.
			if s[j] < s[minimumId] {
				fmt.Printf("%d < %d \n", s[j], s[minimumId])
				minimumId = j
			}
		}
		fmt.Printf("%d ve %d yer değiştiriyor... i: %d ", s[i], s[minimumId], i)
		s[i], s[minimumId] = s[minimumId], s[i]
		fmt.Printf("Slice: %d \n", s)

	}
	return s
}

func SelectionSort() {
	var SliceToSort = []int{3, 6, 8, 2, 9, 5, 1, 7}
	var sortedSlice = selectionSort(SliceToSort)
	fmt.Println(sortedSlice)
}
