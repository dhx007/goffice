package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ComplementTransformConstructor(t *testing.T) {
	v := dml.NewCT_ComplementTransform()
	if v == nil {
		t.Errorf("dml.NewCT_ComplementTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ComplementTransform should validate: %s", err)
	}
}

func TestCT_ComplementTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ComplementTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ComplementTransform()
	xml.Unmarshal(buf, v2)
}
