package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Shape()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
