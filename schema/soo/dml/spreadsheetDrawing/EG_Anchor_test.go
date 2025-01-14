package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/spreadsheetDrawing"
)

func TestEG_AnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewEG_Anchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewEG_Anchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.EG_Anchor should validate: %s", err)
	}
}

func TestEG_AnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewEG_Anchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewEG_Anchor()
	xml.Unmarshal(buf, v2)
}
