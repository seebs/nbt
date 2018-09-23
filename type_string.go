package nbt

import "strconv"

const _Type_name = "EndByteShortIntLongFloatDoubleByteArrayStringListCompoundIntArrayLongArray"

var _Type_index = [...]uint8{0, 3, 7, 12, 15, 19, 24, 30, 39, 45, 49, 57, 65, 74}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
