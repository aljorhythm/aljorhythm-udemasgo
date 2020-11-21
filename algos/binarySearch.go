package algos

import "fmt"

func _binarySearchRecursive(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := int((low + high) / 2)
	fmt.Println(low, high, mid)

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return _binarySearchRecursive(arr, target, low, mid-1)
	} else {
		return _binarySearchRecursive(arr, target, mid+1, high)
	}
}

// BinarySearchRecursive searches target in sorted array recursively
func BinarySearchRecursive(arr []int, target int) int {
	return _binarySearchRecursive(arr, target, 0, len(arr)-1)
}

// BinarySearchIterative searches target in sorted array recursively
func BinarySearchIterative(arr []int, target int) int {
	low := 0
	high := len(arr)
	mid := int((low + high) / 2)

	for low < high {
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
