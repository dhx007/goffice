package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestEG_ExtensionListConstructor(t *testing.T) {
	v := sml.NewEG_ExtensionList()
	if v == nil {
		t.Errorf("sml.NewEG_ExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.EG_ExtensionList should validate: %s", err)
	}
}

func TestEG_ExtensionListMarshalUnmarshal(t *testing.T) {
	v := sml.NewEG_ExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewEG_ExtensionList()
	xml.Unmarshal(buf, v2)
}
