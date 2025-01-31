package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_InkConstructor(t *testing.T) {
	v := vml.NewOfcCT_Ink()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Ink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Ink should validate: %s", err)
	}
}

func TestOfcCT_InkMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Ink()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Ink()
	xml.Unmarshal(buf, v2)
}
