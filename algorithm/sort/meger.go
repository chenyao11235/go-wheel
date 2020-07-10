package sort

//MergeSort 归并排序
func MergeSort(a []int) {
	n := len(a)
}

func sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	sort(arr, start, mid)
	sort(arr, mid+1, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr
}
