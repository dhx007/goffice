package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_BaseStylesConstructor(t *testing.T) {
	v := dml.NewCT_BaseStyles()
	if v == nil {
		t.Errorf("dml.NewCT_BaseStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BaseStyles should validate: %s", err)
	}
}

func TestCT_BaseStylesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BaseStyles()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BaseStyles()
	xml.Unmarshal(buf, v2)
}
