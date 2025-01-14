package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_GvmlShapeNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlShapeNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
