package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pml"
)

func TestCT_ConnectorNonVisualConstructor(t *testing.T) {
	v := pml.NewCT_ConnectorNonVisual()
	if v == nil {
		t.Errorf("pml.NewCT_ConnectorNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ConnectorNonVisual should validate: %s", err)
	}
}

func TestCT_ConnectorNonVisualMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ConnectorNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ConnectorNonVisual()
	xml.Unmarshal(buf, v2)
}
