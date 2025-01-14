package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_PrintOptionsConstructor(t *testing.T) {
	v := sml.NewCT_PrintOptions()
	if v == nil {
		t.Errorf("sml.NewCT_PrintOptions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PrintOptions should validate: %s", err)
	}
}

func TestCT_PrintOptionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PrintOptions()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PrintOptions()
	xml.Unmarshal(buf, v2)
}
