// Package keith has helper functions for dealing with Keith Numbers.  For more
// information about keith numbers, see https://youtu.be/uuMwz47LV_w
package keith

import (
	"math/big"
	"strconv"
)

func IsKeith(n int) bool {
	return IsKeithBase(n, 10)
}

func IsKeithBase(n, base int) bool {
	if base != 10 {
		return false
	}

	digits := strconv.Itoa(n)
	nums := make([]*big.Int, len(digits))
	for i, d := range digits {
		nums[i] = big.NewInt(int64(d - '0'))
	}

	sum := big.NewInt(0)
	for i := 0; i < len(nums); i++ {
		sum.Add(sum, nums[i])
	}

	N := big.NewInt(int64(n))
	for sum.Cmp(N) < 0 {
		sub := nums[0]
		nums = append(nums, big.NewInt(sum.Int64()))[1:]
		sum.Add(sum, sum)
		sum.Sub(sum, sub)
	}
	return sum.Cmp(N) == 0
}
