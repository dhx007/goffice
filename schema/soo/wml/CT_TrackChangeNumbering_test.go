package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestCT_TrackChangeNumberingConstructor(t *testing.T) {
	v := wml.NewCT_TrackChangeNumbering()
	if v == nil {
		t.Errorf("wml.NewCT_TrackChangeNumbering must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TrackChangeNumbering should validate: %s", err)
	}
}

func TestCT_TrackChangeNumberingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TrackChangeNumbering()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TrackChangeNumbering()
	xml.Unmarshal(buf, v2)
}
