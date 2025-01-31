package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ShapeStyleConstructor(t *testing.T) {
	v := dml.NewCT_ShapeStyle()
	if v == nil {
		t.Errorf("dml.NewCT_ShapeStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ShapeStyle should validate: %s", err)
	}
}

func TestCT_ShapeStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ShapeStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ShapeStyle()
	xml.Unmarshal(buf, v2)
}
