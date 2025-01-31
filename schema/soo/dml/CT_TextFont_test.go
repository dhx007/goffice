package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml"
)

func TestCT_TextFontConstructor(t *testing.T) {
	v := dml.NewCT_TextFont()
	if v == nil {
		t.Errorf("dml.NewCT_TextFont must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextFont should validate: %s", err)
	}
}

func TestCT_TextFontMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextFont()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextFont()
	xml.Unmarshal(buf, v2)
}
