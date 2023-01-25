package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	ar = quickSort(ar)

	fmt.Println(ar)
}

func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	q := ar[rand.Intn(len(ar)-1)]
	s_nums := make([]int, 0)
	m_nums := make([]int, 0)
	e_nums := make([]int, 0)
	for _, n := range ar {
		if n < q {
			s_nums = append(s_nums, n)
		} else if n > q {
			m_nums = append(m_nums, n)
		} else {
			e_nums = append(e_nums, n)
		}
	}
	return append(append(quickSort(s_nums), e_nums...), quickSort(m_nums)...)
}

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	k := len(ar) / 2
	a1 := ar[:k]
	a2 := ar[k:]
	left := mergeSort(a1)
	right := mergeSort(a2)
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		t := ar[i]
		ti := i
		for j := i - 1; j >= 0; j-- {
			if t < ar[j] {
				ar[j+1] = ar[j]
				ti--
			}
		}
		if ti != i {
			ar[ti] = t
		}
	}
}

func selectionSortBidirectional(ar []int) {
	for i := 0; i < len(ar)/2; i++ {
		minIndex := i
		maxIndex := len(ar) - 1 - i
		for j := i; j < len(ar)-i; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		if minIndex != i {
			ar[minIndex], ar[i] = ar[i], ar[minIndex]
			if maxIndex == i {
				maxIndex = minIndex
			}
		}
		if maxIndex != len(ar)-1-i {
			ar[maxIndex], ar[len(ar)-1-i] = ar[len(ar)-1-i], ar[maxIndex]
		}
	}
}

func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i >= 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			ar[maxIndex], ar[i] = ar[i], ar[maxIndex]
		}
	}
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			ar[minIndex], ar[i] = ar[i], ar[minIndex]
		}
	}
}

func bubbleSortRecursive(ar []int, n int) {
	for i := 0; i < n-1; i++ {
		if ar[i] > ar[i+1] {
			ar[i+1], ar[i] = ar[i], ar[i+1]
		}
	}
	if n-1 > 1 {
		bubbleSortRecursive(ar, n-1)
	}
}

func bubbleSortReversed(ar []int) {
	swap := false
	for i := 0; i < len(ar)-1; i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j] < ar[j-1] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

func bubbleSortBraked(ar []int) {
	swap := false
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

func bubbleSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
			}
		}
	}
}
