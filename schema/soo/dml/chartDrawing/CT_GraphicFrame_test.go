package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_GraphicFrameConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GraphicFrame()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GraphicFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GraphicFrame should validate: %s", err)
	}
}

func TestCT_GraphicFrameMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GraphicFrame()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GraphicFrame()
	xml.Unmarshal(buf, v2)
}
