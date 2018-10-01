package nbt

// GENERATED CODE: Do not edit. See taggen/main.go and tag.go.

import (
	"fmt"
	"io"
)

// End represents the NBT type TAG_End
// Type() tells you that End represents TypeEnd.
func (End) Type() Type { return TypeEnd }

func GetEnd(t Tag) (out End, ok bool) {
	if t.Type() != TypeEnd {
		return out, false
	}
	return out, true
}

// Byte represents the NBT type TAG_Byte
// Type() tells you that Byte represents TypeByte.
func (Byte) Type() Type { return TypeByte }

// GetByte performs a type-assertion that n is of type TypeByte. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Byte, otherwise you get a zero-valued Byte and ok is
// false.
func GetByte(t Tag) (out Byte, ok bool) {
	if t.Type() != TypeByte {
		return out, false
	}
	out, ok = t.(Byte)
	return out, ok
}

// GetByteList performs a type-assertion that l is a list of Byte,
// and returns the corresponding slice.
func (l List) GetByteList() (out []Byte, ok bool) {
	if l.Contents != TypeByte {
		return out, false
	}
	out, ok = l.data.([]Byte)
	return out, ok
}

// MakeByteList creates a list of the appropriate type of payload.
func MakeByteList(in []Byte) (l List) {
	l.Contents = TypeByte
	l.data = in
	return l
}


// Short represents the NBT type TAG_Short
// Type() tells you that Short represents TypeShort.
func (Short) Type() Type { return TypeShort }

// GetShort performs a type-assertion that n is of type TypeShort. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Short, otherwise you get a zero-valued Short and ok is
// false.
func GetShort(t Tag) (out Short, ok bool) {
	if t.Type() != TypeShort {
		return out, false
	}
	out, ok = t.(Short)
	return out, ok
}

// GetShortList performs a type-assertion that l is a list of Short,
// and returns the corresponding slice.
func (l List) GetShortList() (out []Short, ok bool) {
	if l.Contents != TypeShort {
		return out, false
	}
	out, ok = l.data.([]Short)
	return out, ok
}

// MakeShortList creates a list of the appropriate type of payload.
func MakeShortList(in []Short) (l List) {
	l.Contents = TypeShort
	l.data = in
	return l
}


// Int represents the NBT type TAG_Int
// Type() tells you that Int represents TypeInt.
func (Int) Type() Type { return TypeInt }

// GetInt performs a type-assertion that n is of type TypeInt. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Int, otherwise you get a zero-valued Int and ok is
// false.
func GetInt(t Tag) (out Int, ok bool) {
	if t.Type() != TypeInt {
		return out, false
	}
	out, ok = t.(Int)
	return out, ok
}

// GetIntList performs a type-assertion that l is a list of Int,
// and returns the corresponding slice.
func (l List) GetIntList() (out []Int, ok bool) {
	if l.Contents != TypeInt {
		return out, false
	}
	out, ok = l.data.([]Int)
	return out, ok
}

// MakeIntList creates a list of the appropriate type of payload.
func MakeIntList(in []Int) (l List) {
	l.Contents = TypeInt
	l.data = in
	return l
}


// Long represents the NBT type TAG_Long
// Type() tells you that Long represents TypeLong.
func (Long) Type() Type { return TypeLong }

// GetLong performs a type-assertion that n is of type TypeLong. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Long, otherwise you get a zero-valued Long and ok is
// false.
func GetLong(t Tag) (out Long, ok bool) {
	if t.Type() != TypeLong {
		return out, false
	}
	out, ok = t.(Long)
	return out, ok
}

// GetLongList performs a type-assertion that l is a list of Long,
// and returns the corresponding slice.
func (l List) GetLongList() (out []Long, ok bool) {
	if l.Contents != TypeLong {
		return out, false
	}
	out, ok = l.data.([]Long)
	return out, ok
}

// MakeLongList creates a list of the appropriate type of payload.
func MakeLongList(in []Long) (l List) {
	l.Contents = TypeLong
	l.data = in
	return l
}


// Float represents the NBT type TAG_Float
// Type() tells you that Float represents TypeFloat.
func (Float) Type() Type { return TypeFloat }

// GetFloat performs a type-assertion that n is of type TypeFloat. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Float, otherwise you get a zero-valued Float and ok is
// false.
func GetFloat(t Tag) (out Float, ok bool) {
	if t.Type() != TypeFloat {
		return out, false
	}
	out, ok = t.(Float)
	return out, ok
}

// GetFloatList performs a type-assertion that l is a list of Float,
// and returns the corresponding slice.
func (l List) GetFloatList() (out []Float, ok bool) {
	if l.Contents != TypeFloat {
		return out, false
	}
	out, ok = l.data.([]Float)
	return out, ok
}

// MakeFloatList creates a list of the appropriate type of payload.
func MakeFloatList(in []Float) (l List) {
	l.Contents = TypeFloat
	l.data = in
	return l
}


// Double represents the NBT type TAG_Double
// Type() tells you that Double represents TypeDouble.
func (Double) Type() Type { return TypeDouble }

// GetDouble performs a type-assertion that n is of type TypeDouble. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Double, otherwise you get a zero-valued Double and ok is
// false.
func GetDouble(t Tag) (out Double, ok bool) {
	if t.Type() != TypeDouble {
		return out, false
	}
	out, ok = t.(Double)
	return out, ok
}

// GetDoubleList performs a type-assertion that l is a list of Double,
// and returns the corresponding slice.
func (l List) GetDoubleList() (out []Double, ok bool) {
	if l.Contents != TypeDouble {
		return out, false
	}
	out, ok = l.data.([]Double)
	return out, ok
}

// MakeDoubleList creates a list of the appropriate type of payload.
func MakeDoubleList(in []Double) (l List) {
	l.Contents = TypeDouble
	l.data = in
	return l
}


// ByteArray represents the NBT type TAG_ByteArray
// Type() tells you that ByteArray represents TypeByteArray.
func (ByteArray) Type() Type { return TypeByteArray }

// GetByteArray performs a type-assertion that n is of type TypeByteArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to ByteArray, otherwise you get a zero-valued ByteArray and ok is
// false.
func GetByteArray(t Tag) (out ByteArray, ok bool) {
	if t.Type() != TypeByteArray {
		return out, false
	}
	out, ok = t.(ByteArray)
	return out, ok
}

// GetByteArrayList performs a type-assertion that l is a list of ByteArray,
// and returns the corresponding slice.
func (l List) GetByteArrayList() (out []ByteArray, ok bool) {
	if l.Contents != TypeByteArray {
		return out, false
	}
	out, ok = l.data.([]ByteArray)
	return out, ok
}

// MakeByteArrayList creates a list of the appropriate type of payload.
func MakeByteArrayList(in []ByteArray) (l List) {
	l.Contents = TypeByteArray
	l.data = in
	return l
}


// String represents the NBT type TAG_String
// Type() tells you that String represents TypeString.
func (String) Type() Type { return TypeString }

// GetString performs a type-assertion that n is of type TypeString. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to String, otherwise you get a zero-valued String and ok is
// false.
func GetString(t Tag) (out String, ok bool) {
	if t.Type() != TypeString {
		return out, false
	}
	out, ok = t.(String)
	return out, ok
}

// GetStringList performs a type-assertion that l is a list of String,
// and returns the corresponding slice.
func (l List) GetStringList() (out []String, ok bool) {
	if l.Contents != TypeString {
		return out, false
	}
	out, ok = l.data.([]String)
	return out, ok
}

// MakeStringList creates a list of the appropriate type of payload.
func MakeStringList(in []String) (l List) {
	l.Contents = TypeString
	l.data = in
	return l
}


// List represents the NBT type TAG_List
// Type() tells you that List represents TypeList.
func (List) Type() Type { return TypeList }

// GetList performs a type-assertion that n is of type TypeList. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to List, otherwise you get a zero-valued List and ok is
// false.
func GetList(t Tag) (out List, ok bool) {
	if t.Type() != TypeList {
		return out, false
	}
	out, ok = t.(List)
	return out, ok
}

// GetListList performs a type-assertion that l is a list of List,
// and returns the corresponding slice.
func (l List) GetListList() (out []List, ok bool) {
	if l.Contents != TypeList {
		return out, false
	}
	out, ok = l.data.([]List)
	return out, ok
}

// MakeListList creates a list of the appropriate type of payload.
func MakeListList(in []List) (l List) {
	l.Contents = TypeList
	l.data = in
	return l
}


// Compound represents the NBT type TAG_Compound
// Type() tells you that Compound represents TypeCompound.
func (Compound) Type() Type { return TypeCompound }

// GetCompound performs a type-assertion that n is of type TypeCompound. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Compound, otherwise you get a zero-valued Compound and ok is
// false.
func GetCompound(t Tag) (out Compound, ok bool) {
	if t.Type() != TypeCompound {
		return out, false
	}
	out, ok = t.(Compound)
	return out, ok
}

// GetCompoundList performs a type-assertion that l is a list of Compound,
// and returns the corresponding slice.
func (l List) GetCompoundList() (out []Compound, ok bool) {
	if l.Contents != TypeCompound {
		return out, false
	}
	out, ok = l.data.([]Compound)
	return out, ok
}

// MakeCompoundList creates a list of the appropriate type of payload.
func MakeCompoundList(in []Compound) (l List) {
	l.Contents = TypeCompound
	l.data = in
	return l
}


// IntArray represents the NBT type TAG_IntArray
// Type() tells you that IntArray represents TypeIntArray.
func (IntArray) Type() Type { return TypeIntArray }

// GetIntArray performs a type-assertion that n is of type TypeIntArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to IntArray, otherwise you get a zero-valued IntArray and ok is
// false.
func GetIntArray(t Tag) (out IntArray, ok bool) {
	if t.Type() != TypeIntArray {
		return out, false
	}
	out, ok = t.(IntArray)
	return out, ok
}

// GetIntArrayList performs a type-assertion that l is a list of IntArray,
// and returns the corresponding slice.
func (l List) GetIntArrayList() (out []IntArray, ok bool) {
	if l.Contents != TypeIntArray {
		return out, false
	}
	out, ok = l.data.([]IntArray)
	return out, ok
}

// MakeIntArrayList creates a list of the appropriate type of payload.
func MakeIntArrayList(in []IntArray) (l List) {
	l.Contents = TypeIntArray
	l.data = in
	return l
}


// LongArray represents the NBT type TAG_LongArray
// Type() tells you that LongArray represents TypeLongArray.
func (LongArray) Type() Type { return TypeLongArray }

// GetLongArray performs a type-assertion that n is of type TypeLongArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to LongArray, otherwise you get a zero-valued LongArray and ok is
// false.
func GetLongArray(t Tag) (out LongArray, ok bool) {
	if t.Type() != TypeLongArray {
		return out, false
	}
	out, ok = t.(LongArray)
	return out, ok
}

// GetLongArrayList performs a type-assertion that l is a list of LongArray,
// and returns the corresponding slice.
func (l List) GetLongArrayList() (out []LongArray, ok bool) {
	if l.Contents != TypeLongArray {
		return out, false
	}
	out, ok = l.data.([]LongArray)
	return out, ok
}

// MakeLongArrayList creates a list of the appropriate type of payload.
func MakeLongArrayList(in []LongArray) (l List) {
	l.Contents = TypeLongArray
	l.data = in
	return l
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
		return fmt.Errorf("unhandled tag type in List.storeData: %v", l.Contents)
	}
	return nil
}

// loadData loads the "raw" data array, which we'll later use to build
// the interface array.
func (l *List) loadData(r io.Reader, count int) (err error) {
	switch l.Contents {

	case TypeEnd: // nothing to load
		l.data = nil
		return nil

	case TypeByte:
		raw := make([]Byte, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadByte(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeShort:
		raw := make([]Short, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadShort(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeInt:
		raw := make([]Int, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadInt(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeLong:
		raw := make([]Long, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadLong(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeFloat:
		raw := make([]Float, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadFloat(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeDouble:
		raw := make([]Double, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadDouble(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeByteArray:
		raw := make([]ByteArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadByteArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeString:
		raw := make([]String, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadString(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeList:
		raw := make([]List, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadList(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeCompound:
		raw := make([]Compound, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadCompound(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeIntArray:
		raw := make([]IntArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadIntArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	case TypeLongArray:
		raw := make([]LongArray, count)
		for i := 0; i < count; i++ {
			raw[i], err = loadLongArray(r)
			if err!= nil {
				raw = raw[:i]
				break
			}
		}
		l.data = raw
		return err

	default:
		return fmt.Errorf("unhandled tag type in List.loadData: %v", l.Contents)
	}
}

// Iterate iterates over the list, passing each item in the list (as a Payload)
// to the given function. If fn returns a non-nil error, Iterate stops and returns
// the error.
func (l List) Iterate(fn func(int, Tag) error) (err error) {
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
		return fmt.Errorf("unhandled tag type in List.Iterate: %v", l.Contents)
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

// MakeList makes a list given a slice of any kind of payload object. Note,
// not a slice of Tags, a slice of any of the specific concrete types
// implement tag and aren't End.
func MakeList(in interface{}) (l List, err error) {
	switch in.(type) {

	case []End:
		// We don't allow non-empty lists of Ends
		l.Contents = TypeEnd
		l.data = nil
		return l, err
	case []Byte:
		l.Contents = TypeByte
		l.data = in
		return l, err
	case []Short:
		l.Contents = TypeShort
		l.data = in
		return l, err
	case []Int:
		l.Contents = TypeInt
		l.data = in
		return l, err
	case []Long:
		l.Contents = TypeLong
		l.data = in
		return l, err
	case []Float:
		l.Contents = TypeFloat
		l.data = in
		return l, err
	case []Double:
		l.Contents = TypeDouble
		l.data = in
		return l, err
	case []ByteArray:
		l.Contents = TypeByteArray
		l.data = in
		return l, err
	case []String:
		l.Contents = TypeString
		l.data = in
		return l, err
	case []List:
		l.Contents = TypeList
		l.data = in
		return l, err
	case []Compound:
		l.Contents = TypeCompound
		l.data = in
		return l, err
	case []IntArray:
		l.Contents = TypeIntArray
		l.data = in
		return l, err
	case []LongArray:
		l.Contents = TypeLongArray
		l.data = in
		return l, err
	default:
		return l, fmt.Errorf("can't MakeList on %T", in)
	}
}

// Element gives the ith element of l.
func (l List)Element(i int) (out Tag, ok bool) {
	switch data := l.data.(type) {

	case []End:
		return End{}, false
	case []Byte:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Short:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Int:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Long:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Float:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Double:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []ByteArray:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []String:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []List:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []Compound:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []IntArray:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	case []LongArray:
		if i > 0 && i < len(data) {
			return data[i], true
		}
		return nil, false
	default:
		return nil, false
	}
}
