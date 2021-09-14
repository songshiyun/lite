package string

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {

	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))

}
