package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_ConnectorNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_ConnectorNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_ConnectorNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_ConnectorNonVisual should validate: %s", err)
	}
}

func TestCT_ConnectorNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_ConnectorNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_ConnectorNonVisual()
	xml.Unmarshal(buf, v2)
}
