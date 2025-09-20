package _0045_jump_game_2

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	currentReach, maxReach, steps := nums[0], nums[0], 1
	for i := 0; i < len(nums); i++ {
		if currentReach >= len(nums)-1 {
			break
		}

		maxReach = max(maxReach, i+nums[i])
		if i == currentReach {
			currentReach = maxReach
			steps++
		}
	}

	return steps
}
