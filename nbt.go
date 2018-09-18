// Package nbt provides an implementation of Minecraft's
// NBT data format.
package nbt

import (
	"compress/gzip"
)

type Tag uint8

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
)

type NBT struct {
	Name string
	Type Tag
	Payload
}

// A Payload represents the payload associated with a named tag.
type Payload interface {
	Type() Tag
	Store(w io.Writer)
	Load(r io.Reader)
}

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
type List interface {
	Len() int
	ElemType() Tag
	Elem(int i) NBT
}
type Compound map[string]NBT
type IntArray []int32
type LongArray []int64

// This is an argument for or against generics support.

func (p *End) Type()       { return TagEnd }
func (p *Byte) Type()      { return TagByte }
func (p *Short) Type()     { return TagShort }
func (p *Int) Type()       { return TagInt }
func (p *Long) Type()      { return TagLong }
func (p *Float) Type()     { return TagFloat }
func (p *Double) Type()    { return TagDouble }
func (p *ByteArray) Type() { return TagByteArray }
func (p *String) Type()    { return TagString }
func (p *List) Type()      { return TagList }
func (p *Compound) Type()  { return TagCompound }
func (p *IntArray) Type()  { return TagIntArray }
func (p *LongArray) Type() { return TagLongArray }
