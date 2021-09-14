package heap

import "testing"

func Test1972(t *testing.T) {
	//[[1,2],[3,5],[2,2]],
	class := [][]int{{1,2},{3,5},{2,2}}
	t.Log(maxAverageRatio(class,2))

	class1 :=[][]int{{2,4},{3,9},{4,5},{2,10}}
	t.Log(maxAverageRatio(class1,4))
}
