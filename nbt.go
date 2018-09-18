// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

import (
	"compress/gzip"
	"fmt"
	"io"
	"math"
)

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

type NBT struct {
	Name    string
	Type    Tag
	payload Payload
}

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
	return n.payload.Store(w)
}

func (n *NBT) Load(r io.Reader) error {
	var err error
	if n.Type == TagEnd {
		return nil
	}
	// every other tag has a name:
	var name String
	p, err := name.Load(r)
	if err != nil {
		return err
	}
	n.Name = string(p.(String))
	switch n.Type {
	case TagByte:
		var b Byte
		n.payload, err = b.Load(r)
		return err
	default:
		return fmt.Errorf("unsupported tag type %s", n.Type)
	}
	return nil
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Tag
	Store(w io.Writer) error
	Load(r io.Reader) (Payload, error)
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

// In NBT, a string is represented as a short bytecount followed by a UTF-8 string.
// We just store the string, and handle the bytecount on save/load.
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

func (p Byte) Store(w io.Writer) error {
	b := [1]byte{byte(p)}
	_, err := w.Write(b[0:1])
	return err
}

func (p Byte) Load(r io.Reader) (Payload, error) {
	var b [1]byte
	_, err := r.Read(b[0:1])
	if err != nil {
		return nil, err
	}
	p = Byte(b[0])
	return p, nil
}

func (p Short) Store(w io.Writer) error {
	var b [2]byte
	b[0] = byte((p >> 8) & 0xFF)
	b[1] = byte(p & 0xFF)
	_, err := w.Write(b[0:2])
	return err
}

func (p Short) Load(r io.Reader) (Payload, error) {
	var b [2]byte
	_, err := r.Read(b[0:2])
	if err != nil {
		return nil, err
	}
	p = Short(b[0]<<8 | b[1])
	return p, nil
}

func (p Int) Store(w io.Writer) error {
	var b [4]byte
	b[0] = byte((p >> 24) & 0xFF)
	b[1] = byte((p >> 16) & 0xFF)
	b[2] = byte((p >> 8) & 0xFF)
	b[3] = byte(p & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Long) Store(w io.Writer) error {
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

func (p Float) Store(w io.Writer) error {
	var b [4]byte
	f := math.Float32bits(float32(p))
	b[0] = byte((f >> 24) & 0xFF)
	b[1] = byte((f >> 16) & 0xFF)
	b[2] = byte((f >> 8) & 0xFF)
	b[3] = byte(f & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Double) Store(w io.Writer) error {
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

/*
func (p ByteArray) Store(w io.Writer) error {
	l := Int(len(p))
	err := l.Store(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(p))
	return err
}
*/

/*
func (p String) Store(w io.Writer) error {
}
func (p List) Store(w io.Writer) error {
}

*/

func (p String) Store(w io.Writer) error {
	if len(p) > 32767 {
		return fmt.Errorf("can't store %d-byte string", len(p))
	}
	sh := Short(len(p))
	err := sh.Store(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(p))
	return err
}

func (p String) Load(r io.Reader) (Payload, error) {
	var sh Short
	l, err := sh.Load(r)
	if err != nil {
		return nil, err
	}
	sh = l.(Short)
	b := make([]byte, int(sh))
	n, err := r.Read(b)
	if err != nil {
		return nil, err
	}
	// if you didn't read enough bytes, okay, fine, we'll just accept that
	return String(b[0:n]), nil
}

func (p Compound) Store(w io.Writer) error {
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
func (p IntArray) Store(w io.Writer) error {
}
func (p LongArray) Store(w io.Writer) error {
}
*/

// Load reads the NBT tag(s) found in the gzipped stream r.
func Load(r io.Reader) (NBT, error) {
	uncomp, err := gzip.NewReader(r)
	if err != nil {
		return NBT{}, err
	}
	defer uncomp.Close()
	return LoadUncompressed(uncomp)
}

// LoadUncompressed reads the NBT tag(s) found in the uncompressed
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
	err := StoreUncompressed(comp, n)
	comp.Close()
	return err
}

func StoreUncompressed(w io.Writer, n NBT) error {
	return n.Store(w)
}
