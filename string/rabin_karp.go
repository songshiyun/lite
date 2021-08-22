package string

const prime = 16777619

//
func Index(s, substr string) int {
	n := len(substr)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return indexBytes(s,substr[0])
	case n == len(s):
		if s == substr {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	return indexRabinKarp(s, substr)
}

func Contains(s, substr string) bool {
	return Index(s, substr) != -1
}

func indexBytes(str string,byte2 byte) int {
	for i,b := range str {
		if byte(b) == byte2 {
			return i
		}
	}
	return -1
}

func hashStr(s string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(s); i++ {
		hash = hash*prime + uint32(s[i])
	}
	var pow, sq uint32 = 1, prime
	for i := len(s); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func indexRabinKarp(s, substr string) int {
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*prime + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= prime
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}
