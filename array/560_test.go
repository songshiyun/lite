package array

import "testing"

func TestTT(t *testing.T) {
	t.Log(subarraySum([]int{1},0))
	t.Log(subarraySum1([]int{1},0))
	t.Log(subarraySum1([]int{1,1,1},2))
}
