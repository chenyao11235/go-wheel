package binarysearch

//BinarySearch 二分查找  使用for循环实现
func BinarySearch(arr []int, value int) int {
	n := len(arr)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		} else if value < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//RecursiveBinarySearch 递归实现二分查找
func RecursiveBinarySearch(arr []int, value int) int {
	n := len(arr)
	return bs(arr, 0, value, n-1)
}

func bs(arr []int, low, value, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if value == arr[mid] {
		return arr[mid]
	} else if arr[mid] > value {
		return bs(arr, low, value, mid-1)
	} else {
		return bs(arr, mid+1, value, high)
	}
}
