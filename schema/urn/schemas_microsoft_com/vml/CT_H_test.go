package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_HConstructor(t *testing.T) {
	v := vml.NewCT_H()
	if v == nil {
		t.Errorf("vml.NewCT_H must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_H should validate: %s", err)
	}
}

func TestCT_HMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_H()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_H()
	xml.Unmarshal(buf, v2)
}
