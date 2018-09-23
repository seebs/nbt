package nbt

// GENERATED CODE: Do not edit. See taggen/main.go and tag.go.

import (
	"fmt"
	"io"
)

// End represents the NBT type TAG_End
// Type() tells you that End represents TagEnd.
func (End) Type() Tag { return TagEnd }

func (n NBT) GetEnd() (out End, ok bool) {
	if n.Type != TagEnd {
		return out, false
	}
	return out, true
}

// Byte represents the NBT type TAG_Byte
// Type() tells you that Byte represents TagByte.
func (Byte) Type() Tag { return TagByte }

// GetByte performs a type-assertion that n is of type TagByte. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Byte, otherwise you get a zero-valued Byte and ok is
// false.
func (n NBT) GetByte() (out Byte, ok bool) {
	if n.Type != TagByte {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Byte)
	return out, ok
}

// GetByteList performs a type-assertion that l is a list of Byte,
// and returns the corresponding slice.
func (l List) GetByteList() (out []Byte, ok bool) {
	if l.typ != TagByte {
		return out, false
	}
	out, ok = l.data.([]Byte)
	return out, ok
}


// Short represents the NBT type TAG_Short
// Type() tells you that Short represents TagShort.
func (Short) Type() Tag { return TagShort }

// GetShort performs a type-assertion that n is of type TagShort. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Short, otherwise you get a zero-valued Short and ok is
// false.
func (n NBT) GetShort() (out Short, ok bool) {
	if n.Type != TagShort {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Short)
	return out, ok
}

// GetShortList performs a type-assertion that l is a list of Short,
// and returns the corresponding slice.
func (l List) GetShortList() (out []Short, ok bool) {
	if l.typ != TagShort {
		return out, false
	}
	out, ok = l.data.([]Short)
	return out, ok
}


// Int represents the NBT type TAG_Int
// Type() tells you that Int represents TagInt.
func (Int) Type() Tag { return TagInt }

// GetInt performs a type-assertion that n is of type TagInt. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Int, otherwise you get a zero-valued Int and ok is
// false.
func (n NBT) GetInt() (out Int, ok bool) {
	if n.Type != TagInt {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Int)
	return out, ok
}

// GetIntList performs a type-assertion that l is a list of Int,
// and returns the corresponding slice.
func (l List) GetIntList() (out []Int, ok bool) {
	if l.typ != TagInt {
		return out, false
	}
	out, ok = l.data.([]Int)
	return out, ok
}


// Long represents the NBT type TAG_Long
// Type() tells you that Long represents TagLong.
func (Long) Type() Tag { return TagLong }

// GetLong performs a type-assertion that n is of type TagLong. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Long, otherwise you get a zero-valued Long and ok is
// false.
func (n NBT) GetLong() (out Long, ok bool) {
	if n.Type != TagLong {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Long)
	return out, ok
}

// GetLongList performs a type-assertion that l is a list of Long,
// and returns the corresponding slice.
func (l List) GetLongList() (out []Long, ok bool) {
	if l.typ != TagLong {
		return out, false
	}
	out, ok = l.data.([]Long)
	return out, ok
}


// Float represents the NBT type TAG_Float
// Type() tells you that Float represents TagFloat.
func (Float) Type() Tag { return TagFloat }

// GetFloat performs a type-assertion that n is of type TagFloat. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Float, otherwise you get a zero-valued Float and ok is
// false.
func (n NBT) GetFloat() (out Float, ok bool) {
	if n.Type != TagFloat {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Float)
	return out, ok
}

// GetFloatList performs a type-assertion that l is a list of Float,
// and returns the corresponding slice.
func (l List) GetFloatList() (out []Float, ok bool) {
	if l.typ != TagFloat {
		return out, false
	}
	out, ok = l.data.([]Float)
	return out, ok
}


// Double represents the NBT type TAG_Double
// Type() tells you that Double represents TagDouble.
func (Double) Type() Tag { return TagDouble }

// GetDouble performs a type-assertion that n is of type TagDouble. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Double, otherwise you get a zero-valued Double and ok is
// false.
func (n NBT) GetDouble() (out Double, ok bool) {
	if n.Type != TagDouble {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Double)
	return out, ok
}

// GetDoubleList performs a type-assertion that l is a list of Double,
// and returns the corresponding slice.
func (l List) GetDoubleList() (out []Double, ok bool) {
	if l.typ != TagDouble {
		return out, false
	}
	out, ok = l.data.([]Double)
	return out, ok
}


// ByteArray represents the NBT type TAG_ByteArray
// Type() tells you that ByteArray represents TagByteArray.
func (ByteArray) Type() Tag { return TagByteArray }

// GetByteArray performs a type-assertion that n is of type TagByteArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to ByteArray, otherwise you get a zero-valued ByteArray and ok is
// false.
func (n NBT) GetByteArray() (out ByteArray, ok bool) {
	if n.Type != TagByteArray {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(ByteArray)
	return out, ok
}

// GetByteArrayList performs a type-assertion that l is a list of ByteArray,
// and returns the corresponding slice.
func (l List) GetByteArrayList() (out []ByteArray, ok bool) {
	if l.typ != TagByteArray {
		return out, false
	}
	out, ok = l.data.([]ByteArray)
	return out, ok
}


// String represents the NBT type TAG_String
// Type() tells you that String represents TagString.
func (String) Type() Tag { return TagString }

// GetString performs a type-assertion that n is of type TagString. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to String, otherwise you get a zero-valued String and ok is
// false.
func (n NBT) GetString() (out String, ok bool) {
	if n.Type != TagString {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(String)
	return out, ok
}

// GetStringList performs a type-assertion that l is a list of String,
// and returns the corresponding slice.
func (l List) GetStringList() (out []String, ok bool) {
	if l.typ != TagString {
		return out, false
	}
	out, ok = l.data.([]String)
	return out, ok
}


// List represents the NBT type TAG_List
// Type() tells you that List represents TagList.
func (List) Type() Tag { return TagList }

// GetList performs a type-assertion that n is of type TagList. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to List, otherwise you get a zero-valued List and ok is
// false.
func (n NBT) GetList() (out List, ok bool) {
	if n.Type != TagList {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(List)
	return out, ok
}

// GetListList performs a type-assertion that l is a list of List,
// and returns the corresponding slice.
func (l List) GetListList() (out []List, ok bool) {
	if l.typ != TagList {
		return out, false
	}
	out, ok = l.data.([]List)
	return out, ok
}


// Compound represents the NBT type TAG_Compound
// Type() tells you that Compound represents TagCompound.
func (Compound) Type() Tag { return TagCompound }

// GetCompound performs a type-assertion that n is of type TagCompound. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Compound, otherwise you get a zero-valued Compound and ok is
// false.
func (n NBT) GetCompound() (out Compound, ok bool) {
	if n.Type != TagCompound {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(Compound)
	return out, ok
}

// GetCompoundList performs a type-assertion that l is a list of Compound,
// and returns the corresponding slice.
func (l List) GetCompoundList() (out []Compound, ok bool) {
	if l.typ != TagCompound {
		return out, false
	}
	out, ok = l.data.([]Compound)
	return out, ok
}


// IntArray represents the NBT type TAG_IntArray
// Type() tells you that IntArray represents TagIntArray.
func (IntArray) Type() Tag { return TagIntArray }

// GetIntArray performs a type-assertion that n is of type TagIntArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to IntArray, otherwise you get a zero-valued IntArray and ok is
// false.
func (n NBT) GetIntArray() (out IntArray, ok bool) {
	if n.Type != TagIntArray {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(IntArray)
	return out, ok
}

// GetIntArrayList performs a type-assertion that l is a list of IntArray,
// and returns the corresponding slice.
func (l List) GetIntArrayList() (out []IntArray, ok bool) {
	if l.typ != TagIntArray {
		return out, false
	}
	out, ok = l.data.([]IntArray)
	return out, ok
}


// LongArray represents the NBT type TAG_LongArray
// Type() tells you that LongArray represents TagLongArray.
func (LongArray) Type() Tag { return TagLongArray }

// GetLongArray performs a type-assertion that n is of type TagLongArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to LongArray, otherwise you get a zero-valued LongArray and ok is
// false.
func (n NBT) GetLongArray() (out LongArray, ok bool) {
	if n.Type != TagLongArray {
		return out, false
	}
	if n.payload == nil {
		return out, false
	}
	out, ok = n.payload.(LongArray)
	return out, ok
}

// GetLongArrayList performs a type-assertion that l is a list of LongArray,
// and returns the corresponding slice.
func (l List) GetLongArrayList() (out []LongArray, ok bool) {
	if l.typ != TagLongArray {
		return out, false
	}
	out, ok = l.data.([]LongArray)
	return out, ok
}




func (l List) storeData(w io.Writer) (err error) {
	switch raw := l.data.(type) {

	case []End: // no data to store
		return nil

	case []Byte:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Short:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Int:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Long:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Float:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Double:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []ByteArray:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []String:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []List:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []Compound:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []IntArray:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	case []LongArray:
		count := len(raw)
		for i := 0; i < count; i++ {
			err = raw[i].store(w)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unhandled tag type in List.storeData: %v", l.typ)
	}
	return nil
}

// loadData loads the "raw" data array, which we'll later use to build
// the interface array.
func (l List) loadData(r io.Reader, count int) (err error) {
	switch l.typ {

	case TagEnd: // nothing to load
		l.data = nil
		return nil

	case TagByte:
		raw := make([]Byte, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadByte(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagShort:
		raw := make([]Short, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadShort(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagInt:
		raw := make([]Int, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadInt(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagLong:
		raw := make([]Long, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadLong(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagFloat:
		raw := make([]Float, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadFloat(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagDouble:
		raw := make([]Double, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadDouble(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagByteArray:
		raw := make([]ByteArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadByteArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagString:
		raw := make([]String, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadString(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagList:
		raw := make([]List, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadList(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagCompound:
		raw := make([]Compound, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadCompound(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagIntArray:
		raw := make([]IntArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadIntArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TagLongArray:
		raw := make([]LongArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = LoadLongArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	default:
		return fmt.Errorf("unhandled tag type in List.loadData: %v", l.typ)
	}
}

// Iterate iterates over the list, passing each item in the list (as a Payload)
// to the given function. If fn returns a non-nil error, Iterate stops and returns
// the error.
func (l List) Iterate(fn func(int, Payload) error) (err error) {
	switch raw := l.data.(type) {

	case []End:
		count := len(raw)
		for i := 0; i < count; i++ {

			err = fn(i, End{})
			if err != nil {
				break
			}
		}
	case []Byte:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Short:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Int:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Long:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Float:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Double:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []ByteArray:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []String:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []List:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []Compound:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []IntArray:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	case []LongArray:
		count := len(raw)
		for i := 0; i < count; i++ {
err = fn(i, raw[i])

			if err != nil {
				break
			}
		}
	default:
		return fmt.Errorf("unhandled tag type in List.Iterate: %v", l.typ)
	}
	return err
}

// Length returns the length of the list, if applicable. Note, a list of End
// is (I think) always of length 0, if it's even valid at all.
func (l List) Length() int {
	switch raw := l.data.(type) {

	case []End:

		return 0

	case []Byte:

		return len(raw)

	case []Short:

		return len(raw)

	case []Int:

		return len(raw)

	case []Long:

		return len(raw)

	case []Float:

		return len(raw)

	case []Double:

		return len(raw)

	case []ByteArray:

		return len(raw)

	case []String:

		return len(raw)

	case []List:

		return len(raw)

	case []Compound:

		return len(raw)

	case []IntArray:

		return len(raw)

	case []LongArray:

		return len(raw)

	default:
	 	return 0
	}
}
