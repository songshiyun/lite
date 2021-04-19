package array

//遍历找到第一个大于目标值的数据
//
func findPosByTraverse(src []int, target int) (pos int) {
	for i := 0; i < len(src)-1; i++ {
		if src[i] >= target {
			return i
		}
	}
	return len(src)
}

//4种情况
//1 在最左侧
//2 在最右侧
//3 等于其中的某个元素
//4 位于其中两个元素之间
//定义在其中的一个 [left,right]区间中
func searchInsert(src []int, target int) (pos int) {
	length := len(src)
	left := 0
	right := length - 1

	for left <= right { //left == right -> [left,right]仍然有效
		mid := left + ((right - left) / 2) //防止溢出
		if src[mid] > target {
			right = mid - 1
		} else if src[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right + 1
}

//定义在其中的一个 [left,right)区间中
func searchInsert2(src []int, target int) (pos int) {
	length := len(src)
	left := 0
	right := length
	for left < right { //left == right -> [left,right)无效
		mid := left + ((right - left) / 2) //防止溢出
		if src[mid] > target {             //[left,mid)
			right = mid
		} else if src[mid] < target { //[mid+1,right)
			left = mid + 1
		} else { // target == src[mid]
			return mid
		}
	}
	return right
}
