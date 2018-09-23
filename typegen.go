package nbt

// GENERATED CODE: Do not edit. See taggen/main.go and tag.go.

import (
	"fmt"
	"io"
)

// End represents the NBT type TAG_End
// Type() tells you that End represents TypeEnd.
func (End) Type() Type { return TypeEnd }

func (t Tag) GetEnd() (out End, ok bool) {
	if t.Type != TypeEnd {
		return out, false
	}
	return out, true
}

func GetEnd(p Payload) (out End, ok bool) {
	if p.Type() != TypeEnd {
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
func (t Tag) GetByte() (out Byte, ok bool) {
	if t.Type != TypeByte {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Byte)
	return out, ok
}

// GetByte performs a type-assertion that n is of type TypeByte. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Byte, otherwise you get a zero-valued Byte and ok is
// false.
func GetByte(p Payload) (out Byte, ok bool) {
	if p.Type() != TypeByte {
		return out, false
	}
	out, ok = p.(Byte)
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


// Short represents the NBT type TAG_Short
// Type() tells you that Short represents TypeShort.
func (Short) Type() Type { return TypeShort }

// GetShort performs a type-assertion that n is of type TypeShort. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Short, otherwise you get a zero-valued Short and ok is
// false.
func (t Tag) GetShort() (out Short, ok bool) {
	if t.Type != TypeShort {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Short)
	return out, ok
}

// GetShort performs a type-assertion that n is of type TypeShort. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Short, otherwise you get a zero-valued Short and ok is
// false.
func GetShort(p Payload) (out Short, ok bool) {
	if p.Type() != TypeShort {
		return out, false
	}
	out, ok = p.(Short)
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


// Int represents the NBT type TAG_Int
// Type() tells you that Int represents TypeInt.
func (Int) Type() Type { return TypeInt }

// GetInt performs a type-assertion that n is of type TypeInt. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Int, otherwise you get a zero-valued Int and ok is
// false.
func (t Tag) GetInt() (out Int, ok bool) {
	if t.Type != TypeInt {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Int)
	return out, ok
}

// GetInt performs a type-assertion that n is of type TypeInt. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Int, otherwise you get a zero-valued Int and ok is
// false.
func GetInt(p Payload) (out Int, ok bool) {
	if p.Type() != TypeInt {
		return out, false
	}
	out, ok = p.(Int)
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


// Long represents the NBT type TAG_Long
// Type() tells you that Long represents TypeLong.
func (Long) Type() Type { return TypeLong }

// GetLong performs a type-assertion that n is of type TypeLong. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Long, otherwise you get a zero-valued Long and ok is
// false.
func (t Tag) GetLong() (out Long, ok bool) {
	if t.Type != TypeLong {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Long)
	return out, ok
}

// GetLong performs a type-assertion that n is of type TypeLong. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Long, otherwise you get a zero-valued Long and ok is
// false.
func GetLong(p Payload) (out Long, ok bool) {
	if p.Type() != TypeLong {
		return out, false
	}
	out, ok = p.(Long)
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


// Float represents the NBT type TAG_Float
// Type() tells you that Float represents TypeFloat.
func (Float) Type() Type { return TypeFloat }

// GetFloat performs a type-assertion that n is of type TypeFloat. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Float, otherwise you get a zero-valued Float and ok is
// false.
func (t Tag) GetFloat() (out Float, ok bool) {
	if t.Type != TypeFloat {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Float)
	return out, ok
}

// GetFloat performs a type-assertion that n is of type TypeFloat. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Float, otherwise you get a zero-valued Float and ok is
// false.
func GetFloat(p Payload) (out Float, ok bool) {
	if p.Type() != TypeFloat {
		return out, false
	}
	out, ok = p.(Float)
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


// Double represents the NBT type TAG_Double
// Type() tells you that Double represents TypeDouble.
func (Double) Type() Type { return TypeDouble }

// GetDouble performs a type-assertion that n is of type TypeDouble. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Double, otherwise you get a zero-valued Double and ok is
// false.
func (t Tag) GetDouble() (out Double, ok bool) {
	if t.Type != TypeDouble {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Double)
	return out, ok
}

// GetDouble performs a type-assertion that n is of type TypeDouble. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Double, otherwise you get a zero-valued Double and ok is
// false.
func GetDouble(p Payload) (out Double, ok bool) {
	if p.Type() != TypeDouble {
		return out, false
	}
	out, ok = p.(Double)
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


// ByteArray represents the NBT type TAG_ByteArray
// Type() tells you that ByteArray represents TypeByteArray.
func (ByteArray) Type() Type { return TypeByteArray }

// GetByteArray performs a type-assertion that n is of type TypeByteArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to ByteArray, otherwise you get a zero-valued ByteArray and ok is
// false.
func (t Tag) GetByteArray() (out ByteArray, ok bool) {
	if t.Type != TypeByteArray {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(ByteArray)
	return out, ok
}

// GetByteArray performs a type-assertion that n is of type TypeByteArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to ByteArray, otherwise you get a zero-valued ByteArray and ok is
// false.
func GetByteArray(p Payload) (out ByteArray, ok bool) {
	if p.Type() != TypeByteArray {
		return out, false
	}
	out, ok = p.(ByteArray)
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


// String represents the NBT type TAG_String
// Type() tells you that String represents TypeString.
func (String) Type() Type { return TypeString }

// GetString performs a type-assertion that n is of type TypeString. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to String, otherwise you get a zero-valued String and ok is
// false.
func (t Tag) GetString() (out String, ok bool) {
	if t.Type != TypeString {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(String)
	return out, ok
}

// GetString performs a type-assertion that n is of type TypeString. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to String, otherwise you get a zero-valued String and ok is
// false.
func GetString(p Payload) (out String, ok bool) {
	if p.Type() != TypeString {
		return out, false
	}
	out, ok = p.(String)
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


// List represents the NBT type TAG_List
// Type() tells you that List represents TypeList.
func (List) Type() Type { return TypeList }

// GetList performs a type-assertion that n is of type TypeList. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to List, otherwise you get a zero-valued List and ok is
// false.
func (t Tag) GetList() (out List, ok bool) {
	if t.Type != TypeList {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(List)
	return out, ok
}

// GetList performs a type-assertion that n is of type TypeList. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to List, otherwise you get a zero-valued List and ok is
// false.
func GetList(p Payload) (out List, ok bool) {
	if p.Type() != TypeList {
		return out, false
	}
	out, ok = p.(List)
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


// Compound represents the NBT type TAG_Compound
// Type() tells you that Compound represents TypeCompound.
func (Compound) Type() Type { return TypeCompound }

// GetCompound performs a type-assertion that n is of type TypeCompound. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Compound, otherwise you get a zero-valued Compound and ok is
// false.
func (t Tag) GetCompound() (out Compound, ok bool) {
	if t.Type != TypeCompound {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(Compound)
	return out, ok
}

// GetCompound performs a type-assertion that n is of type TypeCompound. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to Compound, otherwise you get a zero-valued Compound and ok is
// false.
func GetCompound(p Payload) (out Compound, ok bool) {
	if p.Type() != TypeCompound {
		return out, false
	}
	out, ok = p.(Compound)
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


// IntArray represents the NBT type TAG_IntArray
// Type() tells you that IntArray represents TypeIntArray.
func (IntArray) Type() Type { return TypeIntArray }

// GetIntArray performs a type-assertion that n is of type TypeIntArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to IntArray, otherwise you get a zero-valued IntArray and ok is
// false.
func (t Tag) GetIntArray() (out IntArray, ok bool) {
	if t.Type != TypeIntArray {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(IntArray)
	return out, ok
}

// GetIntArray performs a type-assertion that n is of type TypeIntArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to IntArray, otherwise you get a zero-valued IntArray and ok is
// false.
func GetIntArray(p Payload) (out IntArray, ok bool) {
	if p.Type() != TypeIntArray {
		return out, false
	}
	out, ok = p.(IntArray)
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


// LongArray represents the NBT type TAG_LongArray
// Type() tells you that LongArray represents TypeLongArray.
func (LongArray) Type() Type { return TypeLongArray }

// GetLongArray performs a type-assertion that n is of type TypeLongArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to LongArray, otherwise you get a zero-valued LongArray and ok is
// false.
func (t Tag) GetLongArray() (out LongArray, ok bool) {
	if t.Type != TypeLongArray {
		return out, false
	}
	if t.payload == nil {
		return out, false
	}
	out, ok = t.payload.(LongArray)
	return out, ok
}

// GetLongArray performs a type-assertion that n is of type TypeLongArray. If
// it is, and there is a payload, you get the results of a type-assertion
// of payload to LongArray, otherwise you get a zero-valued LongArray and ok is
// false.
func GetLongArray(p Payload) (out LongArray, ok bool) {
	if p.Type() != TypeLongArray {
		return out, false
	}
	out, ok = p.(LongArray)
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
func (l List) loadData(r io.Reader, count int) (err error) {
	switch l.Contents {

	case TypeEnd: // nothing to load
		l.data = nil
		return nil

	case TypeByte:
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

	case TypeShort:
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

	case TypeInt:
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

	case TypeLong:
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

	case TypeFloat:
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

	case TypeDouble:
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

	case TypeByteArray:
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

	case TypeString:
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

	case TypeList:
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

	case TypeCompound:
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

	case TypeIntArray:
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

	case TypeLongArray:
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
		return fmt.Errorf("unhandled tag type in List.loadData: %v", l.Contents)
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
