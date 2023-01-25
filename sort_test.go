package main

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays bubble sort breaked", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSortBraked(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays bubble sort reversed", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSortReversed(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays bubble sort recursive", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			bubbleSortRecursive(ar, len(ar))
			b.StopTimer()
		}
	})
	b.Run("small arrays selection sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays selection sort by max", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			selectionSortByMax(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays selection sort bidirectional", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			selectionSortBidirectional(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays insertion sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays merge sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
	b.Run("small arrays quick sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays bubble sort breaked", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSortBraked(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays bubble sort reversed", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSortReversed(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays bubble sort recursive", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSortRecursive(ar, len(ar))
			b.StopTimer()
		}
	})
	b.Run("middle arrays selection sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays selection sort by max", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSortByMax(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays selection sort bidirectional", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSortBidirectional(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays insertion sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays merge sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
	b.Run("middle arrays quick sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays bubble sort breaked", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSortBraked(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays bubble sort reversed", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSortReversed(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays bubble sort recursive", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSortRecursive(ar, len(ar))
			b.StopTimer()
		}
	})
	b.Run("big arrays selection sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays selection sort by max", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSortByMax(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays selection sort bidirectional", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSortBidirectional(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays insertion sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays merge sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
	b.Run("big arrays quick sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
	b.Run("very small arrays quick sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 100)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
	b.Run("very big arrays quick sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 1000000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
	b.Run("very small arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 100)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("very big arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 1000000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
	b.Run("custom arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		ar := []int{1, 2, 3, 4, 5}
		b.StartTimer()
		bubbleSort(ar)
		b.StopTimer()
	})
	b.Run("custom arrays bubble sort", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		ar := []int{5, 4, 3, 2, 1}
		b.StartTimer()
		bubbleSort(ar)
		b.StopTimer()
	})
}

func generateSlice(max, cap int) []int {
	ar := make([]int, cap)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}
