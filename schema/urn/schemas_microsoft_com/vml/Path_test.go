package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestPathConstructor(t *testing.T) {
	v := vml.NewPath()
	if v == nil {
		t.Errorf("vml.NewPath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Path should validate: %s", err)
	}
}

func TestPathMarshalUnmarshal(t *testing.T) {
	v := vml.NewPath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewPath()
	xml.Unmarshal(buf, v2)
}
