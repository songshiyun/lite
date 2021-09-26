package array

func maximumDifference(nums []int) int {
	// 最长递增子序列
	var stock []int
	var res int = -1
	for i := range nums {
		if len(stock) == 0 {
			stock = append(stock, i)
		}else {
			for {
				var tmp int
				if len(stock) > 0 {
					tmp = nums[stock[len(stock)-1]]
					if tmp >= nums[i] {
						if len(stock) > 1 {
							res = MaxInt(res,tmp-nums[stock[0]])
						}
						stock = stock[:len(stock)-1]
					}else {
						stock = append(stock, i)
						break
					}
				}else {
					stock = append(stock,i)
					break
				}
			}
		}
	}
	if len(stock) >= 2 {
		res = MaxInt(res,nums[stock[len(stock)-1]]-nums[stock[0]])
	}
	return res
}

func MaxInt(i,j int) int {
	if i < j {
		return j
	}
	return i
}