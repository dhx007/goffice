package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/pml"
)

func TestCT_TLBuildParagraphConstructor(t *testing.T) {
	v := pml.NewCT_TLBuildParagraph()
	if v == nil {
		t.Errorf("pml.NewCT_TLBuildParagraph must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLBuildParagraph should validate: %s", err)
	}
}

func TestCT_TLBuildParagraphMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLBuildParagraph()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLBuildParagraph()
	xml.Unmarshal(buf, v2)
}
