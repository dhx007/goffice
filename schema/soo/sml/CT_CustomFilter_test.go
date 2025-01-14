package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_CustomFilterConstructor(t *testing.T) {
	v := sml.NewCT_CustomFilter()
	if v == nil {
		t.Errorf("sml.NewCT_CustomFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomFilter should validate: %s", err)
	}
}

func TestCT_CustomFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomFilter()
	xml.Unmarshal(buf, v2)
}
