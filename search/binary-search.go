package search

func BinarySearchRecursive(list []int, target int, low int, high int) (int, bool) {
	middle := (low + high) / 2
	middleValue := list[middle]

	if middleValue == target {
		return middle, true
	}

	if middleValue > target {
		return BinarySearchRecursive(list, target, low, middle-1)
	}

	return BinarySearchRecursive(list, target, middle+1, high)
}

func BinarySearch(list []int, target int) (int, bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == target {
			return mid, true
		}
		if list[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, false
}
