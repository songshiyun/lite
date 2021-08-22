package graph

import "testing"

func TestNameT(t *testing.T) {

	src := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	//	4
	//	2
	t.Log(networkDelayTime(src, 4, 2))

}
