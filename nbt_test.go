package nbt

import (
	"bytes"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	x := String("foo")
	n := NBT{Type: TagString, Name: "bar", payload: x}
	buf := &bytes.Buffer{}
	// store x into buf
	err := Store(buf, n)
	t.Logf("buf: % x", buf)
	if err != nil {
		t.Logf("unexpected store err: %s", err)
		return
	}
	y, err := Load(buf)
	if err != nil {
		t.Logf("unexpected load err: %s", err)
	}
	if y.Type != TagString {
		t.Logf("didn't get a string back: %T", y.payload)
	}
	if y.payload == nil {
		t.Logf("nil payload, can't convert that")
	}
	str := y.payload.(String)
	if str != x {
		t.Logf("'%s' != '%s'", x, y)
	}
}
