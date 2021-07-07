package util

func Add(first int, slic ...int) int {
	sum := first
	for _, val := range slic {
		sum += val
	}
	return sum
}
