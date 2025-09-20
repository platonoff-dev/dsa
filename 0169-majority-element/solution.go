package _169_majority_element

func majorityElement(nums []int) int {
	ncounter, n := 1, nums[0]

	for _, num := range nums[1:] {
		if num == n {
			ncounter++
		} else {
			ncounter--
		}

		if ncounter == 0 {
			ncounter = num
		}
	}

	return n
}
