package diamonds

func max(vals []int) int {
	max := 0
	for i := 0; i < len(vals); i++ {
		if vals[i] > max {
			max = vals[i]
		}
	}
	return max
}

func avg(vals []int) int {
	if len(vals) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(vals); i++ {
		total += vals[i]
	}
	return total / len(vals)
}
