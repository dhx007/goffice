package picture_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/dml/picture"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := picture.NewCT_Picture()
	if v == nil {
		t.Errorf("picture.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed picture.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := picture.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := picture.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
