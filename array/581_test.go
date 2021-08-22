package array

import "testing"

func TestName(t *testing.T) {
	a := []int{1, 3, 2, 2, 2}
	//a := []int{1,3,2,3,3}
	t.Log(findUnsortedSubarray(a))
	t.Log(findUnsortedSubarray1(a))
}

func TestX(t *testing.T) {
	a := []int{2, 2, 3, 4}
	t.Log(triangleNumber(a))
	t.Log(triangleNumber1(a))
}
