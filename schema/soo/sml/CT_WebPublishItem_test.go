package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_WebPublishItemConstructor(t *testing.T) {
	v := sml.NewCT_WebPublishItem()
	if v == nil {
		t.Errorf("sml.NewCT_WebPublishItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPublishItem should validate: %s", err)
	}
}

func TestCT_WebPublishItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPublishItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPublishItem()
	xml.Unmarshal(buf, v2)
}
