package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_ConnectorNonVisualConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_ConnectorNonVisual()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_ConnectorNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_ConnectorNonVisual should validate: %s", err)
	}
}

func TestCT_ConnectorNonVisualMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_ConnectorNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_ConnectorNonVisual()
	xml.Unmarshal(buf, v2)
}
