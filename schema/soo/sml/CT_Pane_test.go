package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_PaneConstructor(t *testing.T) {
	v := sml.NewCT_Pane()
	if v == nil {
		t.Errorf("sml.NewCT_Pane must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Pane should validate: %s", err)
	}
}

func TestCT_PaneMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Pane()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Pane()
	xml.Unmarshal(buf, v2)
}
