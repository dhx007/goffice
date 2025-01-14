package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_ScRgbColorConstructor(t *testing.T) {
	v := dml.NewCT_ScRgbColor()
	if v == nil {
		t.Errorf("dml.NewCT_ScRgbColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ScRgbColor should validate: %s", err)
	}
}

func TestCT_ScRgbColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ScRgbColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ScRgbColor()
	xml.Unmarshal(buf, v2)
}
