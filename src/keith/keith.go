// Package keith has helper functions for dealing with Keith Numbers.  For more
// information about keith numbers, see https://youtu.be/uuMwz47LV_w
package keith

import (
	"math/big"
)

func IsKeith(n *big.Int) bool {
	return IsKeithBase(n, 10)
}

func IsKeithBase(n *big.Int, base int) bool {
	if base != 10 {
		return false
	}

	digits := n.String()
	nums := make([]*big.Int, len(digits))
	for i, d := range digits {
		nums[i] = big.NewInt(int64(d - '0'))
	}

	sum := big.NewInt(0)
	for i := 0; i < len(nums); i++ {
		sum.Add(sum, nums[i])
	}

	for sum.Cmp(n) < 0 {
		sub := nums[0]
		nums = append(nums, big.NewInt(sum.Int64()))[1:]
		sum.Add(sum, sum)
		sum.Sub(sum, sub)
	}
	return sum.Cmp(n) == 0
}
