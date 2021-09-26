package array

import "testing"

func TestGridGame(t *testing.T) {
	t.Log(gridGame([][]int{{2,5,4},{1,5,1}}))
	t.Log(gridGame([][]int{{3,3,1},{8,5,2}}))
	t.Log(gridGame([][]int{{20,3,20,17,2,12,15,17,4,15},{20,10,13,14,15,5,2,3,14,3}}))
	t.Log(gridGame([][]int{{3,3,1},{8,5,2}}))
	t.Log(gridGame([][]int{{1,2,3},{3,2,1}}))

	t.Log(gridGame1([][]int{{2,5,4},{1,5,1}}))
	t.Log(gridGame1([][]int{{3,3,1},{8,5,2}}))
	t.Log(gridGame1([][]int{{20,3,20,17,2,12,15,17,4,15},{20,10,13,14,15,5,2,3,14,3}}))
	t.Log(gridGame1([][]int{{3,3,1},{8,5,2}}))
	t.Log(gridGame1([][]int{{1,2,3},{3,2,1}}))
}
