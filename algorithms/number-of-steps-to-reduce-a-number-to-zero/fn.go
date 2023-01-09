package number_of_steps_to_reduce_a_number_to_zero

// numberOfSteps is 1342. Number of Steps to Reduce a Number to Zero
// SEE: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps(num int) int {
	steps := 0
	for {
		if num == 0 {
			break
		}
		steps++
		if num&1 == 1 {
			num--
			continue
		}
		num = num >> 1
	}
	return steps
}
