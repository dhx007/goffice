package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/dhx007/goffice/schema/soo/sml"
)

func TestSingleXmlCellsConstructor(t *testing.T) {
	v := sml.NewSingleXmlCells()
	if v == nil {
		t.Errorf("sml.NewSingleXmlCells must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.SingleXmlCells should validate: %s", err)
	}
}

func TestSingleXmlCellsMarshalUnmarshal(t *testing.T) {
	v := sml.NewSingleXmlCells()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewSingleXmlCells()
	xml.Unmarshal(buf, v2)
}
