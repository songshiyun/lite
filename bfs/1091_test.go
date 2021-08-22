package bfs

import (
	"log"
	"testing"
)

func TestT1(t *testing.T) {
	data := [][]int{{0, 1}, {1, 0}}
	res := shortestPathBinaryMatrix(data)
	log.Println(res)
}
