package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcExtrusionConstructor(t *testing.T) {
	v := vml.NewOfcExtrusion()
	if v == nil {
		t.Errorf("vml.NewOfcExtrusion must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcExtrusion should validate: %s", err)
	}
}

func TestOfcExtrusionMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcExtrusion()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcExtrusion()
	xml.Unmarshal(buf, v2)
}
