package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_ShapeNonVisualConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_ShapeNonVisual()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_ShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_ShapeNonVisual should validate: %s", err)
	}
}

func TestCT_ShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_ShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_ShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
