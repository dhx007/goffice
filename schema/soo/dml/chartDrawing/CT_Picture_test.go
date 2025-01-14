package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/chartDrawing"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Picture()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
