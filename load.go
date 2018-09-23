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

// LoadByte loads a Byte payload.
func LoadByte(r io.Reader) (b Byte, e error) {
	var buf [1]byte
	_, err := io.ReadFull(r, buf[0:1])
	if err != nil {
		return b, err
	}
	return Byte(buf[0]), nil
}

// LoadShort loads a Short payload.
func LoadShort(r io.Reader) (s Short, e error) {
	var buf [2]byte
	_, err := io.ReadFull(r, buf[0:2])
	if err != nil {
		return s, err
	}
	return Short(int16(buf[0])<<8 | int16(buf[1])), nil
}

// LoadInt loads an Int payload.
func LoadInt(r io.Reader) (i Int, e error) {
	var buf [4]byte
	_, err := io.ReadFull(r, buf[0:4])
	if err != nil {
		return i, err
	}
	return Int(int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])), nil
}

// LoadLong loads a Long payload.
func LoadLong(r io.Reader) (l Long, e error) {
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

// LoadFloat loads a Float payload.
func LoadFloat(r io.Reader) (f Float, e error) {
	var buf [4]byte
	_, err := io.ReadFull(r, buf[0:4])
	if err != nil {
		return f, err
	}
	return Float(math.Float32frombits(uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3]))), nil
}

// LoadDouble loads a Double payload.
func LoadDouble(r io.Reader) (d Double, e error) {
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

// LoadByteArray loads a byte array, which has a leading Int indicating
// how many bytes it contains.
func LoadByteArray(r io.Reader) (b ByteArray, e error) {
	l, err := LoadInt(r)
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

// LoadIntArray loads an Int array, which has a leading Int indicating
// how many Ints it contains.
func LoadIntArray(r io.Reader) (ia IntArray, e error) {
	l, err := LoadInt(r)
	if err != nil {
		return ia, err
	}
	buf := make([]Int, int(l))
	for i := 0; i < int(l); i++ {
		buf[i], e = LoadInt(r)
		if e != nil {
			return ia, e
		}
	}
	return buf, nil
}

// LoadLongArray loads an Int array, which has a leading Int indicating
// how many Long it contains.
func LoadLongArray(r io.Reader) (ia LongArray, e error) {
	l, err := LoadInt(r)
	if err != nil {
		return ia, err
	}
	buf := make([]Long, int(l))
	for i := 0; i < int(l); i++ {
		buf[i], e = LoadLong(r)
		if e != nil {
			return ia, e
		}
	}
	return buf, nil
}

// LoadString loads a String payload, reading first a Short payload
// for the string's length, then that many bytes of string data.
func LoadString(r io.Reader) (s String, e error) {
	sl, err := LoadShort(r)
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

// LoadList loads a List tag.
func LoadList(r io.Reader) (l List, e error) {
	ttype, e := LoadByte(r)
	if e != nil {
		return l, e
	}
	if Type(ttype) < TypeEnd || Type(ttype) >= TypeMax {
		return l, fmt.Errorf("invalid tag type for list: %d", ttype)
	}
	count, e := LoadInt(r)
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

// LoadCompound loads a Compound tag, thus, loads other tags until it gets
// a TypeEnd.
func LoadCompound(r io.Reader) (c Compound, e error) {
	c = make(map[string]NBT)
	var n NBT
	var err error
	var errored error // an error we handle after the fact
	for n, err = LoadUncompressed(r); err == nil && n.Type != TypeEnd; n, err = LoadUncompressed(r) {
		// fmt.Printf("loaded tag: [%v] %s\n", n.Type, n.Name)
		_, ok := c[n.Name]
		if ok {
			// note the thing, but continue using the newer one
			errored = fmt.Errorf("duplicate name '%s' in compound tag", n.Name)
		}
		c[n.Name] = n
	}
	if n.Type != TypeEnd {
		fmt.Printf("failed load compound\n")
		return c, fmt.Errorf("unterminated compound tag")
	}
	return c, errored
}

// LoadCompressed reads the first NBT found in the gzipped stream r.
func LoadCompressed(r io.Reader) (NBT, error) {
	uncomp, err := gzip.NewReader(r)
	if err != nil {
		return NBT{}, err
	}
	defer uncomp.Close()
	return LoadUncompressed(uncomp)
}

// LoadUncompressed reads the first NBT tag found in the uncompressed
// stream r.
func LoadUncompressed(r io.Reader) (NBT, error) {
	var tagByte [1]byte
	_, err := io.ReadFull(r, tagByte[0:1])
	if err != nil {
		return NBT{}, err
	}
	tag := Type(tagByte[0])
	n := NBT{Type: tag}
	if n.Type == TypeEnd {
		return n, nil
	}
	// every tag other than TypeEnd has a name:
	name, err := LoadString(r)
	if err != nil {
		return n, err
	}
	n.Name = string(name)
	// fmt.Printf("load: %s [%v]\n", n.Name, n.Type)
	switch n.Type {
	case TypeByte:
		n.payload, err = LoadByte(r)
	case TypeShort:
		n.payload, err = LoadShort(r)
	case TypeInt:
		n.payload, err = LoadInt(r)
	case TypeLong:
		n.payload, err = LoadLong(r)
	case TypeFloat:
		n.payload, err = LoadFloat(r)
	case TypeDouble:
		n.payload, err = LoadDouble(r)
	case TypeByteArray:
		n.payload, err = LoadByteArray(r)
	case TypeString:
		n.payload, err = LoadString(r)
	case TypeList:
		n.payload, err = LoadList(r)
	case TypeCompound:
		n.payload, err = LoadCompound(r)
	case TypeIntArray:
		n.payload, err = LoadIntArray(r)
	case TypeLongArray:
		n.payload, err = LoadLongArray(r)
	default:
		err = fmt.Errorf("unsupported tag type %v", n.Type)
	}
	if err != nil {
		fmt.Printf("failed to load %s: %s\n", n.Name, err)
	}
	return n, err
}

// Load attempts to determine whether the stream r is compressed or not,
// and use LoadCompressed/LoadUncompressed accordingly.
func Load(r io.Reader) (NBT, error) {
	buf := bufio.NewReader(r)
	header, err := buf.Peek(512)
	// couldn't read the thing
	if err != nil && err != io.EOF {
		return NBT{}, err
	}
	readBuf := bytes.NewBuffer(header)
	gz, err := gzip.NewReader(readBuf)
	if err == nil {
		gz.Close()
		return LoadCompressed(buf)
	}
	return LoadUncompressed(buf)
}
