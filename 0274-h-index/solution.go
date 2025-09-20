package _0274_h_index

import "slices"

func hIndex(citations []int) int {
	slices.Sort(citations)
	hIndex := 0
	pos := len(citations) - 1
	for pos >= 0 && citations[pos] > hIndex {
		pos--
		hIndex++
	}

	return hIndex
}
