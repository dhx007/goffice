package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestDialogsheetConstructor(t *testing.T) {
	v := sml.NewDialogsheet()
	if v == nil {
		t.Errorf("sml.NewDialogsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Dialogsheet should validate: %s", err)
	}
}

func TestDialogsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewDialogsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewDialogsheet()
	xml.Unmarshal(buf, v2)
}
