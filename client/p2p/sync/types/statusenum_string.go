// Code generated by "stringer -type=statusEnum"; DO NOT EDIT.

package synctypes

import "strconv"

const _statusEnum_name = "IdleSync"

var _statusEnum_index = [...]uint8{0, 4, 8}

func (i statusEnum) String() string {
	if i < 0 || i >= statusEnum(len(_statusEnum_index)-1) {
		return "statusEnum(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _statusEnum_name[_statusEnum_index[i]:_statusEnum_index[i+1]]
}
