package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_AdjPoint2DConstructor(t *testing.T) {
	v := dml.NewCT_AdjPoint2D()
	if v == nil {
		t.Errorf("dml.NewCT_AdjPoint2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AdjPoint2D should validate: %s", err)
	}
}

func TestCT_AdjPoint2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AdjPoint2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AdjPoint2D()
	xml.Unmarshal(buf, v2)
}
