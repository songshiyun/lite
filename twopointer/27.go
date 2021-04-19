package tp

//给定一个数组nums和一个val,原地移除数组中等于val的元素，返回移除后的数组的长度

//solution :
//1.暴力解，两层for循环
//2.双指针（slow & fast pointer）
func removeVal(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

//双指针解法： 排序数组中原地移除重复元素
//the same as : 移除链表中的重复元素
// [0,1,1,1,2,2,3,3,3,4] -> [0,1,3,4]
