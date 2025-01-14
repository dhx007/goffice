package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Shape()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
