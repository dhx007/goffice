package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestCT_CommentPrConstructor(t *testing.T) {
	v := sml.NewCT_CommentPr()
	if v == nil {
		t.Errorf("sml.NewCT_CommentPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CommentPr should validate: %s", err)
	}
}

func TestCT_CommentPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CommentPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CommentPr()
	xml.Unmarshal(buf, v2)
}
