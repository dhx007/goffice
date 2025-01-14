package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestWdCT_WordprocessingContentPartNonVisualConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingContentPartNonVisual()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingContentPartNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingContentPartNonVisual should validate: %s", err)
	}
}

func TestWdCT_WordprocessingContentPartNonVisualMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingContentPartNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingContentPartNonVisual()
	xml.Unmarshal(buf, v2)
}
