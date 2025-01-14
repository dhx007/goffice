package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestWorksheetConstructor(t *testing.T) {
	v := sml.NewWorksheet()
	if v == nil {
		t.Errorf("sml.NewWorksheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Worksheet should validate: %s", err)
	}
}

func TestWorksheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewWorksheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewWorksheet()
	xml.Unmarshal(buf, v2)
}
