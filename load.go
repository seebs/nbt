package nbt

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"unsafe"
)

// Functions related to loading NBT tags from streams.

// loadByte loads a Byte payload.
func loadByte(r io.Reader) (b Byte, e error) {
	var buf [1]byte
	_, err := io.ReadFull(r, buf[0:1])
	if err != nil {
		return b, err
	}
	return Byte(buf[0]), nil
}

// loadShort loads a Short payload.
func loadShort(r io.Reader) (s Short, e error) {
	var buf [2]byte
	_, err := io.ReadFull(r, buf[0:2])
	if err != nil {
		return s, err
	}
	return Short(int16(buf[0])<<8 | int16(buf[1])), nil
}

// loadInt loads an Int payload.
func loadInt(r io.Reader) (i Int, e error) {
	var buf [4]byte
	_, err := io.ReadFull(r, buf[0:4])
	if err != nil {
		return i, err
	}
	return Int(int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])), nil
}

// loadLong loads a Long payload.
func loadLong(r io.Reader) (l Long, e error) {
	var buf [8]byte
	_, err := io.ReadFull(r, buf[0:8])
	if err != nil {
		return l, err
	}
	return Long(
		int64(buf[0])<<56 |
			int64(buf[1])<<48 |
			int64(buf[2])<<40 |
			int64(buf[3])<<32 |
			int64(buf[4])<<24 |
			int64(buf[5])<<16 |
			int64(buf[6])<<8 |
			int64(buf[7])<<0), nil
}

// loadFloat loads a Float payload.
func loadFloat(r io.Reader) (f Float, e error) {
	var buf [4]byte
	_, err := io.ReadFull(r, buf[0:4])
	if err != nil {
		return f, err
	}
	return Float(math.Float32frombits(uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3]))), nil
}

// loadDouble loads a Double payload.
func loadDouble(r io.Reader) (d Double, e error) {
	var buf [8]byte
	_, err := io.ReadFull(r, buf[0:8])
	if err != nil {
		return d, err
	}
	return Double(math.Float64frombits(uint64(buf[0])<<56 |
		uint64(buf[1])<<48 |
		uint64(buf[2])<<40 |
		uint64(buf[3])<<32 |
		uint64(buf[4])<<24 |
		uint64(buf[5])<<16 |
		uint64(buf[6])<<8 |
		uint64(buf[7])<<0)), nil
}

// loadByteArray loads a byte array, which has a leading Int indicating
// how many bytes it contains.
func loadByteArray(r io.Reader) (b ByteArray, e error) {
	l, err := loadInt(r)
	if err != nil {
		return b, err
	}
	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return b, err
	}
	return *(*[]int8)(unsafe.Pointer(&buf)), err
}

// loadIntArray loads an Int array, which has a leading Int indicating
// how many Ints it contains.
func loadIntArray(r io.Reader) (ia IntArray, e error) {
	l, err := loadInt(r)
	if err != nil {
		return ia, err
	}
	buf := make([]Int, int(l))
	for i := 0; i < int(l); i++ {
		buf[i], e = loadInt(r)
		if e != nil {
			return ia, e
		}
	}
	return buf, nil
}

// loadLongArray loads an Int array, which has a leading Int indicating
// how many Long it contains.
func loadLongArray(r io.Reader) (ia LongArray, e error) {
	l, err := loadInt(r)
	if err != nil {
		return ia, err
	}
	buf := make([]Long, int(l))
	for i := 0; i < int(l); i++ {
		buf[i], e = loadLong(r)
		if e != nil {
			return ia, e
		}
	}
	return buf, nil
}

// loadString loads a String payload, reading first a Short payload
// for the string's length, then that many bytes of string data.
func loadString(r io.Reader) (s String, e error) {
	sl, err := loadShort(r)
	if err != nil {
		return s, err
	}
	buf := make([]byte, sl)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return s, err
	}
	return String(buf), nil
}

// loadList loads a List tag.
func loadList(r io.Reader) (l List, e error) {
	ttype, e := loadByte(r)
	if e != nil {
		return l, e
	}
	if Type(ttype) < TypeEnd || Type(ttype) >= TypeMax {
		return l, fmt.Errorf("invalid tag type for list: %d", ttype)
	}
	count, e := loadInt(r)
	if e != nil {
		return l, e
	}
	if count < 0 {
		return l, fmt.Errorf("invalid negative count for list: %d", count)
	}
	l.Contents = Type(ttype)
	e = l.loadData(r, int(count))
	return l, e
}

// loadCompound loads a Compound tag, thus, loads other tags until it gets
// a TypeEnd.
func loadCompound(r io.Reader) (c Compound, e error) {
	c = make(map[String]Tag)
	var t Tag
	var name String
	var err error
	var errored error // an error we handle after the fact
	for t, name, err = LoadUncompressed(r); err == nil && t.Type() != TypeEnd; t, name, err = LoadUncompressed(r) {
		// fmt.Printf("loaded tag: [%v] %s\n", t.Type, t.Name)
		_, ok := c[name]
		if ok {
			// note the thing, but continue using the newer one
			errored = fmt.Errorf("duplicate name '%s' in compound tag", name)
		}
		c[name] = t
	}
	if t.Type() != TypeEnd {
		fmt.Printf("failed load compound\n")
		return c, fmt.Errorf("unterminated compound tag")
	}
	return c, errored
}

// LoadCompressed reads the first Tag found in the gzipped stream r.
func LoadCompressed(r io.Reader) (Tag, String, error) {
	uncomp, err := gzip.NewReader(r)
	if err != nil {
		return nil, "", err
	}
	defer uncomp.Close()
	return LoadUncompressed(uncomp)
}

// LoadUncompressed reads the first Tag found in the uncompressed
// stream r.
func LoadUncompressed(r io.Reader) (Tag, String, error) {
	var tagByte [1]byte
	_, err := io.ReadFull(r, tagByte[0:1])
	if err != nil {
		return nil, "", err
	}
	typ := Type(tagByte[0])
	if typ == TypeEnd {
		return End{}, "", nil
	}
	// every tag other than TypeEnd has a name:
	name, err := loadString(r)
	if err != nil {
		return nil, "", err
	}
	var t Tag
	// fmt.Printf("load: %s [%v]\n", n.Name, n.Type)
	switch typ {
	case TypeByte:
		t, err = loadByte(r)
	case TypeShort:
		t, err = loadShort(r)
	case TypeInt:
		t, err = loadInt(r)
	case TypeLong:
		t, err = loadLong(r)
	case TypeFloat:
		t, err = loadFloat(r)
	case TypeDouble:
		t, err = loadDouble(r)
	case TypeByteArray:
		t, err = loadByteArray(r)
	case TypeString:
		t, err = loadString(r)
	case TypeList:
		t, err = loadList(r)
	case TypeCompound:
		t, err = loadCompound(r)
	case TypeIntArray:
		t, err = loadIntArray(r)
	case TypeLongArray:
		t, err = loadLongArray(r)
	default:
		err = fmt.Errorf("unsupported tag type %v", typ)
	}
	if err != nil {
		fmt.Printf("failed to load %s: %s\n", name, err)
	}
	return t, name, err
}

// Load attempts to determine whether the stream r is compressed or not,
// and use LoadCompressed/LoadUncompressed accordingly.
func Load(r io.Reader) (Tag, String, error) {
	buf := bufio.NewReader(r)
	header, err := buf.Peek(512)
	// couldn't read the thing
	if err != nil && err != io.EOF {
		return nil, "", err
	}
	readBuf := bytes.NewBuffer(header)
	gz, err := gzip.NewReader(readBuf)
	if err == nil {
		gz.Close()
		return LoadCompressed(buf)
	}
	return LoadUncompressed(buf)
}
