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

// Tag represents a single named tag. There is an internal representation
// of contents; use the Get*() methods to obtain contents.
type Tag struct {
	Name    String
	Type    Type
	payload Payload
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Type
	store(w io.Writer) error
}

// Named takes a payload (such as a Compound, or String) and
// wraps it into a Tag object with the given name.
func Named(name String, payload Payload) Tag {
	return Tag{Type: payload.Type(), Name: name, payload: payload}
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
type Compound map[String]Payload
type IntArray []Int
type LongArray []Long

// You never actually have to make an End to put in a List Of End objects,
// so we check the interface thing here for consistency.
var _ Payload = End{}

func (n Tag) String() string {
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

// PrintIndented pretty-prints the given Tag.
func (n Tag) PrintIndented(w io.Writer) {
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
			printIndented(w, v, k, indent+1)
		}
		fmt.Fprintf(w, "%*s}", indent*2, "")
	}
}

func (t Tag) Length() int {
	switch t.Type {
	case TypeByteArray:
		return len(t.payload.(ByteArray))
	case TypeIntArray:
		return len(t.payload.(IntArray))
	case TypeLongArray:
		return len(t.payload.(LongArray))
	case TypeCompound:
		return len(t.payload.(Compound))
	case TypeList:
		x, ok := t.payload.(List)
		if !ok {
			fmt.Printf("TypeList with nil payload [%s]", t.Name)
			return 0
		}
		return x.Length()
	}
	return 0
}

// Element obtains the element t[idx], where idx is a string for a
// Compound element, or an int for Array or List types.
func (t Tag) Element(idx interface{}) (out Tag, ok bool) {
	switch t.Type {
	case TypeCompound:
		sidx, ok := idx.(String)
		if !ok {
			// allow plain Go strings
			str, sok := idx.(string)
			if !sok {
				return Tag{}, false
			}
			sidx = String(str)
		}
		pay, ok := t.payload.(Compound)[sidx]
		if ok {
			return Tag{Type: pay.Type(), Name: sidx, payload: pay}, ok
		} else {
			return Tag{}, false
		}
	case TypeList:
		l := t.payload.(List)
		idx, ok := idx.(int)
		if !ok {
			return Tag{}, false
		}
		data, ok := l.Element(idx)
		return Tag{Type: l.Contents, payload: data}, ok
	case TypeByteArray:
		idx, ok := idx.(int)
		if !ok {
			return Tag{}, false
		}
		a := t.payload.(ByteArray)
		if idx >= 0 && idx < len(a) {
			return Tag{Type: TypeByte, payload: Byte(a[idx])}, true
		}
		return Tag{}, false
	case TypeIntArray:
		idx, ok := idx.(int)
		if !ok {
			return Tag{}, false
		}
		a := t.payload.(IntArray)
		if idx >= 0 && idx < len(a) {
			return Tag{Type: TypeInt, payload: a[idx]}, true
		}
		return Tag{}, false
	case TypeLongArray:
		idx, ok := idx.(int)
		if !ok {
			return Tag{}, false
		}
		a := t.payload.(LongArray)
		if idx >= 0 && idx < len(a) {
			return Tag{Type: TypeLong, payload: a[idx]}, true
		}
		return Tag{}, false
	default:
		return Tag{}, false
	}
}

// HasElements indicates whether an item conceptually has sub-elements.
func (t Tag) HasElements() bool {
	switch t.Type {
	case TypeCompound, TypeList, TypeByteArray, TypeIntArray, TypeLongArray:
		return true
	default:
		return false
	}
}
