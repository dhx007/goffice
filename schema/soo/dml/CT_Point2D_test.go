package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_Point2DConstructor(t *testing.T) {
	v := dml.NewCT_Point2D()
	if v == nil {
		t.Errorf("dml.NewCT_Point2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Point2D should validate: %s", err)
	}
}

func TestCT_Point2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Point2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Point2D()
	xml.Unmarshal(buf, v2)
}
