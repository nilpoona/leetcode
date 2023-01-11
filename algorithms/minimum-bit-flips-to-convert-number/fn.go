package minimum_bit_flips_to_convert_number

// minBitFlips is 2220. Minimum Bit Flips to Convert Number
// SEE: https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/
func minBitFlips(start int, goal int) int {
	n := start ^ goal
	steps := 0
	for i := 0; i < 32; i++ {
		if n>>i&1 == 1 {
			steps++
		}
	}
	return steps
}
