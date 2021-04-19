package array

/**
当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：
若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。
返回 A 的最大湍流子数组的长度。
示例 1：
     a[k] > a[k+1] & a[k] < a[k+1]
      0 1 2 3  4 5 6 7 8
输入：[9,4,2,10,7,8,8,1,9]
输出：5
解释：(A[1] > A[2] < A[3] > A[4] < A[5])
示例 2：
      a[k] > a [k+1]k为偶数 & a[k] < a[k+1] k为奇数
      0 1  2   3
输入：[4,8,12,16]
输出：2
示例 3：
输入：[100]

dp[k] = 1 init
k为奇数： dp[k] = max(dp[k],dp[k-1]+1)

输出：1

提示：
1 <= A.length <= 40000
0 <= A[i] <= 10^9
*/

//双指针
func maxTurbulenceSize(arr []int) int {
	result := 1 //默认至少一个满足
	left, right := 0, 0
	for right < len(arr)-1 {
		if left == right {
			if arr[left] == arr[right] {
				left++
			}
			right++
		} else {
			//arr[right-1] < arr[right] && arr[right] > arr[right+1]
			//arr[right-1] > arr[right] && arr[right] < arr[right] -> right++
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				//不满足直接left = right
				left = right
			}
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
