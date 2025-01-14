package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
)

type EG_RangeMarkupElements struct {
	// Bookmark Start
	BookmarkStart *CT_Bookmark
	// Bookmark End
	BookmarkEnd *CT_MarkupRange
	// Move Source Location Container - Start
	MoveFromRangeStart *CT_MoveBookmark
	// Move Source Location Container - End
	MoveFromRangeEnd *CT_MarkupRange
	// Move Destination Location Container - Start
	MoveToRangeStart *CT_MoveBookmark
	// Move Destination Location Container - End
	MoveToRangeEnd *CT_MarkupRange
	// Comment Anchor Range Start
	CommentRangeStart *CT_MarkupRange
	// Comment Anchor Range End
	CommentRangeEnd *CT_MarkupRange
	// Custom XML Markup Insertion Start
	CustomXmlInsRangeStart *CT_TrackChange
	// Custom XML Markup Insertion End
	CustomXmlInsRangeEnd *CT_Markup
	// Custom XML Markup Deletion Start
	CustomXmlDelRangeStart *CT_TrackChange
	// Custom XML Markup Deletion End
	CustomXmlDelRangeEnd *CT_Markup
	// Custom XML Markup Move Source Start
	CustomXmlMoveFromRangeStart *CT_TrackChange
	// Custom XML Markup Move Source End
	CustomXmlMoveFromRangeEnd *CT_Markup
	// Custom XML Markup Move Destination Location Start
	CustomXmlMoveToRangeStart *CT_TrackChange
	// Custom XML Markup Move Destination Location End
	CustomXmlMoveToRangeEnd *CT_Markup
}

func NewEG_RangeMarkupElements() *EG_RangeMarkupElements {
	ret := &EG_RangeMarkupElements{}
	return ret
}

func (m *EG_RangeMarkupElements) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BookmarkStart != nil {
		sebookmarkStart := xml.StartElement{Name: xml.Name{Local: "w:bookmarkStart"}}
		e.EncodeElement(m.BookmarkStart, sebookmarkStart)
	}
	if m.BookmarkEnd != nil {
		sebookmarkEnd := xml.StartElement{Name: xml.Name{Local: "w:bookmarkEnd"}}
		e.EncodeElement(m.BookmarkEnd, sebookmarkEnd)
	}
	if m.MoveFromRangeStart != nil {
		semoveFromRangeStart := xml.StartElement{Name: xml.Name{Local: "w:moveFromRangeStart"}}
		e.EncodeElement(m.MoveFromRangeStart, semoveFromRangeStart)
	}
	if m.MoveFromRangeEnd != nil {
		semoveFromRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:moveFromRangeEnd"}}
		e.EncodeElement(m.MoveFromRangeEnd, semoveFromRangeEnd)
	}
	if m.MoveToRangeStart != nil {
		semoveToRangeStart := xml.StartElement{Name: xml.Name{Local: "w:moveToRangeStart"}}
		e.EncodeElement(m.MoveToRangeStart, semoveToRangeStart)
	}
	if m.MoveToRangeEnd != nil {
		semoveToRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:moveToRangeEnd"}}
		e.EncodeElement(m.MoveToRangeEnd, semoveToRangeEnd)
	}
	if m.CommentRangeStart != nil {
		secommentRangeStart := xml.StartElement{Name: xml.Name{Local: "w:commentRangeStart"}}
		e.EncodeElement(m.CommentRangeStart, secommentRangeStart)
	}
	if m.CommentRangeEnd != nil {
		secommentRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:commentRangeEnd"}}
		e.EncodeElement(m.CommentRangeEnd, secommentRangeEnd)
	}
	if m.CustomXmlInsRangeStart != nil {
		secustomXmlInsRangeStart := xml.StartElement{Name: xml.Name{Local: "w:customXmlInsRangeStart"}}
		e.EncodeElement(m.CustomXmlInsRangeStart, secustomXmlInsRangeStart)
	}
	if m.CustomXmlInsRangeEnd != nil {
		secustomXmlInsRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:customXmlInsRangeEnd"}}
		e.EncodeElement(m.CustomXmlInsRangeEnd, secustomXmlInsRangeEnd)
	}
	if m.CustomXmlDelRangeStart != nil {
		secustomXmlDelRangeStart := xml.StartElement{Name: xml.Name{Local: "w:customXmlDelRangeStart"}}
		e.EncodeElement(m.CustomXmlDelRangeStart, secustomXmlDelRangeStart)
	}
	if m.CustomXmlDelRangeEnd != nil {
		secustomXmlDelRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:customXmlDelRangeEnd"}}
		e.EncodeElement(m.CustomXmlDelRangeEnd, secustomXmlDelRangeEnd)
	}
	if m.CustomXmlMoveFromRangeStart != nil {
		secustomXmlMoveFromRangeStart := xml.StartElement{Name: xml.Name{Local: "w:customXmlMoveFromRangeStart"}}
		e.EncodeElement(m.CustomXmlMoveFromRangeStart, secustomXmlMoveFromRangeStart)
	}
	if m.CustomXmlMoveFromRangeEnd != nil {
		secustomXmlMoveFromRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:customXmlMoveFromRangeEnd"}}
		e.EncodeElement(m.CustomXmlMoveFromRangeEnd, secustomXmlMoveFromRangeEnd)
	}
	if m.CustomXmlMoveToRangeStart != nil {
		secustomXmlMoveToRangeStart := xml.StartElement{Name: xml.Name{Local: "w:customXmlMoveToRangeStart"}}
		e.EncodeElement(m.CustomXmlMoveToRangeStart, secustomXmlMoveToRangeStart)
	}
	if m.CustomXmlMoveToRangeEnd != nil {
		secustomXmlMoveToRangeEnd := xml.StartElement{Name: xml.Name{Local: "w:customXmlMoveToRangeEnd"}}
		e.EncodeElement(m.CustomXmlMoveToRangeEnd, secustomXmlMoveToRangeEnd)
	}
	return nil
}

func (m *EG_RangeMarkupElements) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_RangeMarkupElements:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkStart"}:
				m.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(m.BookmarkStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookmarkEnd"}:
				m.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(m.BookmarkEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeStart"}:
				m.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(m.MoveFromRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFromRangeEnd"}:
				m.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(m.MoveFromRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeStart"}:
				m.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(m.MoveToRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveToRangeEnd"}:
				m.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(m.MoveToRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeStart"}:
				m.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(m.CommentRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "commentRangeEnd"}:
				m.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(m.CommentRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeStart"}:
				m.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(m.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlInsRangeEnd"}:
				m.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(m.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeStart"}:
				m.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(m.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlDelRangeEnd"}:
				m.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(m.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeStart"}:
				m.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(m.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveFromRangeEnd"}:
				m.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(m.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeStart"}:
				m.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(m.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "customXmlMoveToRangeEnd"}:
				m.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(m.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on EG_RangeMarkupElements %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_RangeMarkupElements
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_RangeMarkupElements and its children
func (m *EG_RangeMarkupElements) Validate() error {
	return m.ValidateWithPath("EG_RangeMarkupElements")
}

// ValidateWithPath validates the EG_RangeMarkupElements and its children, prefixing error messages with path
func (m *EG_RangeMarkupElements) ValidateWithPath(path string) error {
	if m.BookmarkStart != nil {
		if err := m.BookmarkStart.ValidateWithPath(path + "/BookmarkStart"); err != nil {
			return err
		}
	}
	if m.BookmarkEnd != nil {
		if err := m.BookmarkEnd.ValidateWithPath(path + "/BookmarkEnd"); err != nil {
			return err
		}
	}
	if m.MoveFromRangeStart != nil {
		if err := m.MoveFromRangeStart.ValidateWithPath(path + "/MoveFromRangeStart"); err != nil {
			return err
		}
	}
	if m.MoveFromRangeEnd != nil {
		if err := m.MoveFromRangeEnd.ValidateWithPath(path + "/MoveFromRangeEnd"); err != nil {
			return err
		}
	}
	if m.MoveToRangeStart != nil {
		if err := m.MoveToRangeStart.ValidateWithPath(path + "/MoveToRangeStart"); err != nil {
			return err
		}
	}
	if m.MoveToRangeEnd != nil {
		if err := m.MoveToRangeEnd.ValidateWithPath(path + "/MoveToRangeEnd"); err != nil {
			return err
		}
	}
	if m.CommentRangeStart != nil {
		if err := m.CommentRangeStart.ValidateWithPath(path + "/CommentRangeStart"); err != nil {
			return err
		}
	}
	if m.CommentRangeEnd != nil {
		if err := m.CommentRangeEnd.ValidateWithPath(path + "/CommentRangeEnd"); err != nil {
			return err
		}
	}
	if m.CustomXmlInsRangeStart != nil {
		if err := m.CustomXmlInsRangeStart.ValidateWithPath(path + "/CustomXmlInsRangeStart"); err != nil {
			return err
		}
	}
	if m.CustomXmlInsRangeEnd != nil {
		if err := m.CustomXmlInsRangeEnd.ValidateWithPath(path + "/CustomXmlInsRangeEnd"); err != nil {
			return err
		}
	}
	if m.CustomXmlDelRangeStart != nil {
		if err := m.CustomXmlDelRangeStart.ValidateWithPath(path + "/CustomXmlDelRangeStart"); err != nil {
			return err
		}
	}
	if m.CustomXmlDelRangeEnd != nil {
		if err := m.CustomXmlDelRangeEnd.ValidateWithPath(path + "/CustomXmlDelRangeEnd"); err != nil {
			return err
		}
	}
	if m.CustomXmlMoveFromRangeStart != nil {
		if err := m.CustomXmlMoveFromRangeStart.ValidateWithPath(path + "/CustomXmlMoveFromRangeStart"); err != nil {
			return err
		}
	}
	if m.CustomXmlMoveFromRangeEnd != nil {
		if err := m.CustomXmlMoveFromRangeEnd.ValidateWithPath(path + "/CustomXmlMoveFromRangeEnd"); err != nil {
			return err
		}
	}
	if m.CustomXmlMoveToRangeStart != nil {
		if err := m.CustomXmlMoveToRangeStart.ValidateWithPath(path + "/CustomXmlMoveToRangeStart"); err != nil {
			return err
		}
	}
	if m.CustomXmlMoveToRangeEnd != nil {
		if err := m.CustomXmlMoveToRangeEnd.ValidateWithPath(path + "/CustomXmlMoveToRangeEnd"); err != nil {
			return err
		}
	}
	return nil
}
