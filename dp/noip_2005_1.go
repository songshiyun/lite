package dp

/**
山洞里有 $M$ 株不同的草药，采每一株都需要一些时间 $t_i$，
每一株也有它自身的价值 $v_i$。给你一段时间 $T$，在这段时间里，
你可以采到一些草药。让采到的草药的总价值最大
*/
func maxCost(tcost, mget []int, t int) int {
	if len(tcost) == 0 || len(mget) == 0 || t <= 0 {
		return 0
	}
	ans := 0
	var dfs func(pos int, tleft int, ans int)
	dfs = func(pos int, tleft int, tans int) {
		if tleft < 0 {
			return
		}
		if pos == len(tcost) {
			ans = max(ans, tans)
			return
		}
		dfs(pos+1, tleft, tans)                          //不采集pos +1
		dfs(pos+1, tleft-tcost[pos+1], tans+mget[pos+1]) //采集pos+1,会消耗成本和获得收益
	}
	dfs(0, t, 0)
	return ans
}
