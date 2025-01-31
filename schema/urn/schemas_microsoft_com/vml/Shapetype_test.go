package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestShapetypeConstructor(t *testing.T) {
	v := vml.NewShapetype()
	if v == nil {
		t.Errorf("vml.NewShapetype must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Shapetype should validate: %s", err)
	}
}

func TestShapetypeMarshalUnmarshal(t *testing.T) {
	v := vml.NewShapetype()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewShapetype()
	xml.Unmarshal(buf, v2)
}
