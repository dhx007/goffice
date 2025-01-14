package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_LockConstructor(t *testing.T) {
	v := vml.NewOfcCT_Lock()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Lock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Lock should validate: %s", err)
	}
}

func TestOfcCT_LockMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Lock()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Lock()
	xml.Unmarshal(buf, v2)
}
