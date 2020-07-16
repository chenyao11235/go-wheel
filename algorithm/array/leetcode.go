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

// 搜索插入位置 在一个排序数组中找到指定元素的索引，如果不存在就返回应该插入的位置 这种是暴力解题 时间复杂度是O(N)
// 可以使用二分查找 时间负载度是log(N)
func searchInsert(nums []int, target int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if target <= nums[i] {
			return i
		}
	}
	return n
}
