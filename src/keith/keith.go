// Package keith has helper functions for dealing with Keith Numbers.  For more
// information about keith numbers, see https://youtu.be/uuMwz47LV_w
package keith

import (
	"container/list"
	coll "keith/collection"
	"math/big"
)

func IsKeith(n *big.Int) bool {
	return IsKeithDo(n, func(_ *big.Int) {})
}

func IsKeithCollect(n *big.Int) bool {
	for _, c := range coll.Collections {
		c.Reset()
		c.Set(n)
	}

	return IsKeithDo(n, func(i *big.Int) {
		for _, c := range coll.Collections {
			c.Add(i)
		}
	})
}

func IsKeithDo(n *big.Int, f func(*big.Int)) bool {
	digits := n.String()
	nums := list.New()
	sum := big.NewInt(0)
	for _, d := range digits {
		b := big.NewInt(int64(d - '0'))
		nums.PushBack(b)
		sum.Add(sum, b)
		f(sum)
	}

	f(sum)

	for sum.Cmp(n) < 0 {
		sub := nums.Remove(nums.Front()).(*big.Int)
		nums.PushBack(big.NewInt(0).Set(sum))
		sum.Add(sum, sum)
		sum.Sub(sum, sub)
		f(sum)
	}

	return sum.Cmp(n) == 0
}
