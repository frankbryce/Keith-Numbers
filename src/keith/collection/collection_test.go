package collection_test

import (
	"fmt"
	"keith"
	coll "keith/collection"
	"math/big"
	"strings"
	"testing"
)

type testCase struct {
	n         *big.Int
	isKeith   bool
	last      *big.Int
	underflow *big.Int
	overflow  *big.Int
	count     uint64
}

func bigFromString(s string) *big.Int {
	n := big.NewInt(0)
	n.SetString(s, 10)
	return n
}

func TestKeithCollection(t *testing.T) {
	tests := []testCase{
		// 1,0,1,1,2,3,5,8,13
		testCase{
			n:         bigFromString("10"),
			isKeith:   false,
			last:      bigFromString("13"),
			underflow: bigFromString("2"),
			overflow:  bigFromString("3"),
			count:     9,
		},
		// 1,4,5,9,14
		testCase{
			n:         bigFromString("14"),
			isKeith:   true,
			last:      bigFromString("14"),
			underflow: bigFromString("5"),
			overflow:  bigFromString("0"),
			count:     5,
		},
		// 1,9,7,17,33,57,107,197
		testCase{
			n:         bigFromString("197"),
			isKeith:   true,
			last:      bigFromString("197"),
			underflow: bigFromString("90"),
			overflow:  bigFromString("0"),
			count:     8,
		},
		// 1,9,8,18,35,61,114,210
		testCase{
			n:         bigFromString("198"),
			isKeith:   false,
			last:      bigFromString("210"),
			underflow: bigFromString("84"),
			overflow:  bigFromString("12"),
			count:     8,
		},
	}

	for _, test := range tests {
		isKeith := keith.IsKeithCollect(test.n)
		if isKeith != test.isKeith {
			t.Errorf("case %v \"isKeith\", want %v got %v", test.n, test.isKeith, isKeith)
		}
		for _, c := range coll.Collections {
			s := fmt.Sprintf("%v", c.Get())
			switch strings.TrimPrefix(fmt.Sprintf("%T", c), "*collection.") {
			case "last":
				if test.last.String() != s {
					t.Errorf("case %v \"last\", want %v got %v", test.n, test.last, s)
				}
			case "underflow":
				if test.underflow.String() != s {
					t.Errorf("case %v \"underflow\", want %v got %v", test.n, test.underflow, s)
				}
			case "overflow":
				if test.overflow.String() != s {
					t.Errorf("case %v \"overflow\", want %v got %v", test.n, test.overflow, s)
				}
			case "count":
				if test.count != c.Get() {
					t.Errorf("case %v \"count\", want %v got %v", test.n, test.count, s)
				}
			}
		}
	}
}
