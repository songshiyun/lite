package array

import "math"

func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	mp := make(map[item]map[int][]item)
	for i := 0; i < n-1; i++ {
		for j := i +1; j < n; j++ {
			s := distance(points[i],points[j])
			it1 := item{
				x: points[i][0],
				y: points[i][1],
			}
			it2 := item{
				x: points[j][0],
				y: points[j][1],
			}
			if _,ok := mp[it1]; !ok {
				mp[it1] = make(map[int][]item)
			}
			mp[it1][s] = append(mp[it1][s], it2)

			if _,ok := mp[it2]; !ok {
				mp[it2] = make(map[int][]item)
			}
			mp[it2][s] = append(mp[it2][s], it1)
		}
	}

	for _,_ = range mp {



	}

	return 0

}

type item struct {
	x int
	y int
}
func distance(i,j []int) int {
	x := abs(i[0],j[0])
	y := abs(i[0],j[0])
	return x*x + y *y
}

func abs(i,j int) int {
	if i >= 0 && j >=0 || i <= 0 && j <= 0{
		return int(math.Abs(float64(i-j)))
	}
	return int(math.Abs(float64(i))) + int(math.Abs(float64(j)))
}