// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

//go:generate go run taggen/main.go -in tag.tmp -out tagdata.go End Byte Short Int Long Float Double ByteArray String List Compound IntArray LongArray

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

// List is a pain to work with; see tag.tmp for the template code, and
// tagdata.go for the fancy implementations with type switches.
type List struct {
	typ  Tag
	data interface{}
}
type Compound map[string]NBT
type IntArray []Int
type LongArray []Long

// You never actually have to make an End to put in a List Of End objects,
// so we check the interface thing here for consistency.
var _ Payload = End{}

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
		return fmt.Sprintf("list[%d elements] of %v", x.Length(), x.typ)
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
		length := x.Length()
		fmt.Fprintf(w, "[%d %v list] {", length, x.typ)
		if length != 0 {
			fmt.Fprintf(w, "\n")
			x.Iterate(func(i int, p Payload) error { printIndented(w, p, i, indent+1); return nil })
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
		return x.Length()
	}
	return 0
}
