package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_MergeCellConstructor(t *testing.T) {
	v := sml.NewCT_MergeCell()
	if v == nil {
		t.Errorf("sml.NewCT_MergeCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MergeCell should validate: %s", err)
	}
}

func TestCT_MergeCellMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MergeCell()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MergeCell()
	xml.Unmarshal(buf, v2)
}
