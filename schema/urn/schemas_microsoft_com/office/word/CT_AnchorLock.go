package word

import (
	"encoding/xml"
	"fmt"
)

type CT_AnchorLock struct {
}

func NewCT_AnchorLock() *CT_AnchorLock {
	ret := &CT_AnchorLock{}
	return ret
}

func (m *CT_AnchorLock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AnchorLock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_AnchorLock: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_AnchorLock and its children
func (m *CT_AnchorLock) Validate() error {
	return m.ValidateWithPath("CT_AnchorLock")
}

// ValidateWithPath validates the CT_AnchorLock and its children, prefixing error messages with path
func (m *CT_AnchorLock) ValidateWithPath(path string) error {
	return nil
}
