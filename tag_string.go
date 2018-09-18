package nbt

import "strconv"

const __name = "EndByteShortIntLongFloatDoubleByteArrayStringListCompoundIntArrayLongArray"

var _Tag_index = [...]uint8{0, 3, 7, 12, 15, 19, 24, 30, 39, 45, 49, 57, 65, 74}

func (i Tag) String() string {
	if i >= Tag(len(_Tag_index)-1) {
		return "Tag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Tag_name[_Tag_index[i]:_Tag_index[i+1]]
}
