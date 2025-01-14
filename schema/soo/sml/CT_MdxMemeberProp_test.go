package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_MdxMemeberPropConstructor(t *testing.T) {
	v := sml.NewCT_MdxMemeberProp()
	if v == nil {
		t.Errorf("sml.NewCT_MdxMemeberProp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MdxMemeberProp should validate: %s", err)
	}
}

func TestCT_MdxMemeberPropMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MdxMemeberProp()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MdxMemeberProp()
	xml.Unmarshal(buf, v2)
}
