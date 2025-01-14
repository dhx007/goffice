package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_BorderConstructor(t *testing.T) {
	v := sml.NewCT_Border()
	if v == nil {
		t.Errorf("sml.NewCT_Border must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Border should validate: %s", err)
	}
}

func TestCT_BorderMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Border()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Border()
	xml.Unmarshal(buf, v2)
}
