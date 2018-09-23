package nbt

import (
	"compress/gzip"
	"fmt"
	"io"
	"math"
	"unsafe"
)

// functionality related to storing NBT data to streams

// Store stores `n` to the provided io.Writer. It does
// not handle compression; for that, use the non-method Store.
func (n NBT) Store(w io.Writer) error {
	// TagEnd doesn't get its name written.
	if n.Type == TagEnd {
		_, err := w.Write([]byte{0})
		return err
	}
	l := len(n.Name)
	var b [3]byte
	b[0] = byte(n.Type)
	b[1] = byte((l >> 8) & 0xFF)
	b[2] = byte(l & 0xFF)
	_, err := w.Write(b[0:3])
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(n.Name))
	if err != nil {
		return err
	}
	return n.payload.store(w)
}

func (p End) store(w io.Writer) error {
	return nil
}

func (p Byte) store(w io.Writer) error {
	b := [1]byte{byte(p)}
	_, err := w.Write(b[0:1])
	return err
}

func (p Short) store(w io.Writer) error {
	var b [2]byte
	b[0] = byte((p >> 8) & 0xFF)
	b[1] = byte(p & 0xFF)
	_, err := w.Write(b[0:2])
	return err
}

func (p Int) store(w io.Writer) error {
	var b [4]byte
	b[0] = byte((p >> 24) & 0xFF)
	b[1] = byte((p >> 16) & 0xFF)
	b[2] = byte((p >> 8) & 0xFF)
	b[3] = byte(p & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Long) store(w io.Writer) error {
	var b [8]byte
	b[0] = byte((p >> 56) & 0xFF)
	b[1] = byte((p >> 48) & 0xFF)
	b[2] = byte((p >> 40) & 0xFF)
	b[3] = byte((p >> 32) & 0xFF)
	b[4] = byte((p >> 24) & 0xFF)
	b[5] = byte((p >> 16) & 0xFF)
	b[6] = byte((p >> 8) & 0xFF)
	b[7] = byte(p & 0xFF)
	_, err := w.Write(b[0:8])
	return err
}

func (p Float) store(w io.Writer) error {
	var b [4]byte
	f := math.Float32bits(float32(p))
	b[0] = byte((f >> 24) & 0xFF)
	b[1] = byte((f >> 16) & 0xFF)
	b[2] = byte((f >> 8) & 0xFF)
	b[3] = byte(f & 0xFF)
	_, err := w.Write(b[0:4])
	return err
}

func (p Double) store(w io.Writer) error {
	var b [8]byte
	f := math.Float64bits(float64(p))
	b[0] = byte((f >> 56) & 0xFF)
	b[1] = byte((f >> 48) & 0xFF)
	b[2] = byte((f >> 40) & 0xFF)
	b[3] = byte((f >> 32) & 0xFF)
	b[4] = byte((f >> 24) & 0xFF)
	b[5] = byte((f >> 16) & 0xFF)
	b[6] = byte((f >> 8) & 0xFF)
	b[7] = byte(f & 0xFF)
	_, err := w.Write(b[0:8])
	return err
}

func (p ByteArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	_, err = w.Write(*(*[]byte)(unsafe.Pointer(&p)))
	return err
}

func (p String) store(w io.Writer) error {
	if len(p) > 32767 {
		return fmt.Errorf("can't store %d-byte string", len(p))
	}
	sh := Short(len(p))
	err := sh.store(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(p))
	return err
}

func (p List) store(w io.Writer) error {
	err := Byte(p.typ).store(w)
	if err != nil {
		return err
	}
	l := Int(len(p.data))
	err = l.store(w)
	if err != nil {
		return err
	}
	return p.storeData(w)
}

func (p Compound) store(w io.Writer) error {
	for k, v := range p {
		n := NBT{Name: k, Type: v.Type, payload: v.payload}
		err := n.Store(w)
		if err != nil {
			return err
		}
	}
	end := NBT{Name: "", Type: TagEnd, payload: nil}
	return end.Store(w)
}

func (p IntArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	for _, i := range p {
		err = i.store(w)
		if err != nil {
			return err
		}
	}
	return err
}

func (p LongArray) store(w io.Writer) error {
	l := Int(len(p))
	err := l.store(w)
	if err != nil {
		return err
	}
	for _, i := range p {
		err = i.store(w)
		if err != nil {
			return err
		}
	}
	return err
}

// StoreCompressed writes n to w, compressed via gzip.
func StoreCompressed(w io.Writer, n NBT) error {
	comp := gzip.NewWriter(w)
	defer comp.Close()
	return StoreUncompressed(comp, n)
}

// StoreUncompressed writes n to w, not compressing it. This is not
// usually useful except for debugging.
func StoreUncompressed(w io.Writer, n NBT) error {
	return n.Store(w)
}

// Store is just an alias for StoreCompressed, since the NBT spec
// says everything is gzipped.
func Store(w io.Writer, n NBT) error {
	return StoreCompressed(w, n)
}
