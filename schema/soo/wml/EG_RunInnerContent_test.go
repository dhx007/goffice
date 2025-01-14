package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/wml"
)

func TestEG_RunInnerContentConstructor(t *testing.T) {
	v := wml.NewEG_RunInnerContent()
	if v == nil {
		t.Errorf("wml.NewEG_RunInnerContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RunInnerContent should validate: %s", err)
	}
}

func TestEG_RunInnerContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RunInnerContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RunInnerContent()
	xml.Unmarshal(buf, v2)
}
