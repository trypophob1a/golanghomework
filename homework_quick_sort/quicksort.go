package homeworkquicksort

func QuickSort(arr []int) []int {
	return fullSorted(arr, 0, len(arr)-1)
}

func fullSorted(arr []int, l, r int) []int {
	if l >= r {
		return arr
	}

	index := partition(arr, l, r)
	fullSorted(arr, index, r)

	return fullSorted(arr, l, index-1)
}

func partition(arr []int, l, r int) int {
	pivotPosition := (l + r) / 2
	pivot := arr[pivotPosition]

	for l <= r {
		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	return l
}
