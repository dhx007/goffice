package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_MapInfoConstructor(t *testing.T) {
	v := sml.NewCT_MapInfo()
	if v == nil {
		t.Errorf("sml.NewCT_MapInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MapInfo should validate: %s", err)
	}
}

func TestCT_MapInfoMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MapInfo()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MapInfo()
	xml.Unmarshal(buf, v2)
}
