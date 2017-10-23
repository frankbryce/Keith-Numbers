// Package keith has helper functions for dealing with Keith Numbers.  For more
// information about keith numbers, see https://youtu.be/uuMwz47LV_w
package keith

import (
	"container/list"
	"math/big"
)

func IsKeith(n *big.Int) bool {
	digits := n.String()
	nums := list.New()
	sum := big.NewInt(0)
	for _, d := range digits {
		b := big.NewInt(int64(d - '0'))
		nums.PushBack(b)
		sum.Add(sum, b)
	}

	for sum.Cmp(n) < 0 {
		sub := nums.Remove(nums.Front()).(*big.Int)
		nums.PushBack(big.NewInt(sum.Int64()))
		sum.Add(sum, sum)
		sum.Sub(sum, sub)
	}
	return sum.Cmp(n) == 0
}
