package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_GraphicalObjectConstructor(t *testing.T) {
	v := dml.NewCT_GraphicalObject()
	if v == nil {
		t.Errorf("dml.NewCT_GraphicalObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GraphicalObject should validate: %s", err)
	}
}

func TestCT_GraphicalObjectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GraphicalObject()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GraphicalObject()
	xml.Unmarshal(buf, v2)
}
