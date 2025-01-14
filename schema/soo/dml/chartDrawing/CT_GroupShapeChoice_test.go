package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_GroupShapeChoiceConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeChoice()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
