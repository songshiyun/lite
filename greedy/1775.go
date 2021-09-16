package greedy

import "sort"

func minOperations(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	// 不可能的情况是 较短的每个都变成最大都不能满足，较长的变为最小的和
	if min(l1, l2)*6 < max(l1, l2) {
		return -1
	}
	s1 := sum(nums1)
	s2 := sum(nums2)
	if s1 == s2 {
		return 0
	}
	// 假设 s1 < s2, 则变换之后的和则满足： s1 < s < s2
	// s1 sum小
	// s2 sum大
	// 保证s1 < s2
	if s1 > s2 {
		return minOperations(nums2, nums1)
	}
	// 排序，将num1排序成： 1，2，3，4，5从小到大大顺序，每次变换选择最小的变换 1 -> 6, 直接变换到最大，这样能够操作最小的步骤，1->2->3...
	// 排序，将nums排序成： 4，3，2，1，每次从最大的开始变换，挑选最大的变换，则也能够最优的步骤
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})
	var i, j, ans int
	for s1 != s2 {
		var d1, d2 int
		diff := s2 - s1
		if i != l1 {
			// 如果已经遍历完了nums1数组，
			// 每次都贪心的变换最大
			d1 = 6 - nums1[i]
		}
		if j != l2 {
			// 如果已经便利完了nums2数组
			d2 = nums2[j] - 1
		}
		// 每次变换都选取较大的
		d := min(diff, max(d1, d2))
		if d1 >= d2 {
			s1 += d
			i++
		} else {
			s2 -= d
			j++
		}
		ans++
	}
	return ans
}

func sum(nums []int) int {
	var result int
	for _, item := range nums {
		result += item
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
