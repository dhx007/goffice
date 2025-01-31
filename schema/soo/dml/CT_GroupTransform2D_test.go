package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_GroupTransform2DConstructor(t *testing.T) {
	v := dml.NewCT_GroupTransform2D()
	if v == nil {
		t.Errorf("dml.NewCT_GroupTransform2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GroupTransform2D should validate: %s", err)
	}
}

func TestCT_GroupTransform2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GroupTransform2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GroupTransform2D()
	xml.Unmarshal(buf, v2)
}
