package _026_remove_duplicates

func removeDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[index-1] != nums[i] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
