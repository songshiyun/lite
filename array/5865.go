package array

func firstDayBeenInAllRooms(nextVisit []int) int {
	N := 10 ^ 5 + 1
	mod := int64(10 ^ 9 + 7)
	n := len(nextVisit)
	if n == 2 {
		return 2
	}
	times := make([]int64, N)
	sums := make([]int64, N)
	times[0], sums[0] = 0, 0
	times[1], sums[1] = 2, 2
	for i := 1; i+1 < n; i++ {
		if nextVisit[i] == i {
			times[i+1] = 2
			sums[i+1] = (sums[i] + times[i+1]) % mod
		} else {
			times[i+1] = (sums[i] - sums[nextVisit[i]] + 2) % mod
			sums[i+1] = (sums[i] + times[i+1]) % mod
		}
	}
	return int((sums[n-1]%mod + mod) % mod)
}

/**
const int N = 1e5 + 10;
long long f[N], g[N], h[N];
const int MM = 1e9 + 7;

class Solution {
public:
    int firstDayBeenInAllRooms(vector<int>& nxt) {
        int n = nxt.size();
        if (n == 2) return 2;
        f[0] = g[0] = h[0] = 0;
        f[1] = g[1] = h[1] = 2;
        for (int i = 1; i + 1 < n; i++) {
            // i -> i+1
            if (nxt[i] == i) {
                f[i + 1] = 2;
                g[i + 1] = (g[i] + f[i + 1]) % MM;
                // h[i + 1] = (h[i] + 2) % MM;
            } else {
                f[i + 1] = (g[i] - g[nxt[i]] + 2) % MM;
                g[i + 1] = (g[i] + f[i + 1]) % MM;
                // h[i + 1] = (h[i] + f[i + 1]) % MM;
            }
        }
        return (g[n - 1] % MM + MM) % MM;
    }
};
*/
