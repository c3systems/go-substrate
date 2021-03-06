// Code generated by "stringer -type=prefixEnum"; DO NOT EDIT.

package address

import "strconv"

const (
	_prefixEnum_name_0 = "ZeroOne"
	_prefixEnum_name_1 = "Three"
	_prefixEnum_name_2 = "FortyTwoFortyThree"
	_prefixEnum_name_3 = "SixtyEightSixtyNine"
)

var (
	_prefixEnum_index_0 = [...]uint8{0, 4, 7}
	_prefixEnum_index_2 = [...]uint8{0, 8, 18}
	_prefixEnum_index_3 = [...]uint8{0, 10, 19}
)

func (i prefixEnum) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _prefixEnum_name_0[_prefixEnum_index_0[i]:_prefixEnum_index_0[i+1]]
	case i == 3:
		return _prefixEnum_name_1
	case 42 <= i && i <= 43:
		i -= 42
		return _prefixEnum_name_2[_prefixEnum_index_2[i]:_prefixEnum_index_2[i+1]]
	case 68 <= i && i <= 69:
		i -= 68
		return _prefixEnum_name_3[_prefixEnum_index_3[i]:_prefixEnum_index_3[i+1]]
	default:
		return "prefixEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
