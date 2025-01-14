package wml

import (
	"encoding/xml"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/ofc/math"
	"github.com/dhx007/goffice/schema/soo/schemaLibrary"
)

type Settings struct {
	CT_Settings
}

func NewSettings() *Settings {
	ret := &Settings{}
	ret.CT_Settings = *NewCT_Settings()
	return ret
}

func (m *Settings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:settings"
	return m.CT_Settings.MarshalXML(e, start)
}

func (m *Settings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Settings = *NewCT_Settings()
lSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "writeProtection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "writeProtection"}:
				m.WriteProtection = NewCT_WriteProtection()
				if err := d.DecodeElement(m.WriteProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "view"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "view"}:
				m.View = NewCT_View()
				if err := d.DecodeElement(m.View, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "zoom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "zoom"}:
				m.Zoom = NewCT_Zoom()
				if err := d.DecodeElement(m.Zoom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "removePersonalInformation"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "removePersonalInformation"}:
				m.RemovePersonalInformation = NewCT_OnOff()
				if err := d.DecodeElement(m.RemovePersonalInformation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "removeDateAndTime"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "removeDateAndTime"}:
				m.RemoveDateAndTime = NewCT_OnOff()
				if err := d.DecodeElement(m.RemoveDateAndTime, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotDisplayPageBoundaries"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotDisplayPageBoundaries"}:
				m.DoNotDisplayPageBoundaries = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDisplayPageBoundaries, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayBackgroundShape"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "displayBackgroundShape"}:
				m.DisplayBackgroundShape = NewCT_OnOff()
				if err := d.DecodeElement(m.DisplayBackgroundShape, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printPostScriptOverText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "printPostScriptOverText"}:
				m.PrintPostScriptOverText = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintPostScriptOverText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printFractionalCharacterWidth"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "printFractionalCharacterWidth"}:
				m.PrintFractionalCharacterWidth = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFractionalCharacterWidth, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printFormsData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "printFormsData"}:
				m.PrintFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintFormsData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "embedTrueTypeFonts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "embedTrueTypeFonts"}:
				m.EmbedTrueTypeFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedTrueTypeFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "embedSystemFonts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "embedSystemFonts"}:
				m.EmbedSystemFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.EmbedSystemFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveSubsetFonts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveSubsetFonts"}:
				m.SaveSubsetFonts = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSubsetFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveFormsData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveFormsData"}:
				m.SaveFormsData = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveFormsData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "mirrorMargins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "mirrorMargins"}:
				m.MirrorMargins = NewCT_OnOff()
				if err := d.DecodeElement(m.MirrorMargins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alignBordersAndEdges"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "alignBordersAndEdges"}:
				m.AlignBordersAndEdges = NewCT_OnOff()
				if err := d.DecodeElement(m.AlignBordersAndEdges, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bordersDoNotSurroundHeader"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bordersDoNotSurroundHeader"}:
				m.BordersDoNotSurroundHeader = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundHeader, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bordersDoNotSurroundFooter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bordersDoNotSurroundFooter"}:
				m.BordersDoNotSurroundFooter = NewCT_OnOff()
				if err := d.DecodeElement(m.BordersDoNotSurroundFooter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gutterAtTop"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "gutterAtTop"}:
				m.GutterAtTop = NewCT_OnOff()
				if err := d.DecodeElement(m.GutterAtTop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hideSpellingErrors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hideSpellingErrors"}:
				m.HideSpellingErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideSpellingErrors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hideGrammaticalErrors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hideGrammaticalErrors"}:
				m.HideGrammaticalErrors = NewCT_OnOff()
				if err := d.DecodeElement(m.HideGrammaticalErrors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "activeWritingStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "activeWritingStyle"}:
				tmp := NewCT_WritingStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ActiveWritingStyle = append(m.ActiveWritingStyle, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofState"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "proofState"}:
				m.ProofState = NewCT_Proof()
				if err := d.DecodeElement(m.ProofState, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "formsDesign"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "formsDesign"}:
				m.FormsDesign = NewCT_OnOff()
				if err := d.DecodeElement(m.FormsDesign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attachedTemplate"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "attachedTemplate"}:
				m.AttachedTemplate = NewCT_Rel()
				if err := d.DecodeElement(m.AttachedTemplate, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "linkStyles"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "linkStyles"}:
				m.LinkStyles = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkStyles, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "stylePaneFormatFilter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "stylePaneFormatFilter"}:
				m.StylePaneFormatFilter = NewCT_StylePaneFilter()
				if err := d.DecodeElement(m.StylePaneFormatFilter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "stylePaneSortMethod"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "stylePaneSortMethod"}:
				m.StylePaneSortMethod = NewCT_StyleSort()
				if err := d.DecodeElement(m.StylePaneSortMethod, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "documentType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "documentType"}:
				m.DocumentType = NewCT_DocType()
				if err := d.DecodeElement(m.DocumentType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "mailMerge"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "mailMerge"}:
				m.MailMerge = NewCT_MailMerge()
				if err := d.DecodeElement(m.MailMerge, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "revisionView"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "revisionView"}:
				m.RevisionView = NewCT_TrackChangesView()
				if err := d.DecodeElement(m.RevisionView, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "trackRevisions"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "trackRevisions"}:
				m.TrackRevisions = NewCT_OnOff()
				if err := d.DecodeElement(m.TrackRevisions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotTrackMoves"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotTrackMoves"}:
				m.DoNotTrackMoves = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackMoves, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotTrackFormatting"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotTrackFormatting"}:
				m.DoNotTrackFormatting = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotTrackFormatting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "documentProtection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "documentProtection"}:
				m.DocumentProtection = NewCT_DocProtect()
				if err := d.DecodeElement(m.DocumentProtection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoFormatOverride"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "autoFormatOverride"}:
				m.AutoFormatOverride = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoFormatOverride, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "styleLockTheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "styleLockTheme"}:
				m.StyleLockTheme = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockTheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "styleLockQFSet"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "styleLockQFSet"}:
				m.StyleLockQFSet = NewCT_OnOff()
				if err := d.DecodeElement(m.StyleLockQFSet, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "defaultTabStop"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "defaultTabStop"}:
				m.DefaultTabStop = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DefaultTabStop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoHyphenation"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "autoHyphenation"}:
				m.AutoHyphenation = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoHyphenation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "consecutiveHyphenLimit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "consecutiveHyphenLimit"}:
				m.ConsecutiveHyphenLimit = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ConsecutiveHyphenLimit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hyphenationZone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hyphenationZone"}:
				m.HyphenationZone = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.HyphenationZone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotHyphenateCaps"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotHyphenateCaps"}:
				m.DoNotHyphenateCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotHyphenateCaps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "showEnvelope"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "showEnvelope"}:
				m.ShowEnvelope = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowEnvelope, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "summaryLength"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "summaryLength"}:
				m.SummaryLength = NewCT_DecimalNumberOrPrecent()
				if err := d.DecodeElement(m.SummaryLength, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "clickAndTypeStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "clickAndTypeStyle"}:
				m.ClickAndTypeStyle = NewCT_String()
				if err := d.DecodeElement(m.ClickAndTypeStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "defaultTableStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "defaultTableStyle"}:
				m.DefaultTableStyle = NewCT_String()
				if err := d.DecodeElement(m.DefaultTableStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "evenAndOddHeaders"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "evenAndOddHeaders"}:
				m.EvenAndOddHeaders = NewCT_OnOff()
				if err := d.DecodeElement(m.EvenAndOddHeaders, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldRevPrinting"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookFoldRevPrinting"}:
				m.BookFoldRevPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldRevPrinting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldPrinting"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookFoldPrinting"}:
				m.BookFoldPrinting = NewCT_OnOff()
				if err := d.DecodeElement(m.BookFoldPrinting, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookFoldPrintingSheets"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bookFoldPrintingSheets"}:
				m.BookFoldPrintingSheets = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.BookFoldPrintingSheets, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridHorizontalSpacing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawingGridHorizontalSpacing"}:
				m.DrawingGridHorizontalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalSpacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridVerticalSpacing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawingGridVerticalSpacing"}:
				m.DrawingGridVerticalSpacing = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalSpacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayHorizontalDrawingGridEvery"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "displayHorizontalDrawingGridEvery"}:
				m.DisplayHorizontalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayHorizontalDrawingGridEvery, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "displayVerticalDrawingGridEvery"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "displayVerticalDrawingGridEvery"}:
				m.DisplayVerticalDrawingGridEvery = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.DisplayVerticalDrawingGridEvery, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotUseMarginsForDrawingGridOrigin"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotUseMarginsForDrawingGridOrigin"}:
				m.DoNotUseMarginsForDrawingGridOrigin = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseMarginsForDrawingGridOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridHorizontalOrigin"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawingGridHorizontalOrigin"}:
				m.DrawingGridHorizontalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridHorizontalOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawingGridVerticalOrigin"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawingGridVerticalOrigin"}:
				m.DrawingGridVerticalOrigin = NewCT_TwipsMeasure()
				if err := d.DecodeElement(m.DrawingGridVerticalOrigin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotShadeFormData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotShadeFormData"}:
				m.DoNotShadeFormData = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotShadeFormData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noPunctuationKerning"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noPunctuationKerning"}:
				m.NoPunctuationKerning = NewCT_OnOff()
				if err := d.DecodeElement(m.NoPunctuationKerning, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "characterSpacingControl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "characterSpacingControl"}:
				m.CharacterSpacingControl = NewCT_CharacterSpacing()
				if err := d.DecodeElement(m.CharacterSpacingControl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "printTwoOnOne"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "printTwoOnOne"}:
				m.PrintTwoOnOne = NewCT_OnOff()
				if err := d.DecodeElement(m.PrintTwoOnOne, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "strictFirstAndLastChars"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "strictFirstAndLastChars"}:
				m.StrictFirstAndLastChars = NewCT_OnOff()
				if err := d.DecodeElement(m.StrictFirstAndLastChars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noLineBreaksAfter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noLineBreaksAfter"}:
				m.NoLineBreaksAfter = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksAfter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noLineBreaksBefore"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noLineBreaksBefore"}:
				m.NoLineBreaksBefore = NewCT_Kinsoku()
				if err := d.DecodeElement(m.NoLineBreaksBefore, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "savePreviewPicture"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "savePreviewPicture"}:
				m.SavePreviewPicture = NewCT_OnOff()
				if err := d.DecodeElement(m.SavePreviewPicture, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotValidateAgainstSchema"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotValidateAgainstSchema"}:
				m.DoNotValidateAgainstSchema = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotValidateAgainstSchema, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveInvalidXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveInvalidXml"}:
				m.SaveInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveInvalidXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ignoreMixedContent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ignoreMixedContent"}:
				m.IgnoreMixedContent = NewCT_OnOff()
				if err := d.DecodeElement(m.IgnoreMixedContent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alwaysShowPlaceholderText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "alwaysShowPlaceholderText"}:
				m.AlwaysShowPlaceholderText = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysShowPlaceholderText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotDemarcateInvalidXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotDemarcateInvalidXml"}:
				m.DoNotDemarcateInvalidXml = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotDemarcateInvalidXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveXmlDataOnly"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveXmlDataOnly"}:
				m.SaveXmlDataOnly = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveXmlDataOnly, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "useXSLTWhenSaving"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "useXSLTWhenSaving"}:
				m.UseXSLTWhenSaving = NewCT_OnOff()
				if err := d.DecodeElement(m.UseXSLTWhenSaving, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveThroughXslt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveThroughXslt"}:
				m.SaveThroughXslt = NewCT_SaveThroughXslt()
				if err := d.DecodeElement(m.SaveThroughXslt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "showXMLTags"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "showXMLTags"}:
				m.ShowXMLTags = NewCT_OnOff()
				if err := d.DecodeElement(m.ShowXMLTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "alwaysMergeEmptyNamespace"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "alwaysMergeEmptyNamespace"}:
				m.AlwaysMergeEmptyNamespace = NewCT_OnOff()
				if err := d.DecodeElement(m.AlwaysMergeEmptyNamespace, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "updateFields"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "updateFields"}:
				m.UpdateFields = NewCT_OnOff()
				if err := d.DecodeElement(m.UpdateFields, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hdrShapeDefaults"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hdrShapeDefaults"}:
				m.HdrShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.HdrShapeDefaults, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "footnotePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "footnotePr"}:
				m.FootnotePr = NewCT_FtnDocProps()
				if err := d.DecodeElement(m.FootnotePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "endnotePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "endnotePr"}:
				m.EndnotePr = NewCT_EdnDocProps()
				if err := d.DecodeElement(m.EndnotePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "compat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "compat"}:
				m.Compat = NewCT_Compat()
				if err := d.DecodeElement(m.Compat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "docVars"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "docVars"}:
				m.DocVars = NewCT_DocVars()
				if err := d.DecodeElement(m.DocVars, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rsids"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rsids"}:
				m.Rsids = NewCT_DocRsids()
				if err := d.DecodeElement(m.Rsids, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mathPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "mathPr"}:
				m.MathPr = math.NewMathPr()
				if err := d.DecodeElement(m.MathPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "attachedSchema"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "attachedSchema"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AttachedSchema = append(m.AttachedSchema, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "themeFontLang"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "themeFontLang"}:
				m.ThemeFontLang = NewCT_Language()
				if err := d.DecodeElement(m.ThemeFontLang, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "clrSchemeMapping"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "clrSchemeMapping"}:
				m.ClrSchemeMapping = NewCT_ColorSchemeMapping()
				if err := d.DecodeElement(m.ClrSchemeMapping, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotIncludeSubdocsInStats"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotIncludeSubdocsInStats"}:
				m.DoNotIncludeSubdocsInStats = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotIncludeSubdocsInStats, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotAutoCompressPictures"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotAutoCompressPictures"}:
				m.DoNotAutoCompressPictures = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotAutoCompressPictures, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "forceUpgrade"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "forceUpgrade"}:
				m.ForceUpgrade = NewCT_Empty()
				if err := d.DecodeElement(m.ForceUpgrade, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "captions"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "captions"}:
				m.Captions = NewCT_Captions()
				if err := d.DecodeElement(m.Captions, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "readModeInkLockDown"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "readModeInkLockDown"}:
				m.ReadModeInkLockDown = NewCT_ReadingModeInkLockDown()
				if err := d.DecodeElement(m.ReadModeInkLockDown, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smartTagType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "smartTagType"}:
				tmp := NewCT_SmartTagType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SmartTagType = append(m.SmartTagType, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/schemaLibrary/2006/main", Local: "schemaLibrary"}:
				m.SchemaLibrary = schemaLibrary.NewSchemaLibrary()
				if err := d.DecodeElement(m.SchemaLibrary, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "shapeDefaults"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "shapeDefaults"}:
				m.ShapeDefaults = NewCT_ShapeDefaults()
				if err := d.DecodeElement(m.ShapeDefaults, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotEmbedSmartTags"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotEmbedSmartTags"}:
				m.DoNotEmbedSmartTags = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotEmbedSmartTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "decimalSymbol"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "decimalSymbol"}:
				m.DecimalSymbol = NewCT_String()
				if err := d.DecodeElement(m.DecimalSymbol, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "listSeparator"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "listSeparator"}:
				m.ListSeparator = NewCT_String()
				if err := d.DecodeElement(m.ListSeparator, &el); err != nil {
					return err
				}
			default:
				any := &goffice.XSDAny{}
				if err := d.DecodeElement(any, &el); err != nil {
					return err
				}
				m.Extra = append(m.Extra, any)
			}
		case xml.EndElement:
			break lSettings
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Settings and its children
func (m *Settings) Validate() error {
	return m.ValidateWithPath("Settings")
}

// ValidateWithPath validates the Settings and its children, prefixing error messages with path
func (m *Settings) ValidateWithPath(path string) error {
	if err := m.CT_Settings.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
