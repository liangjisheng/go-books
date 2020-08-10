package search

// BinarySearch ...
func BinarySearch(array []int, target int) int {
	n := len(array)
	if n <= 0 {
		return -1
	}

	i := 0
	j := n - 1
	for i <= j {
		mid := (i + j) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
