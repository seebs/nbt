// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

import (
	"fmt"
	"io"
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

// You never actually have to make an End to put in a List Of End objects,
// so we check the interface thing here for consistency.
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
		fmt.Fprintf(w, "%*s}", indent*2, "")
	case Compound:
		fmt.Fprintf(w, "compound [%d elements] {\n", len(x))
		for k, v := range x {
			printIndented(w, v.payload, k, indent+1)
		}
		fmt.Fprintf(w, "%*s}", indent*2, "")
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
