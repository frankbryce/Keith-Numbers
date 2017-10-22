// Package keith has helper functions for dealing with Keith Numbers.  For more
// information about keith numbers, see https://youtu.be/uuMwz47LV_w
package keith

import (
	"container/list"
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
	nums := list.New()
	sum := big.NewInt(0)
	for _, d := range digits {
		b := big.NewInt(int64(d - '0'))
		nums.PushBack(b)
		sum.Add(sum, b)
	}

	for sum.Cmp(n) < 0 {
		sub := nums.Front().Value.(*big.Int)
		nums.Remove(nums.Front())
		nums.PushBack(big.NewInt(sum.Int64()))
		sum.Add(sum, sum)
		sum.Sub(sum, sub)
	}
	return sum.Cmp(n) == 0
}
