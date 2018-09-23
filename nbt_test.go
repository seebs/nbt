package nbt

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestSampleData(t *testing.T) {
	bigtest, err := ioutil.ReadFile("examples/bigtest.nbt")
	if err != nil {
		t.Fatalf("couldn't open bigtest.nbt: %s", err)
	}
	t.Run("load", func(t *testing.T) { DoTestLoad(t, bytes.NewBuffer(bigtest)) })
}

func DoTestLoad(t *testing.T, r io.Reader) {
	tag, err := Load(r)
	if err != nil {
		t.Fatalf("couldn't read sample data: %s", err)
	}
	if tag.Name != "Level" {
		t.Fatalf("top level tag named '%q', not 'Level'.", tag.Name)
	}
	_, ok := tag.GetCompound()
	if !ok {
		t.Fatalf("sample data did not give a compound: %v", tag.Type)
	}
}

func TestRoundTrip(t *testing.T) {
	c := make(Compound)
	c["foo"] = String("bar")
	buf := &bytes.Buffer{}
	n := Named("top", c)
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
	if y.Type != TypeCompound {
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
	str, ok := GetString(foo)
	if !ok {
		t.Logf("'foo' is not a string")
	}
	if str != String("bar") {
		t.Logf("'%s' != '%s'", str, "bar")
	}
}
