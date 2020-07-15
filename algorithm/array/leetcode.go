package array

/* 做过的leetcode中的题目
 */

// 寻找数据的中心索引
func pivotIndex(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return -1
	}
	var sum = 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	var preSum = 0
	for i := 0; i < n; i++ {
		if preSum*2+arr[i] == sum {
			return i
		}
		preSum += arr[i-1]
	}
	return -1
}
