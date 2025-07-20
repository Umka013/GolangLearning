package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func quickSort(array []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := l + (r-l)/2
	pivotVal := array[pivot]

	i := l
	j := r

	for i <= j {
		for array[i] < pivotVal {
			i++
		}
		for array[j] > pivotVal {
			j--
		}
		if i <= j {
			swap(&array[i], &array[j])
			i++
			j--
		}
	}
	if l < j {
		quickSort(array, l, j)
	}
	if i < r {
		quickSort(array, i, r)
	}

}

func quickSort2(array []int, l int, r int) {
	pivot := l + (r-l)/2
	pivotVal := array[pivot]

	i := l
	j := r

	if i >= j {
		return
	}

	for i <= j {
		for array[i] < pivotVal {
			i++
		}
		for array[j] > pivotVal {
			j--
		}
		if i <= j {
			swap(&array[i], &array[j])
			i++
			j--
		}
	}

	if j > l {
		quickSort2(array, l, j)
	}
	if i < r {
		quickSort2(array, i, r)
	}
}

func main() {
	array := []int{5, 3, 1, 4, 5, 6, 2}
	quickSort2(array, 0, len(array)-1)
	fmt.Println(array)
}
