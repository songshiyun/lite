package heap

import "testing"

func TestKn(t *testing.T) {
	x := kthLargestNumber([]string{"3","6","7","10"},4)
	t.Log(x)
}
