package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_FunctionGroupConstructor(t *testing.T) {
	v := sml.NewCT_FunctionGroup()
	if v == nil {
		t.Errorf("sml.NewCT_FunctionGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FunctionGroup should validate: %s", err)
	}
}

func TestCT_FunctionGroupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FunctionGroup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FunctionGroup()
	xml.Unmarshal(buf, v2)
}
