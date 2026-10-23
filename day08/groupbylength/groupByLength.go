package groupbylength

func GroupByLength(words []string) map[int][]string {

	m := make(map[int][]string)

	for _, w := range words {
		n := len(w)
		m[n] = append(m[n], w)
	}
	return m
}
