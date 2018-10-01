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
	tag, name, err := Load(r)
	if err != nil {
		t.Fatalf("couldn't read sample data: %s", err)
	}
	if name != "Level" {
		t.Fatalf("top level tag named '%q', not 'Level'.", name)
	}
	_, ok := GetCompound(tag)
	if !ok {
		t.Fatalf("sample data did not give a compound: %v", tag.Type())
	}
}

func TestRoundTrip(t *testing.T) {
	var err error
	c := make(Compound)
	c["foo"] = String("bar")
	var x List
	x, err = MakeList([]Int{1, 2})
	c["list"] = x
	t.Logf("x: %v, contents %v", x, x.Contents)
	if err != nil {
		t.Fatalf("error trying to make []Int list: %s", err)
	}
	buf := &bytes.Buffer{}
	// store c into buf
	err = StoreTag(buf, c, "top")
	// fmt.Printf("buf: % x", buf)
	if err != nil {
		t.Logf("unexpected store err: %s", err)
		return
	}
	y, name, err := Load(buf)
	if err != nil {
		t.Logf("unexpected load err: %s", err)
	}
	if name != "top" {
		t.Logf("wrong name, got '%q', wanted 'top'", name)
	}
	c2, ok := GetCompound(y)
	if !ok {
		t.Logf("didn't get a compound back: %v", y.Type())
	}
	foo, ok := c2["foo"]
	if !ok {
		t.Logf("no 'foo' in compound")
	}
	bar, ok := TagElement(y, "foo")
	if !ok {
		t.Logf("no 'foo' in compound using Element")
	}
	s1, ok := GetString(foo)
	if !ok {
		t.Logf("'foo' is not a string")
	}
	s2, ok := GetString(bar)
	if s1 != String("bar") {
		t.Logf("s1: '%s' != '%s'", s1, "bar")
	}
	if s2 != String("bar") {
		t.Logf("s2: '%s' != '%s'", s2, "bar")
	}
	list, ok := GetList(c2["list"])
	if !ok {
		t.Logf("no 'list' in compound")
	}
	ints, ok := list.GetIntList()
	if !ok {
		t.Fatalf("compound['list'] doesn't seem to be an Int list: %v", list.Contents)
	}
	if ints[0] != 1 {
		t.Logf("ints[0]: got %d, expecting 1", ints[0])
	}
}
