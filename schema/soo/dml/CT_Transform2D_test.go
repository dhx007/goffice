package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_Transform2DConstructor(t *testing.T) {
	v := dml.NewCT_Transform2D()
	if v == nil {
		t.Errorf("dml.NewCT_Transform2D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Transform2D should validate: %s", err)
	}
}

func TestCT_Transform2DMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Transform2D()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Transform2D()
	xml.Unmarshal(buf, v2)
}
