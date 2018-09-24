package nbt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Path represents a path from the root of a tag (typically the file's top-level
// Compound) to a specific node or value.
type Path struct {
	Tags       []Tag
	Components []PathComponent
}

// PathComponent represents a component of a path, which will be either an
// integer (for List and Array tags) or a string (for Compound tags).
type PathComponent interface {
	Follow(t Tag) (Tag, error)
}

var (
	PathAboveRoot    = errors.New("cannot move above root")
	PathNilComponent = errors.New("cannot have nil component")
)

func PathWrongType(t Tag, idx interface{}) error {
	return fmt.Errorf("path component type mismatch: %v can't be indexed by %T.", t.Type, idx)
}

func PathNoEntry(s String) error {
	return fmt.Errorf("no such entry %q", s)
}

func PathNoIndex(i Int) error {
	return fmt.Errorf("no such entry %d", i)
}

// Follow tries to follow a given name to a new tag.
func (s String) Follow(t Tag) (Tag, error) {
	if t.Type != TypeCompound {
		// try to handle integer indexes that were passed in as strings
		i, err := strconv.ParseInt(string(s), 10, 32)
		if err == nil {
			return Int(i).Follow(t)
		}
		return Tag{}, PathWrongType(t, s)
	}
	out, ok := t.Element(s)
	if !ok {
		return Tag{}, PathNoEntry(s)
	}
	return out, nil
}

// Follow tries to follow a given index to a new tag.
func (i Int) Follow(t Tag) (Tag, error) {
	switch t.Type {
	case TypeList, TypeByteArray, TypeIntArray, TypeLongArray:
		out, ok := t.Element(i)
		if !ok {
			return Tag{}, PathNoIndex(i)
		}
		return out, nil
	default:
		return Tag{}, PathWrongType(t, i)
	}
}

// Cd tries to "change directory" to a new PathComponent.
func (p *Path) Cd(comp PathComponent) (Tag, error) {
	cur := p.Tags[len(p.Components)]
	if comp == nil {
		return cur, PathNilComponent
	}
	if str, ok := comp.(String); ok && str == ".." {
		if len(p.Components) < 1 {
			return p.Tags[0], PathAboveRoot
		}
		newLen := len(p.Components) - 1
		// Tags has one more entry than Components; a zero-entry
		// path still has the root, Tags[0].
		p.Tags = p.Tags[:newLen+1]
		p.Components = p.Components[:newLen]
		return p.Tags[newLen], nil
	}
	newTag, err := comp.Follow(cur)
	if err != nil {
		return newTag, err
	}
	fmt.Printf("cd success: %v\n", newTag)
	// append the requested items to this path
	p.Tags = append(p.Tags, newTag)
	p.Components = append(p.Components, comp)
	return newTag, nil
}

// Current yields the current tag this path points to.
func (p Path) Current() Tag {
	return p.Tags[len(p.Tags)-1]
}

func (p Path) String() string {
	if len(p.Components) == 0 {
		return "/"
	}
	buf := &strings.Builder{}
	for _, c := range p.Components {
		switch c := c.(type) {
		case Int:
			fmt.Fprintf(buf, "%d/", c)
		case String:
			fmt.Fprintf(buf, "%s/", c)
		}
	}
	return buf.String()
}

// NewPath creates a new Path rooted in the given tag.
func NewPath(t Tag) Path {
	return Path{Tags: []Tag{t}}
}
