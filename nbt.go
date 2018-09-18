// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

import (
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"unsafe"
)

// Tag represents the types of tags available.
type Tag uint8

const (
	// TagEnd and friends are the NBT tag forms.
	TagEnd Tag = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

// NBT represents a single named tag. There is an internal representation
// of contents; use the Get*() methods to obtain contents.
type NBT struct {
	Name    string
	Type    Tag
	payload Payload
}

// GetEnd returns the (useless) End value, plus a boolean indicating
// whether the NBT in question was in fact a TagEnd.
func (n NBT) GetEnd() (e End, ok bool) {
	if n.Type != TagEnd {
		return End{}, false
	}
	return End{}, true
}

// GetByte returns the Byte value, plus a boolean indicating whether
// the NBT in question was in fact a TagByte.
func (n NBT) GetByte() (b Byte, ok bool) {
	if n.Type != TagByte {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return b, true
	}
	b = n.payload.(Byte)
	return b, true
}

// GetString returns the String value, plus a boolean indicating
// whether the NBT in question was in fact a TagString.
func (n NBT) GetString() (s String, ok bool) {
	if n.Type != TagString {
		return s, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return s, true
	}
	s = n.payload.(String)
	return s, true
}

// Store stores `n` to the provided io.Writer. It does
// not handle compression; for that, use the non-method Store.
func (n NBT) Store(w io.Writer) error {
	// TagEnd doesn't get its name written.
	if n.Type == TagEnd {
		_, err := w.Write([]byte{0})
		return err
	}
	l := len(n.Name)
	var b [3]byte
	b[0] = byte(n.Type)
	b[1] = byte((l >> 8) & 0xFF)
	b[2] = byte(l & 0xFF)
	_, err := w.Write(b[0:3])
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(n.Name))
	if err != nil {
		return err
	}
	return n.payload.store(w)
}

// Load tries to load the name and contents of a tag. The tag's
// type should already be set. For TagEnd, no name is loaded.
func (n *NBT) Load(r io.Reader) error {
	var err error
	if n.Type == TagEnd {
		return nil
	}
	// every tag other than TagEnd has a name:
	name, err := LoadString(r)
	if err != nil {
		return err
	}
	n.Name = string(name)
	switch n.Type {
	case TagByte:
		n.payload, err = LoadByte(r)
	case TagShort:
		n.payload, err = LoadShort(r)
	case TagInt:
		n.payload, err = LoadInt(r)
	case TagLong:
		n.payload, err = LoadLong(r)
	case TagFloat:
		n.payload, err = LoadFloat(r)
	case TagDouble:
		n.payload, err = LoadDouble(r)
	case TagByteArray:
		n.payload, err = LoadByteArray(r)
	case TagString:
		n.payload, err = LoadString(r)
	case TagList:
	case TagCompound:
		n.payload, err = LoadCompound(r)
	case TagIntArray:
	case TagLongArray:
	default:
		err = fmt.Errorf("unsupported tag type %v", n.Type)
	}
	return err
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Tag
	store(w io.Writer) error
}

// End doesn't even have a name, let alone contents.
type End struct{}
type Byte int8
type Short int16
type Int int32
type Long int64
type Float float32
type Double float64
type ByteArray []int8
type String string
type listInternal interface {
	Len() int
	ElemType() Tag
	Elem(i int) NBT
}
type List struct {
	internal listInternal
}
type Compound map[string]NBT
type IntArray []int32
type LongArray []int64

var _ Payload = End{}

// This is an argument for or against generics support.
func (p End) Type() Tag       { return TagEnd }
func (p Byte) Type() Tag      { return TagByte }
func (p Short) Type() Tag     { return TagShort }
func (p Int) Type() Tag       { return TagInt }
func (p Long) Type() Tag      { return TagLong }
func (p Float) Type() Tag     { return TagFloat }
func (p Double) Type() Tag    { return TagDouble }
func (p ByteArray) Type() Tag { return TagByteArray }
func (p String) Type() Tag    { return TagString }
func (p List) Type() Tag      { return TagList }
func (p Compound) Type() Tag  { return TagCompound }
func (p IntArray) Type() Tag  { return TagIntArray }
func (p LongArray) Type() Tag { return TagLongArray }

func (p End) store(w io.Writer) error {
	return nil
}

func (p Byte) store(w io.Writer) error {
	b := [1]byte{byte(p)}
	_, err := w.Write(b[0:1])
	return err
}

func (p Short) store(w io.Writer) error {
	var b [2]byte
	b[0] = byte((p >> 8) & 0xFF)
	b[1] = byte(p & 0xFF)
	_, err := w.Write(b[0:2])
	return err
}

func (p Int) store(w io.Writer) error {
	var b [4]byte
	b[0] = byte((p >> 24) & 0xFF)
	b[1] = byte((p >> 16) & 0xFF)
	b[2] = byte((p >> 8) & 0xFF)
	b[3] = byte(p & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Long) store(w io.Writer) error {
	var b [8]byte
	b[0] = byte((p >> 56) & 0xFF)
	b[1] = byte((p >> 48) & 0xFF)
	b[2] = byte((p >> 40) & 0xFF)
	b[3] = byte((p >> 32) & 0xFF)
	b[4] = byte((p >> 24) & 0xFF)
	b[5] = byte((p >> 16) & 0xFF)
	b[6] = byte((p >> 8) & 0xFF)
	b[7] = byte(p & 0xFF)
	_, err := w.Write(b[0:8])
	return err
}

func (p Float) store(w io.Writer) error {
	var b [4]byte
	f := math.Float32bits(float32(p))
	b[0] = byte((f >> 24) & 0xFF)
	b[1] = byte((f >> 16) & 0xFF)
	b[2] = byte((f >> 8) & 0xFF)
	b[3] = byte(f & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Double) store(w io.Writer) error {
	var b [8]byte
	f := math.Float64bits(float64(p))
	b[0] = byte((f >> 56) & 0xFF)
	b[1] = byte((f >> 48) & 0xFF)
	b[2] = byte((f >> 40) & 0xFF)
	b[3] = byte((f >> 32) & 0xFF)
	b[4] = byte((f >> 24) & 0xFF)
	b[5] = byte((f >> 16) & 0xFF)
	b[6] = byte((f >> 8) & 0xFF)
	b[7] = byte(f & 0xFF)
	_, err := w.Write(b[0:8])
	return err
}

func (p ByteArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	_, err = w.Write(*(*[]byte)(unsafe.Pointer(&p)))
	return err
}

/*
func (p String) store(w io.Writer) error {
}
func (p List) store(w io.Writer) error {
}

*/

func (p String) store(w io.Writer) error {
	if len(p) > 32767 {
		return fmt.Errorf("can't store %d-byte string", len(p))
	}
	sh := Short(len(p))
	err := sh.store(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(p))
	return err
}

func (p Compound) store(w io.Writer) error {
	for k, v := range p {
		n := NBT{Name: k, Type: v.Type, payload: v.payload}
		err := n.Store(w)
		if err != nil {
			return err
		}
	}
	end := NBT{Name: "", Type: TagEnd, payload: nil}
	return end.Store(w)
}

/*
func (p IntArray) store(w io.Writer) error {
}
func (p LongArray) store(w io.Writer) error {
}
*/

// LoadByte loads a Byte payload.
func LoadByte(r io.Reader) (b Byte, e error) {
	var buf [1]byte
	n, err := r.Read(buf[0:1])
	if n == 1 {
		return Byte(buf[0]), nil
	}
	return b, err
}

// LoadShort loads a Short payload.
func LoadShort(r io.Reader) (s Short, e error) {
	var buf [2]byte
	n, err := r.Read(buf[0:2])
	if n == 2 {
		return Short(int16(buf[0])<<8 | int16(buf[1])), nil
	}
	return s, err
}

// LoadInt loads an Int payload.
func LoadInt(r io.Reader) (i Int, e error) {
	var buf [4]byte
	n, err := r.Read(buf[0:4])
	if n == 2 {
		return Int(int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])), nil
	}
	return i, err
}

// LoadLong loads a Long payload.
func LoadLong(r io.Reader) (l Long, e error) {
	var buf [8]byte
	n, err := r.Read(buf[0:8])
	if n == 2 {
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
	return l, err
}

// LoadFloat loads a Float payload.
func LoadFloat(r io.Reader) (f Float, e error) {
	var buf [4]byte
	n, err := r.Read(buf[0:4])
	if n == 2 {
		return Float(math.Float32frombits(uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3]))), nil
	}
	return f, err
}

// LoadDouble loads a Double payload.
func LoadDouble(r io.Reader) (d Double, e error) {
	var buf [8]byte
	n, err := r.Read(buf[0:8])
	if n == 2 {
		return Double(math.Float64frombits(uint64(buf[0])<<56 |
			uint64(buf[1])<<48 |
			uint64(buf[2])<<40 |
			uint64(buf[3])<<32 |
			uint64(buf[4])<<24 |
			uint64(buf[5])<<16 |
			uint64(buf[6])<<8 |
			uint64(buf[7])<<0)), nil
	}
	return d, err
}

// LoadByteArray loads a byte array, which has a leading Int indicating
// how many bytes it contains.
func LoadByteArray(r io.Reader) (b ByteArray, e error) {
	l, err := LoadInt(r)
	if err != nil {
		return b, err
	}
	buf := make([]byte, int(l))
	n, err := r.Read(buf)
	if err != nil && n != int(l) {
		return *(*[]int8)(unsafe.Pointer(&buf)), err
	}
	// if you didn't read enough bytes, okay, fine, we'll just accept that
	return *(*[]int8)(unsafe.Pointer(&buf)), nil
}

// LoadString loads a String payload, reading first a Short payload
// for the string's length, then that many bytes of string data.
func LoadString(r io.Reader) (s String, e error) {
	l, err := LoadShort(r)
	if err != nil {
		return s, err
	}
	b := make([]byte, int(l))
	n, err := r.Read(b)
	if err != nil && n != int(l) {
		return s, err
	}
	// if you didn't read enough bytes, okay, fine, we'll just accept that
	return String(b[0:n]), nil
}

// LoadCompound loads a Compound tag, thus, loads other tags until it gets
// a TagEnd.
func LoadCompound(r io.Reader) (c Compound, e error) {
	c = make(map[string]NBT)
	var n NBT
	var err error
	for n, err = LoadUncompressed(r); err == nil && n.Type != TagEnd; n, err = LoadUncompressed(r) {
		_, ok := c[n.Name]
		if ok {
			return c, fmt.Errorf("duplicate name '%s' in compound tag", n.Name)
		}
		c[n.Name] = n
	}
	if n.Type != TagEnd {
		return c, fmt.Errorf("unterminated compound tag")
	}
	return c, nil
}

// Load reads the first NBT found in the gzipped stream r.
func Load(r io.Reader) (NBT, error) {
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
	n, err := r.Read(tagByte[0:1])
	if err != nil {
		return NBT{}, err
	}
	if n != 1 {
		panic("no byte read on a non-error read, this shouldn't happen")
	}
	tag := Tag(tagByte[0])
	nbt := NBT{Type: tag}
	return nbt, nbt.Load(r)
}

func Store(w io.Writer, n NBT) error {
	comp := gzip.NewWriter(w)
	defer comp.Close()
	return StoreUncompressed(comp, n)
}

func StoreUncompressed(w io.Writer, n NBT) error {
	return n.Store(w)
}
