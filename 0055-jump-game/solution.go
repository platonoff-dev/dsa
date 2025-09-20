package _0055_jump_game

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		maxReach = max(i+nums[i], maxReach)

		if i == maxReach && nums[i] == 0 {
			break
		}
	}

	return maxReach >= len(nums)-1
}
