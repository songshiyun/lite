package array

func countQuadruplets(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}
	mp := make(map[int][]int)
	for i := 3; i < len(nums); i++ {
		mp[nums[i]] = append(mp[nums[i]], i)
	}
	var count int
	for i := len(nums) - 2; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				c := nums[i] + nums[j] + nums[k]
				v, ok := mp[c]
				if ok {
					count += containBigger(v, i)
				}
			}
		}
	}

	return count
}

func containBigger(src []int, dst int) (res int) {
	for _, item := range src {
		if item > dst {
			res++
		}
	}
	return
}
