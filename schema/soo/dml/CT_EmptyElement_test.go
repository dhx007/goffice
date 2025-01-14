package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_EmptyElementConstructor(t *testing.T) {
	v := dml.NewCT_EmptyElement()
	if v == nil {
		t.Errorf("dml.NewCT_EmptyElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EmptyElement should validate: %s", err)
	}
}

func TestCT_EmptyElementMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EmptyElement()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EmptyElement()
	xml.Unmarshal(buf, v2)
}
