// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

//go:generate go run typegen/main.go -in type.tmp -out typegen.go End Byte Short Int Long Float Double ByteArray String List Compound IntArray LongArray

import (
	"fmt"
	"io"
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

// NBT represents a single named tag. There is an internal representation
// of contents; use the Get*() methods to obtain contents.
type NBT struct {
	Name    string
	Type    Type
	payload Payload
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Type
	store(w io.Writer) error
}

// Named takes a payload (such as a Compound, or String) and
// wraps it into an NBT object with the given name.
func Named(name string, payload Payload) NBT {
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
// typegen.go for the fancy generated-code implementations with type
// switches.
type List struct {
	Contents Type
	data     interface{}
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
	case TypeEnd:
		return ""
	case TypeByte:
		x, _ := n.GetByte()
		return fmt.Sprintf("%q", x)
	case TypeShort:
		x, _ := n.GetShort()
		return fmt.Sprintf("%d", x)
	case TypeInt:
		x, _ := n.GetInt()
		return fmt.Sprintf("%d", x)
	case TypeLong:
		x, _ := n.GetLong()
		return fmt.Sprintf("%d", x)
	case TypeFloat:
		x, _ := n.GetFloat()
		return fmt.Sprintf("%f", x)
	case TypeDouble:
		x, _ := n.GetDouble()
		return fmt.Sprintf("%f", x)
	case TypeString:
		x, _ := n.GetString()
		return fmt.Sprintf("%s", x)
	case TypeList:
		x, _ := n.GetList()
		return fmt.Sprintf("list[%d elements] of %v", x.Length(), x.Contents)
	case TypeByteArray, TypeIntArray, TypeLongArray, TypeCompound:
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
		fmt.Fprintf(w, "[%d %v list] {", length, x.Contents)
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
	case TypeByteArray:
		return len(n.payload.(ByteArray))
	case TypeIntArray:
		return len(n.payload.(IntArray))
	case TypeLongArray:
		return len(n.payload.(LongArray))
	case TypeCompound:
		return len(n.payload.(Compound))
	case TypeList:
		x, ok := n.payload.(List)
		if !ok {
			fmt.Printf("TypeList with nil payload [%s]", n.Name)
			return 0
		}
		return x.Length()
	}
	return 0
}
