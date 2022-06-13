package number_of_1_bits

import (
	"strconv"
	"strings"
)

func hammingWeight(num uint32) int {
	b := strconv.FormatUint(uint64(num), 2)
	return strings.Count(b, "1")
}
