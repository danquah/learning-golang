package myexample

func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
