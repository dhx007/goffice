package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_RelSizeAnchorConstructor(t *testing.T) {
	v := chartDrawing.NewCT_RelSizeAnchor()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_RelSizeAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_RelSizeAnchor should validate: %s", err)
	}
}

func TestCT_RelSizeAnchorMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_RelSizeAnchor()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_RelSizeAnchor()
	xml.Unmarshal(buf, v2)
}
