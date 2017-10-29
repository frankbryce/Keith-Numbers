package keith_test

import (
	"keith"
	"math/big"
	"testing"
)

type testCase struct {
	num     *big.Int
	isKeith bool
}

func bigFromString(s string) *big.Int {
	n := big.NewInt(0)
	n.SetString(s, 10)
	return n
}

func TestIsKeith(t *testing.T) {
	testCases := []testCase{
		testCase{num: big.NewInt(10), isKeith: false},
		testCase{num: big.NewInt(14), isKeith: true},
		testCase{num: big.NewInt(19), isKeith: true},
		testCase{num: big.NewInt(28), isKeith: true},
		testCase{num: big.NewInt(47), isKeith: true},
		testCase{num: big.NewInt(61), isKeith: true},
		testCase{num: big.NewInt(75), isKeith: true},
		testCase{num: big.NewInt(197), isKeith: true},
		testCase{num: big.NewInt(742), isKeith: true},
		testCase{num: big.NewInt(1104), isKeith: true},
		testCase{num: big.NewInt(1537), isKeith: true},
		testCase{num: big.NewInt(2208), isKeith: true},
		testCase{num: big.NewInt(2580), isKeith: true},
		testCase{num: big.NewInt(3684), isKeith: true},
		testCase{num: big.NewInt(4788), isKeith: true},
		testCase{num: big.NewInt(7385), isKeith: true},
		testCase{num: big.NewInt(7647), isKeith: true},
		testCase{num: big.NewInt(7909), isKeith: true},
		testCase{num: big.NewInt(31331), isKeith: true},
		testCase{num: big.NewInt(34285), isKeith: true},
		testCase{num: big.NewInt(34348), isKeith: true},
		testCase{num: big.NewInt(55604), isKeith: true},
		testCase{num: big.NewInt(62662), isKeith: true},
		testCase{num: big.NewInt(86935), isKeith: true},
		testCase{num: big.NewInt(93993), isKeith: true},
		testCase{num: big.NewInt(120284), isKeith: true},
		testCase{num: big.NewInt(129106), isKeith: true},
		testCase{num: big.NewInt(147640), isKeith: true},
		testCase{num: big.NewInt(156146), isKeith: true},
		testCase{num: big.NewInt(174680), isKeith: true},
		testCase{num: big.NewInt(183186), isKeith: true},
		testCase{num: big.NewInt(298320), isKeith: true},
		testCase{num: big.NewInt(355419), isKeith: true},
		testCase{num: big.NewInt(694280), isKeith: true},
		testCase{num: big.NewInt(925993), isKeith: true},
		testCase{num: big.NewInt(1084051), isKeith: true},
		testCase{num: big.NewInt(7913837), isKeith: true},
		testCase{num: big.NewInt(11436171), isKeith: true},
		testCase{num: big.NewInt(33445755), isKeith: true},
		testCase{num: big.NewInt(87114075), isKeith: false},
		testCase{num: big.NewInt(8380909), isKeith: false},
		testCase{num: big.NewInt(3430268), isKeith: false},
		testCase{num: big.NewInt(431480), isKeith: false},
		testCase{num: big.NewInt(2625840), isKeith: false},
		testCase{num: big.NewInt(903633), isKeith: false},
		testCase{num: big.NewInt(198092), isKeith: false},
		testCase{num: big.NewInt(96882), isKeith: false},
		testCase{num: big.NewInt(79138), isKeith: false},
		testCase{num: big.NewInt(2356), isKeith: false},
		testCase{num: big.NewInt(6632), isKeith: false},
		testCase{num: big.NewInt(471), isKeith: false},
		testCase{num: big.NewInt(450), isKeith: false},
		testCase{num: big.NewInt(26), isKeith: false},
		testCase{num: big.NewInt(23), isKeith: false},
		testCase{num: bigFromString("988242310393860390066911414"), isKeith: true},
	}

	for _, test := range testCases {
		isKeith := keith.IsKeith(test.num)
		if isKeith != test.isKeith {
			t.Errorf("case %v, want %v got %v", test.num, test.isKeith, isKeith)
		}
	}
}

func BenchmarkIsKeith(b *testing.B) {
	// 1000 digit number
	n := bigFromString("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		keith.IsKeith(n)
	}
}
