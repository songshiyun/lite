package array

import "fmt"

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(findSubSequence(nums, 7))
}

func findSubSequence(nums []int, target int) (result int) {
	for i := 0; i < len(nums)-1; i++ {
		sum := 0
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				result = min(result, j-i+1)
				break
			}
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
