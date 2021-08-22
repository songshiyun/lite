package array

import "sort"

//2,2,3,4
//超时
func triangleNumber(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	var temp [][]int
	var res int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				t := []int{nums[i], nums[j], nums[k]}
				if is(nums[i], nums[j], nums[k]) {
					temp = append(temp, t)
					res++
				}
			}
		}
	}

	return res
}

func is(i, j, k int) bool {
	if i+j > k && i+k > j && j+k > i && i-j < k && i-k < j && j-i < k && j-k < i && k-i < j && k-j < i {
		return true
	}
	return false
}

func triangleNumber1(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	// 排序，然后逆序遍历两边之和大于第三变的数据
	sort.Ints(nums)
	var res int
	for k := len(nums) - 1; k >= 2; k-- {
		l, r := 0, k-1
		//遍历所有的l->r
		// 二分查找
		for l < r {
			// 如果满足num[l] + num[r] > num[k]
			// 则满足 num[l->r-1]+num[r] > num[k]
			// 然后满足r--
			if nums[l]+nums[r] > nums[k] {
				res += r - l
				r--
			} else {
				// 如果相加小于，则应该移动左边部分
				l++
			}
		}
	}
	return res
}
