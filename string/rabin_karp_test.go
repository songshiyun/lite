package string

import "testing"

func TestIndex2(t *testing.T) {
	idx := indexRabinKarp("helllo","ll")
	t.Log(idx)
}
