package main

import (
	"sort"
	"strings"
)

func main() {

}

func DoPage(data []*Item, pageSize int) [][]*Item {
	var result [][]*Item
	if len(data) == 0 {
		return result
	}
	globalQueue := make([]*Item, 0, 0)
	//先判断global中的数据
	//再判断data
	//如果存在则放入global中
	//[1,1,1] -> page_size=3 -> 3页
	for len(globalQueue) != 0 || len(data) != 0 {
		var temp *Row
		temp, globalQueue = makeDataFromGlobal(globalQueue, pageSize)
		if len(temp.Items) == pageSize {
			continue
		}
		lenght := len(temp.Items)
		for index, item := range data {
			if lenght < pageSize {
				if exist(temp.ItemSet, item.HostID) {
					globalQueue = append(globalQueue, item)
				} else {
					temp.Items = append(temp.Items, item)
					temp.ItemSet[item.HostID] = struct{}{}
					lenght++
				}
			} else {
				data = data[index:]
				break
			}
		}
		result = append(result, temp.Items)
	}
	for _, item := range result {
		sort.Slice(item, func(i, j int) bool {
			return item[i].Score < item[j].Score
		})
	}
	return result
}

func makeDataFromGlobal(global []*Item, size int) (*Row, []*Item) {
	result := &Row{
		Items:   make([]*Item, 0, 0),
		ItemSet: map[int]struct{}{},
	}
	g1 := make([]*Item, 0, 0) //需要拷贝一次，update in-place
	for _, item := range global {
		if exist(result.ItemSet, item.HostID) {
			g1 = append(g1, item)
		} else {
			if len(result.Items) < size {
				result.Items = append(result.Items, item)
				result.ItemSet[item.HostID] = struct{}{}
			} else {
				g1 = append(g1, item)
			}
		}
	}
	return result, g1
}

func exist(temp map[int]struct{}, id int) bool {
	for k, _ := range temp {
		if k == id {
			return true
		}
	}
	return false
}

type Pages struct {
	Rows []Row
}

type Row struct {
	Items   []*Item
	ItemSet map[int]struct{}
}

type Item struct {
	HostID    int
	ListingID int
	Score     float64
	City      string
}

type Items []Item

func (item Items) Less(i, j int) bool {
	return item[i].Score < item[j].Score
}

func (item Items) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}

func (item Items) Len() int {
	return len(item)
}

/**
<a          -> parse error
<></>       -> prase error
<<a></a>    -> parse error
<a>         -> missing close tag of <a>
<a><b></b>  -> missing close tag of <a>
</a>        -> missing open tag of </a>
<a><b></a></b>  -> missing open tag of </a>
<a></a>         -> valid
this is <a>a</a> <b>test</b> -> valid
*/
func xmlValidator(string2 string) bool {
	if strings.Trim(string2, " ") == "" {
		return true
	}
	var stack []byte
	for string2 != "" {

	}
	if len(stack) != 0 {
		return false
	}
	return true
}
