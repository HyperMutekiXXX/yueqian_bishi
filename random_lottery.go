package main

import (
	"math/rand"
)

// RandomLottery 随机抽奖，积分越大中奖率越高,时间复杂度O(n + log n)
func RandomLottery(arr []int) int {
	// 对数组进行切分
	length := len(arr)
	pre := make([]int, length)
	pre[0] = arr[0]
	for i := 1; i < length; i++ {
		pre[i] = arr[i] + pre[i-1]
	}

	x := rand.Intn(pre[length-1]) + 1
	// 寻找随机数 x <= pre[i],这时i就是随机数下标,使用二分查找
	return binarySearch(pre, x)
}

// binarySearch 二分查找
func binarySearch(arr []int, target int) int {
	length := len(arr)
	low, high := 0, length-1
	for low < high {
		mid := (high-low)/2 + low
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
