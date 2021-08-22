package compaction

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(getLucky("leetcode", 2))
	//"334111"
	//[0,9,2,3,3,2,5,5,5,5]
	fmt.Println(maximumNumber1("132", []int{9, 8, 5, 0, 3, 6, 4, 2, 6, 8}))
	fmt.Println(maximumNumber1("021", []int{9, 4, 3, 5, 7, 2, 1, 9, 0, 6}))
	fmt.Println(maximumNumber1("334111", []int{0, 9, 2, 3, 3, 2, 5, 5, 5, 5}))
	//"214010"
	//[6,7,9,7,4,0,3,4,4,7]
	fmt.Println(maximumNumber1("214010", []int{6, 7, 9, 7, 4, 0, 3, 4, 4}))
}

func TestSplit(t *testing.T) {
	fmt.Println(Split(262124))
}
