package sorting

func Merge(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	m := len(data) / 2
	l := Merge(data[:m])
	r := Merge(data[m:])
	return mergeSortCombine(l, r)
}

func mergeSortCombine(l, r []int) []int {
	tmp := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		switch {
		case len(l) == 0:
			return append(tmp, r...)
		case len(r) == 0:
			return append(tmp, l...)
		case l[0] <= r[0]:
			tmp = append(tmp, l[0])
			l = l[1:]
		case l[0] >= r[0]:
			tmp = append(tmp, r[0])
			r = r[1:]
		}
	}
	return tmp
}
