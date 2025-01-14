package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_MeasureGroupsConstructor(t *testing.T) {
	v := sml.NewCT_MeasureGroups()
	if v == nil {
		t.Errorf("sml.NewCT_MeasureGroups must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MeasureGroups should validate: %s", err)
	}
}

func TestCT_MeasureGroupsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MeasureGroups()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MeasureGroups()
	xml.Unmarshal(buf, v2)
}
