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

// We define tags for the various NBT tag types.
const (
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
	TagMax
)

// NBT represents a single named tag. There is an internal representation
// of contents; use the Get*() methods to obtain contents.
type NBT struct {
	Name    string
	Type    Tag
	payload Payload
}

// GetEnd returns the (useless) End value.
func (n NBT) GetEnd() (e End, ok bool) {
	if n.Type != TagEnd {
		return End{}, false
	}
	return End{}, true
}

// GetByte returns the Byte value.
func (n NBT) GetByte() (o Byte, ok bool) {
	if n.Type != TagByte {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Byte)
	return o, true
}

// GetShort returns the Short value.
func (n NBT) GetShort() (o Short, ok bool) {
	if n.Type != TagShort {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Short)
	return o, true
}

// GetInt returns the Int value.
func (n NBT) GetInt() (o Int, ok bool) {
	if n.Type != TagInt {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Int)
	return o, true
}

// GetLong returns the Long value.
func (n NBT) GetLong() (o Long, ok bool) {
	if n.Type != TagLong {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Long)
	return o, true
}

// GetFloat returns the Float value.
func (n NBT) GetFloat() (o Float, ok bool) {
	if n.Type != TagFloat {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Float)
	return o, true
}

// GetDouble returns the Double value.
func (n NBT) GetDouble() (o Double, ok bool) {
	if n.Type != TagDouble {
		return 0, false
	}
	// handle nil payload with zero value
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Double)
	return o, true
}

// GetString returns the String value.
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

// GetList returns the List value.
func (n NBT) GetList() (o List, ok bool) {
	if n.Type != TagList {
		return o, false
	}
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(List)
	return o, true
}

// GetCompound returns the Compound value.
func (n NBT) GetCompound() (o Compound, ok bool) {
	if n.Type != TagCompound {
		return o, false
	}
	if n.payload == nil {
		return o, true
	}
	o = n.payload.(Compound)
	return o, true
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
	fmt.Printf("load: %s [%v]\n", n.Name, n.Type)
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
		n.payload, err = LoadList(r)
	case TagCompound:
		n.payload, err = LoadCompound(r)
	case TagIntArray:
		n.payload, err = LoadIntArray(r)
	case TagLongArray:
		n.payload, err = LoadLongArray(r)
	default:
		err = fmt.Errorf("unsupported tag type %v", n.Type)
	}
	if err != nil {
		fmt.Printf("failed to load %s: %s\n", n.Name, err)
	} else {
		fmt.Printf("load ok: %s\n", n.Name)
	}
	return err
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Tag
	store(w io.Writer) error
}

// NBTNamed takes a payload (such as a Compound, or String) and
// wraps it into an NBT object with the given name.
func NBTNamed(name string, payload Payload) NBT {
	return NBT{Type: payload.Type(), Name: name, payload: payload}
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
type List struct {
	typ  Tag
	data []Payload
}
type Compound map[string]NBT
type IntArray []Int
type LongArray []Long

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

func (n NBT) String() string {
	switch n.Type {
	default:
		return fmt.Sprintf("[unknown tag %v]", n.Type)
	case TagEnd:
		return ""
	case TagByte:
		x, _ := n.GetByte()
		return fmt.Sprintf("%q", x)
	case TagShort:
		x, _ := n.GetShort()
		return fmt.Sprintf("%d", x)
	case TagInt:
		x, _ := n.GetInt()
		return fmt.Sprintf("%d", x)
	case TagLong:
		x, _ := n.GetLong()
		return fmt.Sprintf("%d", x)
	case TagFloat:
		x, _ := n.GetFloat()
		return fmt.Sprintf("%f", x)
	case TagDouble:
		x, _ := n.GetDouble()
		return fmt.Sprintf("%f", x)
	case TagString:
		x, _ := n.GetString()
		return fmt.Sprintf("%s", x)
	case TagList:
		x, _ := n.GetList()
		return fmt.Sprintf("list[%d elements] of %v", len(x.data), x.typ)
	case TagByteArray, TagIntArray, TagLongArray, TagCompound:
		return fmt.Sprintf("%v [%d elements]", n.Type, n.Length())
	}
}

// PrintIndented pretty-prints the given NBT.
func (n NBT) PrintIndented(w io.Writer) {
	printIndented(w, n.payload, n.Name, 0)
}

func printIndented(w io.Writer, p Payload, prefix interface{}, indent int) {
	fmt.Fprintf(w, "%*s", indent*2, "")
	switch v := prefix.(type) {
	case string:
		fmt.Fprintf(w, "%s: ", v)
	case int:
		fmt.Fprintf(w, "[%d]: ", v)
	default:
		fmt.Fprintf(w, "%s: ", v)
	}
	defer fmt.Fprintln(w) // newline after this
	switch x := p.(type) {
	default:
		fmt.Fprintf(w, "[unknown tag %v]", p.Type())
	case End:
		fmt.Fprintf(w, "}")
	case Byte:
		fmt.Fprintf(w, "%q", x)
	case Short:
		fmt.Fprintf(w, "%d", x)
	case Int:
		fmt.Fprintf(w, "%d", x)
	case Long:
		fmt.Fprintf(w, "%d", x)
	case Float:
		fmt.Fprintf(w, "%f", x)
	case Double:
		fmt.Fprintf(w, "%f", x)
	case String:
		fmt.Fprintf(w, "%s", x)
	case ByteArray:
		fmt.Fprintf(w, "[%d item byte]", len(x))
	case IntArray:
		fmt.Fprintf(w, "[%d item int]", len(x))
	case LongArray:
		fmt.Fprintf(w, "[%d item long]", len(x))
	case List:
		fmt.Fprintf(w, "[%d %v list] {", len(x.data), x.typ)
		if len(x.data) != 0 {
			fmt.Fprintf(w, "\n")
			for i, v := range x.data {
				printIndented(w, v, i, indent+1)
			}
		}
		fmt.Fprintf(w, "%*s}\n", indent*2, "")
	case Compound:
		fmt.Fprintf(w, "compound [%d elements] {\n", len(x))
		for k, v := range x {
			printIndented(w, v.payload, k, indent+1)
		}
		fmt.Fprintf(w, "%*s}\n", indent*2, "")
	}
}

func (n NBT) Length() int {
	switch n.Type {
	case TagByteArray:
		return len(n.payload.(ByteArray))
	case TagIntArray:
		return len(n.payload.(IntArray))
	case TagLongArray:
		return len(n.payload.(LongArray))
	case TagCompound:
		return len(n.payload.(Compound))
	case TagList:
		x, ok := n.payload.(List)
		if !ok {
			fmt.Printf("TagList with nil payload [%s]", n.Name)
			return 0
		}
		return len(x.data)
	}
	return 0
}

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

func (p List) store(w io.Writer) error {
	err := Byte(p.typ).store(w)
	if err != nil {
		return err
	}
	l := Int(len(p.data))
	err = l.store(w)
	if err != nil {
		return err
	}
	for i := 0; i < int(l); i++ {
		if p.data[i].Type() != p.typ {
			return fmt.Errorf("list mismatch, expecting %v, found %v", p.typ, p.data[i].Type())
		}
		// store just the payloads, not the whole objects
		err = p.data[i].store(w)
		if err != nil {
			return err
		}
	}
	return nil
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

func (p IntArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	for _, i := range p {
		err = i.store(w)
		if err != nil {
			return err
		}
	}
	return err
}

func (p LongArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	for _, i := range p {
		err = i.store(w)
		if err != nil {
			return err
		}
	}
	return err
}

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
	if Tag(ttype) < TagEnd || Tag(ttype) >= TagMax {
		return l, fmt.Errorf("invalid tag type for list: %d", ttype)
	}
	count, e := LoadInt(r)
	if e != nil {
		return l, e
	}
	if count < 0 {
		return l, fmt.Errorf("invalid negative count for list: %d", count)
	}
	l.typ = Tag(ttype)
	l.data = make([]Payload, count)
	fmt.Printf("list: %v[%d]\n", l.typ, count)
	for i := 0; i < int(count); i++ {
		fmt.Printf("item %d\n", i)
		l.data[i], e = LoadPayload(l.typ, r)
		if e != nil {
			fmt.Printf("list failed at %d: %s\n", i, e)
			break
		}
	}
	return l, e
}

// LoadCompound loads a Compound tag, thus, loads other tags until it gets
// a TagEnd.
func LoadCompound(r io.Reader) (c Compound, e error) {
	c = make(map[string]NBT)
	var n NBT
	var err error
	var errored error // an error we handle after the fact
	for n, err = LoadUncompressed(r); err == nil && n.Type != TagEnd; n, err = LoadUncompressed(r) {
		fmt.Printf("loaded tag: [%v] %s\n", n.Type, n.Name)
		_, ok := c[n.Name]
		if ok {
			// note the thing, but continue using the newer one
			errored = fmt.Errorf("duplicate name '%s' in compound tag", n.Name)
		}
		c[n.Name] = n
	}
	if n.Type != TagEnd {
		fmt.Printf("failed load compound\n")
		return c, fmt.Errorf("unterminated compound tag")
	}
	return c, errored
}

// LoadPayload wraps the type-specific payload loaders.
func LoadPayload(typ Tag, r io.Reader) (p Payload, e error) {
	switch typ {
	case TagByte:
		return LoadByte(r)
	case TagShort:
		return LoadShort(r)
	case TagInt:
		return LoadInt(r)
	case TagLong:
		return LoadLong(r)
	case TagFloat:
		return LoadFloat(r)
	case TagDouble:
		return LoadDouble(r)
	case TagByteArray:
		return LoadByteArray(r)
	case TagString:
		return LoadString(r)
	case TagList:
		return LoadList(r)
	case TagCompound:
		return LoadCompound(r)
	case TagIntArray:
		return LoadIntArray(r)
	case TagLongArray:
		return LoadLongArray(r)
	}
	return nil, fmt.Errorf("unhandled type %v", typ)
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
	_, err := io.ReadFull(r, tagByte[0:1])
	if err != nil {
		return NBT{}, err
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
