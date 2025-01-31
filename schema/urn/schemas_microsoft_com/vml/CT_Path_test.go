package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_PathConstructor(t *testing.T) {
	v := vml.NewCT_Path()
	if v == nil {
		t.Errorf("vml.NewCT_Path must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Path should validate: %s", err)
	}
}

func TestCT_PathMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Path()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Path()
	xml.Unmarshal(buf, v2)
}
