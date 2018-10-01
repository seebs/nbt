// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

//go:generate go run typegen/main.go -in type.tmp -out typegen.go End Byte Short Int Long Float Double ByteArray String List Compound IntArray LongArray

import (
	"fmt"
	"io"
	"strconv"
)

// Type represents the types of tags available.
type Type uint8

// We define tags for the various NBT tag types.
const (
	TypeEnd Type = iota
	TypeByte
	TypeShort
	TypeInt
	TypeLong
	TypeFloat
	TypeDouble
	TypeByteArray
	TypeString
	TypeList
	TypeCompound
	TypeIntArray
	TypeLongArray
	TypeMax
)

// A Tag is one of several concrete types which represent NBT "payloads",
// because it turns out that's the correct conceptual entity to think of
// as a "tag".
type Tag interface {
	Type() Type
	store(w io.Writer) error
}

type End struct{}
type Byte int8
type Short int16
type Int int32
type Long int64
type Float float32
type Double float64
type ByteArray []int8
type String string

// List is a pain to work with; see tag.tmp for the template code, and
// typegen.go for the fancy generated-code implementations with type
// switches.
type List struct {
	Contents Type
	data     interface{}
}
type Compound map[String]Tag
type IntArray []Int
type LongArray []Long

// You never actually have to make an End to put in a List Of End objects,
// so we check the interface thing here for consistency.
var _ Tag = End{}

// String() makes End objects printable.
func (x End) String() string {
	return ""
}

// String() makes Byte objects printable.
func (x Byte) String() string {
	return fmt.Sprintf("%q", byte(x))
}

// String() makes Short objects printable.
func (x Short) String() string {
	return strconv.FormatInt(int64(x), 10)
}

// String() makes Int objects printable.
func (x Int) String() string {
	return strconv.FormatInt(int64(x), 10)
}

// String() makes Long objects printable.
func (x Long) String() string {
	return strconv.FormatInt(int64(x), 10)
}

// String() makes Float objects printable.
func (x Float) String() string {
	return strconv.FormatFloat(float64(x), 'f', -1, 32)
}

// String() makes Double objects printable.
func (x Double) String() string {
	return strconv.FormatFloat(float64(x), 'f', -1, 64)
}

// String() makes ByteArray objects printable.
func (x ByteArray) String() string {
	return fmt.Sprintf("%v [%d elements]", x.Type(), len(x))
}

// String() makes String objects printable.
func (x String) String() string {
	return string(x)
}

// String() makes List objects printable.
func (x List) String() string {
	return fmt.Sprintf("list[%d elements] of %v", x.Length(), x.Contents)
}

// String() makes Compound objects printable.
func (x Compound) String() string {
	return fmt.Sprintf("%v [%d elements]", x.Type(), len(x))
}

// String() makes IntArray objects printable.
func (x IntArray) String() string {
	return fmt.Sprintf("%v [%d elements]", x.Type(), len(x))
}

// String() makes LongArray objects printable.
func (x LongArray) String() string {
	return fmt.Sprintf("%v [%d elements]", x.Type(), len(x))
}

// PrintIndented pretty-prints the given Tag.
func PrintIndented(w io.Writer, t Tag) {
	printIndented(w, t, nil, 0)
}

// printIndented tries to print the given tag,
func printIndented(w io.Writer, p Tag, prefix interface{}, indent int) {
	fmt.Fprintf(w, "%*s", indent*2, "")
	switch v := prefix.(type) {
	case nil:
		// do nothing with a nil prefix
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
		length := x.Length()
		fmt.Fprintf(w, "[%d %v list] {", length, x.Contents)
		if length != 0 {
			fmt.Fprintf(w, "\n")
			x.Iterate(func(i int, t Tag) error { printIndented(w, p, i, indent+1); return nil })
		}
		fmt.Fprintf(w, "%*s}", indent*2, "")
	case Compound:
		fmt.Fprintf(w, "compound [%d elements] {\n", len(x))
		for k, v := range x {
			printIndented(w, v, k, indent+1)
		}
		fmt.Fprintf(w, "%*s}", indent*2, "")
	}
}

func TagLength(t Tag) int {
	switch tag := t.(type) {
	case ByteArray:
		return len(tag)
	case IntArray:
		return len(tag)
	case LongArray:
		return len(tag)
	case Compound:
		return len(tag)
	case List:
		return tag.Length()
	}
	return 0
}

// Element obtains the element t[idx], where idx is a string for a
// Compound element, or an int for Array or List types.
func TagElement(t Tag, idx interface{}) (out Tag, ok bool) {
	if t == nil {
		return nil, false
	}
	switch tag := t.(type) {
	case Compound:
		sidx, ok := idx.(String)
		if !ok {
			// allow plain Go strings
			str, sok := idx.(string)
			if !sok {
				return nil, false
			}
			sidx = String(str)
		}
		pay, ok := tag[sidx]
		return pay, ok
	case List:
		idx, ok := idx.(int)
		if !ok {
			return nil, false
		}
		data, ok := tag.Element(idx)
		return data, ok
	case ByteArray:
		idx, ok := idx.(int)
		if !ok {
			return nil, false
		}
		if idx >= 0 && idx < len(tag) {
			return Byte(tag[idx]), true
		}
		return nil, false
	case IntArray:
		idx, ok := idx.(int)
		if !ok {
			return nil, false
		}
		if idx >= 0 && idx < len(tag) {
			return tag[idx], true
		}
		return nil, false
	case LongArray:
		idx, ok := idx.(int)
		if !ok {
			return nil, false
		}
		if idx >= 0 && idx < len(tag) {
			return tag[idx], true
		}
		return nil, false
	default:
		return nil, false
	}
}

// HasElements indicates whether an item conceptually has sub-elements.
func TagHasElements(t Tag) bool {
	switch t.Type() {
	case TypeCompound, TypeList, TypeByteArray, TypeIntArray, TypeLongArray:
		return true
	default:
		return false
	}
}
