package jump_game

func jump(current int, nums []int, lastIndex int, visited map[int]bool) bool {
	if current == lastIndex {
		return true
	}
	if current > lastIndex {
		return false
	}

	if _, ok := visited[current]; ok {
		return false
	}

	visited[current] = true

	distance := nums[current]

	for distance > 0 {
		if result := jump(current+distance, nums, lastIndex, visited); result {
			return result
		}
		distance--
	}
	return false
}

func canJump(nums []int) bool {
	visited := make(map[int]bool)
	return jump(0, nums, len(nums)-1, visited)
}
