package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_MetadataRecordConstructor(t *testing.T) {
	v := sml.NewCT_MetadataRecord()
	if v == nil {
		t.Errorf("sml.NewCT_MetadataRecord must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MetadataRecord should validate: %s", err)
	}
}

func TestCT_MetadataRecordMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MetadataRecord()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MetadataRecord()
	xml.Unmarshal(buf, v2)
}
