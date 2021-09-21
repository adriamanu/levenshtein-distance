package levenshtein_distance

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}

func Levenshtein(a, b string) int {
	// Initialize distance array
	var d = make([][]int, len(a)+1)
	for i := range d {
		d[i] = make([]int, len(b)+1)
	}

	// Initialize first row and first column
	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			var substitution int = 1
			if a[i-1] == b[j-1] {
				substitution = 0
			}
			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+substitution)
		}
	}

	return d[len(a)][len(b)]
}
