package goffice_test

import (
	"testing"

	"github.com/dhx007/goffice"
)

func TestFloat32(t *testing.T) {
	exp := float32(1.234)
	got := goffice.Float32(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestFloat64(t *testing.T) {
	exp := 1.234
	got := goffice.Float64(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestUint64(t *testing.T) {
	exp := uint64(123)
	got := goffice.Uint64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint32(t *testing.T) {
	exp := uint32(123)
	got := goffice.Uint32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt64(t *testing.T) {
	exp := int64(123)
	got := goffice.Int64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt32(t *testing.T) {
	exp := int32(123)
	got := goffice.Int32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt8(t *testing.T) {
	exp := int8(123)
	got := goffice.Int8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestBool(t *testing.T) {
	exp := bool(true)
	got := goffice.Bool(exp)
	if *got != exp {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestString(t *testing.T) {
	exp := "foo"
	got := goffice.String(exp)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}

func TestUint8(t *testing.T) {
	exp := uint8(123)
	got := goffice.Uint8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint16(t *testing.T) {
	exp := uint16(123)
	got := goffice.Uint16(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestStringf(t *testing.T) {
	exp := "foobar123"
	got := goffice.Stringf("foo%s%d", "bar", 123)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
