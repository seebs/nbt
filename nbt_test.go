package nbt

import (
	"bytes"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	c := make(Compound)
	c["foo"] = NBT{payload: String("bar"), Type: TagString, Name: "foo"}
	buf := &bytes.Buffer{}
	n := NBT{Type: TagCompound, Name: "top", payload: c}
	// store x into buf
	err := Store(buf, n)
	// fmt.Printf("buf: % x", buf)
	if err != nil {
		t.Logf("unexpected store err: %s", err)
		return
	}
	y, err := Load(buf)
	if err != nil {
		t.Logf("unexpected load err: %s", err)
	}
	if y.Type != TagCompound {
		t.Logf("didn't get a string back: %v", y.Type)
	}
	if y.payload == nil {
		t.Logf("nil payload, can't convert that")
	}
	c2 := y.payload.(Compound)
	foo, ok := c2["foo"]
	if !ok {
		t.Logf("no 'foo' in compound")
	}
	str, ok := foo.GetString()
	if !ok {
		t.Logf("'foo' is not a string")
	}
	if str != String("bar") {
		t.Logf("'%s' != '%s'", str, "bar")
	}
}
