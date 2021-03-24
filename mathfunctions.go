package mathfunctions

func Min(x ...float64) float64 {
	min := x[0]
	for _, value := range x {
		if value < min {
			min = value
		}
	}
	return min
}
