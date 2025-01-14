package excel

import (
	"encoding/xml"
	"fmt"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/schema/soo/ofc/sharedTypes"
)

type CT_ClientData struct {
	ObjectTypeAttr ST_ObjectType
	MoveWithCells  sharedTypes.ST_TrueFalseBlank
	SizeWithCells  sharedTypes.ST_TrueFalseBlank
	Anchor         *string
	Locked         sharedTypes.ST_TrueFalseBlank
	DefaultSize    sharedTypes.ST_TrueFalseBlank
	PrintObject    sharedTypes.ST_TrueFalseBlank
	Disabled       sharedTypes.ST_TrueFalseBlank
	AutoFill       sharedTypes.ST_TrueFalseBlank
	AutoLine       sharedTypes.ST_TrueFalseBlank
	AutoPict       sharedTypes.ST_TrueFalseBlank
	FmlaMacro      *string
	TextHAlign     *string
	TextVAlign     *string
	LockText       sharedTypes.ST_TrueFalseBlank
	JustLastX      sharedTypes.ST_TrueFalseBlank
	SecretEdit     sharedTypes.ST_TrueFalseBlank
	Default        sharedTypes.ST_TrueFalseBlank
	Help           sharedTypes.ST_TrueFalseBlank
	Cancel         sharedTypes.ST_TrueFalseBlank
	Dismiss        sharedTypes.ST_TrueFalseBlank
	Accel          *int64
	Accel2         *int64
	Row            *int64
	Column         *int64
	Visible        sharedTypes.ST_TrueFalseBlank
	RowHidden      sharedTypes.ST_TrueFalseBlank
	ColHidden      sharedTypes.ST_TrueFalseBlank
	VTEdit         *int64
	MultiLine      sharedTypes.ST_TrueFalseBlank
	VScroll        sharedTypes.ST_TrueFalseBlank
	ValidIds       sharedTypes.ST_TrueFalseBlank
	FmlaRange      *string
	WidthMin       *int64
	Sel            *int64
	NoThreeD2      sharedTypes.ST_TrueFalseBlank
	SelType        *string
	MultiSel       *string
	LCT            *string
	ListItem       *string
	DropStyle      *string
	Colored        sharedTypes.ST_TrueFalseBlank
	DropLines      *int64
	Checked        *int64
	FmlaLink       *string
	FmlaPict       *string
	NoThreeD       sharedTypes.ST_TrueFalseBlank
	FirstButton    sharedTypes.ST_TrueFalseBlank
	FmlaGroup      *string
	Val            *int64
	Min            *int64
	Max            *int64
	Inc            *int64
	Page           *int64
	Horiz          sharedTypes.ST_TrueFalseBlank
	Dx             *int64
	MapOCX         sharedTypes.ST_TrueFalseBlank
	CF             []string
	Camera         sharedTypes.ST_TrueFalseBlank
	RecalcAlways   sharedTypes.ST_TrueFalseBlank
	AutoScale      sharedTypes.ST_TrueFalseBlank
	DDE            sharedTypes.ST_TrueFalseBlank
	UIObj          sharedTypes.ST_TrueFalseBlank
	ScriptText     *string
	ScriptExtended *string
	ScriptLanguage *uint32
	ScriptLocation *uint32
	FmlaTxbx       *string
}

func NewCT_ClientData() *CT_ClientData {
	ret := &CT_ClientData{}
	ret.ObjectTypeAttr = ST_ObjectType(1)
	return ret
}

func (m *CT_ClientData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ObjectTypeAttr.MarshalXMLAttr(xml.Name{Local: "ObjectType"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.MoveWithCells != sharedTypes.ST_TrueFalseBlankUnset {
		seMoveWithCells := xml.StartElement{Name: xml.Name{Local: "x:MoveWithCells"}}
		e.EncodeElement(m.MoveWithCells, seMoveWithCells)
	}
	if m.SizeWithCells != sharedTypes.ST_TrueFalseBlankUnset {
		seSizeWithCells := xml.StartElement{Name: xml.Name{Local: "x:SizeWithCells"}}
		e.EncodeElement(m.SizeWithCells, seSizeWithCells)
	}
	if m.Anchor != nil {
		seAnchor := xml.StartElement{Name: xml.Name{Local: "x:Anchor"}}
		goffice.AddPreserveSpaceAttr(&seAnchor, *m.Anchor)
		e.EncodeElement(m.Anchor, seAnchor)
	}
	if m.Locked != sharedTypes.ST_TrueFalseBlankUnset {
		seLocked := xml.StartElement{Name: xml.Name{Local: "x:Locked"}}
		e.EncodeElement(m.Locked, seLocked)
	}
	if m.DefaultSize != sharedTypes.ST_TrueFalseBlankUnset {
		seDefaultSize := xml.StartElement{Name: xml.Name{Local: "x:DefaultSize"}}
		e.EncodeElement(m.DefaultSize, seDefaultSize)
	}
	if m.PrintObject != sharedTypes.ST_TrueFalseBlankUnset {
		sePrintObject := xml.StartElement{Name: xml.Name{Local: "x:PrintObject"}}
		e.EncodeElement(m.PrintObject, sePrintObject)
	}
	if m.Disabled != sharedTypes.ST_TrueFalseBlankUnset {
		seDisabled := xml.StartElement{Name: xml.Name{Local: "x:Disabled"}}
		e.EncodeElement(m.Disabled, seDisabled)
	}
	if m.AutoFill != sharedTypes.ST_TrueFalseBlankUnset {
		seAutoFill := xml.StartElement{Name: xml.Name{Local: "x:AutoFill"}}
		e.EncodeElement(m.AutoFill, seAutoFill)
	}
	if m.AutoLine != sharedTypes.ST_TrueFalseBlankUnset {
		seAutoLine := xml.StartElement{Name: xml.Name{Local: "x:AutoLine"}}
		e.EncodeElement(m.AutoLine, seAutoLine)
	}
	if m.AutoPict != sharedTypes.ST_TrueFalseBlankUnset {
		seAutoPict := xml.StartElement{Name: xml.Name{Local: "x:AutoPict"}}
		e.EncodeElement(m.AutoPict, seAutoPict)
	}
	if m.FmlaMacro != nil {
		seFmlaMacro := xml.StartElement{Name: xml.Name{Local: "x:FmlaMacro"}}
		goffice.AddPreserveSpaceAttr(&seFmlaMacro, *m.FmlaMacro)
		e.EncodeElement(m.FmlaMacro, seFmlaMacro)
	}
	if m.TextHAlign != nil {
		seTextHAlign := xml.StartElement{Name: xml.Name{Local: "x:TextHAlign"}}
		goffice.AddPreserveSpaceAttr(&seTextHAlign, *m.TextHAlign)
		e.EncodeElement(m.TextHAlign, seTextHAlign)
	}
	if m.TextVAlign != nil {
		seTextVAlign := xml.StartElement{Name: xml.Name{Local: "x:TextVAlign"}}
		goffice.AddPreserveSpaceAttr(&seTextVAlign, *m.TextVAlign)
		e.EncodeElement(m.TextVAlign, seTextVAlign)
	}
	if m.LockText != sharedTypes.ST_TrueFalseBlankUnset {
		seLockText := xml.StartElement{Name: xml.Name{Local: "x:LockText"}}
		e.EncodeElement(m.LockText, seLockText)
	}
	if m.JustLastX != sharedTypes.ST_TrueFalseBlankUnset {
		seJustLastX := xml.StartElement{Name: xml.Name{Local: "x:JustLastX"}}
		e.EncodeElement(m.JustLastX, seJustLastX)
	}
	if m.SecretEdit != sharedTypes.ST_TrueFalseBlankUnset {
		seSecretEdit := xml.StartElement{Name: xml.Name{Local: "x:SecretEdit"}}
		e.EncodeElement(m.SecretEdit, seSecretEdit)
	}
	if m.Default != sharedTypes.ST_TrueFalseBlankUnset {
		seDefault := xml.StartElement{Name: xml.Name{Local: "x:Default"}}
		e.EncodeElement(m.Default, seDefault)
	}
	if m.Help != sharedTypes.ST_TrueFalseBlankUnset {
		seHelp := xml.StartElement{Name: xml.Name{Local: "x:Help"}}
		e.EncodeElement(m.Help, seHelp)
	}
	if m.Cancel != sharedTypes.ST_TrueFalseBlankUnset {
		seCancel := xml.StartElement{Name: xml.Name{Local: "x:Cancel"}}
		e.EncodeElement(m.Cancel, seCancel)
	}
	if m.Dismiss != sharedTypes.ST_TrueFalseBlankUnset {
		seDismiss := xml.StartElement{Name: xml.Name{Local: "x:Dismiss"}}
		e.EncodeElement(m.Dismiss, seDismiss)
	}
	if m.Accel != nil {
		seAccel := xml.StartElement{Name: xml.Name{Local: "x:Accel"}}
		e.EncodeElement(m.Accel, seAccel)
	}
	if m.Accel2 != nil {
		seAccel2 := xml.StartElement{Name: xml.Name{Local: "x:Accel2"}}
		e.EncodeElement(m.Accel2, seAccel2)
	}
	if m.Row != nil {
		seRow := xml.StartElement{Name: xml.Name{Local: "x:Row"}}
		e.EncodeElement(m.Row, seRow)
	}
	if m.Column != nil {
		seColumn := xml.StartElement{Name: xml.Name{Local: "x:Column"}}
		e.EncodeElement(m.Column, seColumn)
	}
	if m.Visible != sharedTypes.ST_TrueFalseBlankUnset {
		seVisible := xml.StartElement{Name: xml.Name{Local: "x:Visible"}}
		e.EncodeElement(m.Visible, seVisible)
	}
	if m.RowHidden != sharedTypes.ST_TrueFalseBlankUnset {
		seRowHidden := xml.StartElement{Name: xml.Name{Local: "x:RowHidden"}}
		e.EncodeElement(m.RowHidden, seRowHidden)
	}
	if m.ColHidden != sharedTypes.ST_TrueFalseBlankUnset {
		seColHidden := xml.StartElement{Name: xml.Name{Local: "x:ColHidden"}}
		e.EncodeElement(m.ColHidden, seColHidden)
	}
	if m.VTEdit != nil {
		seVTEdit := xml.StartElement{Name: xml.Name{Local: "x:VTEdit"}}
		e.EncodeElement(m.VTEdit, seVTEdit)
	}
	if m.MultiLine != sharedTypes.ST_TrueFalseBlankUnset {
		seMultiLine := xml.StartElement{Name: xml.Name{Local: "x:MultiLine"}}
		e.EncodeElement(m.MultiLine, seMultiLine)
	}
	if m.VScroll != sharedTypes.ST_TrueFalseBlankUnset {
		seVScroll := xml.StartElement{Name: xml.Name{Local: "x:VScroll"}}
		e.EncodeElement(m.VScroll, seVScroll)
	}
	if m.ValidIds != sharedTypes.ST_TrueFalseBlankUnset {
		seValidIds := xml.StartElement{Name: xml.Name{Local: "x:ValidIds"}}
		e.EncodeElement(m.ValidIds, seValidIds)
	}
	if m.FmlaRange != nil {
		seFmlaRange := xml.StartElement{Name: xml.Name{Local: "x:FmlaRange"}}
		goffice.AddPreserveSpaceAttr(&seFmlaRange, *m.FmlaRange)
		e.EncodeElement(m.FmlaRange, seFmlaRange)
	}
	if m.WidthMin != nil {
		seWidthMin := xml.StartElement{Name: xml.Name{Local: "x:WidthMin"}}
		e.EncodeElement(m.WidthMin, seWidthMin)
	}
	if m.Sel != nil {
		seSel := xml.StartElement{Name: xml.Name{Local: "x:Sel"}}
		e.EncodeElement(m.Sel, seSel)
	}
	if m.NoThreeD2 != sharedTypes.ST_TrueFalseBlankUnset {
		seNoThreeD2 := xml.StartElement{Name: xml.Name{Local: "x:NoThreeD2"}}
		e.EncodeElement(m.NoThreeD2, seNoThreeD2)
	}
	if m.SelType != nil {
		seSelType := xml.StartElement{Name: xml.Name{Local: "x:SelType"}}
		goffice.AddPreserveSpaceAttr(&seSelType, *m.SelType)
		e.EncodeElement(m.SelType, seSelType)
	}
	if m.MultiSel != nil {
		seMultiSel := xml.StartElement{Name: xml.Name{Local: "x:MultiSel"}}
		goffice.AddPreserveSpaceAttr(&seMultiSel, *m.MultiSel)
		e.EncodeElement(m.MultiSel, seMultiSel)
	}
	if m.LCT != nil {
		seLCT := xml.StartElement{Name: xml.Name{Local: "x:LCT"}}
		goffice.AddPreserveSpaceAttr(&seLCT, *m.LCT)
		e.EncodeElement(m.LCT, seLCT)
	}
	if m.ListItem != nil {
		seListItem := xml.StartElement{Name: xml.Name{Local: "x:ListItem"}}
		goffice.AddPreserveSpaceAttr(&seListItem, *m.ListItem)
		e.EncodeElement(m.ListItem, seListItem)
	}
	if m.DropStyle != nil {
		seDropStyle := xml.StartElement{Name: xml.Name{Local: "x:DropStyle"}}
		goffice.AddPreserveSpaceAttr(&seDropStyle, *m.DropStyle)
		e.EncodeElement(m.DropStyle, seDropStyle)
	}
	if m.Colored != sharedTypes.ST_TrueFalseBlankUnset {
		seColored := xml.StartElement{Name: xml.Name{Local: "x:Colored"}}
		e.EncodeElement(m.Colored, seColored)
	}
	if m.DropLines != nil {
		seDropLines := xml.StartElement{Name: xml.Name{Local: "x:DropLines"}}
		e.EncodeElement(m.DropLines, seDropLines)
	}
	if m.Checked != nil {
		seChecked := xml.StartElement{Name: xml.Name{Local: "x:Checked"}}
		e.EncodeElement(m.Checked, seChecked)
	}
	if m.FmlaLink != nil {
		seFmlaLink := xml.StartElement{Name: xml.Name{Local: "x:FmlaLink"}}
		goffice.AddPreserveSpaceAttr(&seFmlaLink, *m.FmlaLink)
		e.EncodeElement(m.FmlaLink, seFmlaLink)
	}
	if m.FmlaPict != nil {
		seFmlaPict := xml.StartElement{Name: xml.Name{Local: "x:FmlaPict"}}
		goffice.AddPreserveSpaceAttr(&seFmlaPict, *m.FmlaPict)
		e.EncodeElement(m.FmlaPict, seFmlaPict)
	}
	if m.NoThreeD != sharedTypes.ST_TrueFalseBlankUnset {
		seNoThreeD := xml.StartElement{Name: xml.Name{Local: "x:NoThreeD"}}
		e.EncodeElement(m.NoThreeD, seNoThreeD)
	}
	if m.FirstButton != sharedTypes.ST_TrueFalseBlankUnset {
		seFirstButton := xml.StartElement{Name: xml.Name{Local: "x:FirstButton"}}
		e.EncodeElement(m.FirstButton, seFirstButton)
	}
	if m.FmlaGroup != nil {
		seFmlaGroup := xml.StartElement{Name: xml.Name{Local: "x:FmlaGroup"}}
		goffice.AddPreserveSpaceAttr(&seFmlaGroup, *m.FmlaGroup)
		e.EncodeElement(m.FmlaGroup, seFmlaGroup)
	}
	if m.Val != nil {
		seVal := xml.StartElement{Name: xml.Name{Local: "x:Val"}}
		e.EncodeElement(m.Val, seVal)
	}
	if m.Min != nil {
		seMin := xml.StartElement{Name: xml.Name{Local: "x:Min"}}
		e.EncodeElement(m.Min, seMin)
	}
	if m.Max != nil {
		seMax := xml.StartElement{Name: xml.Name{Local: "x:Max"}}
		e.EncodeElement(m.Max, seMax)
	}
	if m.Inc != nil {
		seInc := xml.StartElement{Name: xml.Name{Local: "x:Inc"}}
		e.EncodeElement(m.Inc, seInc)
	}
	if m.Page != nil {
		sePage := xml.StartElement{Name: xml.Name{Local: "x:Page"}}
		e.EncodeElement(m.Page, sePage)
	}
	if m.Horiz != sharedTypes.ST_TrueFalseBlankUnset {
		seHoriz := xml.StartElement{Name: xml.Name{Local: "x:Horiz"}}
		e.EncodeElement(m.Horiz, seHoriz)
	}
	if m.Dx != nil {
		seDx := xml.StartElement{Name: xml.Name{Local: "x:Dx"}}
		e.EncodeElement(m.Dx, seDx)
	}
	if m.MapOCX != sharedTypes.ST_TrueFalseBlankUnset {
		seMapOCX := xml.StartElement{Name: xml.Name{Local: "x:MapOCX"}}
		e.EncodeElement(m.MapOCX, seMapOCX)
	}
	if m.CF != nil {
		seCF := xml.StartElement{Name: xml.Name{Local: "x:CF"}}
		for _, c := range m.CF {
			e.EncodeElement(c, seCF)
		}
	}
	if m.Camera != sharedTypes.ST_TrueFalseBlankUnset {
		seCamera := xml.StartElement{Name: xml.Name{Local: "x:Camera"}}
		e.EncodeElement(m.Camera, seCamera)
	}
	if m.RecalcAlways != sharedTypes.ST_TrueFalseBlankUnset {
		seRecalcAlways := xml.StartElement{Name: xml.Name{Local: "x:RecalcAlways"}}
		e.EncodeElement(m.RecalcAlways, seRecalcAlways)
	}
	if m.AutoScale != sharedTypes.ST_TrueFalseBlankUnset {
		seAutoScale := xml.StartElement{Name: xml.Name{Local: "x:AutoScale"}}
		e.EncodeElement(m.AutoScale, seAutoScale)
	}
	if m.DDE != sharedTypes.ST_TrueFalseBlankUnset {
		seDDE := xml.StartElement{Name: xml.Name{Local: "x:DDE"}}
		e.EncodeElement(m.DDE, seDDE)
	}
	if m.UIObj != sharedTypes.ST_TrueFalseBlankUnset {
		seUIObj := xml.StartElement{Name: xml.Name{Local: "x:UIObj"}}
		e.EncodeElement(m.UIObj, seUIObj)
	}
	if m.ScriptText != nil {
		seScriptText := xml.StartElement{Name: xml.Name{Local: "x:ScriptText"}}
		goffice.AddPreserveSpaceAttr(&seScriptText, *m.ScriptText)
		e.EncodeElement(m.ScriptText, seScriptText)
	}
	if m.ScriptExtended != nil {
		seScriptExtended := xml.StartElement{Name: xml.Name{Local: "x:ScriptExtended"}}
		goffice.AddPreserveSpaceAttr(&seScriptExtended, *m.ScriptExtended)
		e.EncodeElement(m.ScriptExtended, seScriptExtended)
	}
	if m.ScriptLanguage != nil {
		seScriptLanguage := xml.StartElement{Name: xml.Name{Local: "x:ScriptLanguage"}}
		e.EncodeElement(m.ScriptLanguage, seScriptLanguage)
	}
	if m.ScriptLocation != nil {
		seScriptLocation := xml.StartElement{Name: xml.Name{Local: "x:ScriptLocation"}}
		e.EncodeElement(m.ScriptLocation, seScriptLocation)
	}
	if m.FmlaTxbx != nil {
		seFmlaTxbx := xml.StartElement{Name: xml.Name{Local: "x:FmlaTxbx"}}
		goffice.AddPreserveSpaceAttr(&seFmlaTxbx, *m.FmlaTxbx)
		e.EncodeElement(m.FmlaTxbx, seFmlaTxbx)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ClientData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ObjectTypeAttr = ST_ObjectType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "ObjectType" {
			m.ObjectTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_ClientData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MoveWithCells"}:
				m.MoveWithCells = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MoveWithCells, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SizeWithCells"}:
				m.SizeWithCells = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.SizeWithCells, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Anchor"}:
				m.Anchor = new(string)
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Locked"}:
				m.Locked = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Locked, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DefaultSize"}:
				m.DefaultSize = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.DefaultSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "PrintObject"}:
				m.PrintObject = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.PrintObject, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Disabled"}:
				m.Disabled = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Disabled, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoFill"}:
				m.AutoFill = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoLine"}:
				m.AutoLine = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoLine, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoPict"}:
				m.AutoPict = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoPict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaMacro"}:
				m.FmlaMacro = new(string)
				if err := d.DecodeElement(m.FmlaMacro, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "TextHAlign"}:
				m.TextHAlign = new(string)
				if err := d.DecodeElement(m.TextHAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "TextVAlign"}:
				m.TextVAlign = new(string)
				if err := d.DecodeElement(m.TextVAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "LockText"}:
				m.LockText = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.LockText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "JustLastX"}:
				m.JustLastX = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.JustLastX, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SecretEdit"}:
				m.SecretEdit = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.SecretEdit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Default"}:
				m.Default = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Default, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Help"}:
				m.Help = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Help, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Cancel"}:
				m.Cancel = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Cancel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Dismiss"}:
				m.Dismiss = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Dismiss, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Accel"}:
				m.Accel = new(int64)
				if err := d.DecodeElement(m.Accel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Accel2"}:
				m.Accel2 = new(int64)
				if err := d.DecodeElement(m.Accel2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Row"}:
				m.Row = new(int64)
				if err := d.DecodeElement(m.Row, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Column"}:
				m.Column = new(int64)
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Visible"}:
				m.Visible = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Visible, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "RowHidden"}:
				m.RowHidden = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.RowHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ColHidden"}:
				m.ColHidden = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.ColHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "VTEdit"}:
				m.VTEdit = new(int64)
				if err := d.DecodeElement(m.VTEdit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MultiLine"}:
				m.MultiLine = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MultiLine, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "VScroll"}:
				m.VScroll = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.VScroll, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ValidIds"}:
				m.ValidIds = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.ValidIds, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaRange"}:
				m.FmlaRange = new(string)
				if err := d.DecodeElement(m.FmlaRange, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "WidthMin"}:
				m.WidthMin = new(int64)
				if err := d.DecodeElement(m.WidthMin, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Sel"}:
				m.Sel = new(int64)
				if err := d.DecodeElement(m.Sel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "NoThreeD2"}:
				m.NoThreeD2 = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.NoThreeD2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "SelType"}:
				m.SelType = new(string)
				if err := d.DecodeElement(m.SelType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MultiSel"}:
				m.MultiSel = new(string)
				if err := d.DecodeElement(m.MultiSel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "LCT"}:
				m.LCT = new(string)
				if err := d.DecodeElement(m.LCT, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ListItem"}:
				m.ListItem = new(string)
				if err := d.DecodeElement(m.ListItem, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DropStyle"}:
				m.DropStyle = new(string)
				if err := d.DecodeElement(m.DropStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Colored"}:
				m.Colored = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Colored, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DropLines"}:
				m.DropLines = new(int64)
				if err := d.DecodeElement(m.DropLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Checked"}:
				m.Checked = new(int64)
				if err := d.DecodeElement(m.Checked, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaLink"}:
				m.FmlaLink = new(string)
				if err := d.DecodeElement(m.FmlaLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaPict"}:
				m.FmlaPict = new(string)
				if err := d.DecodeElement(m.FmlaPict, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "NoThreeD"}:
				m.NoThreeD = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.NoThreeD, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FirstButton"}:
				m.FirstButton = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.FirstButton, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaGroup"}:
				m.FmlaGroup = new(string)
				if err := d.DecodeElement(m.FmlaGroup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Val"}:
				m.Val = new(int64)
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Min"}:
				m.Min = new(int64)
				if err := d.DecodeElement(m.Min, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Max"}:
				m.Max = new(int64)
				if err := d.DecodeElement(m.Max, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Inc"}:
				m.Inc = new(int64)
				if err := d.DecodeElement(m.Inc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Page"}:
				m.Page = new(int64)
				if err := d.DecodeElement(m.Page, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Horiz"}:
				m.Horiz = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Horiz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Dx"}:
				m.Dx = new(int64)
				if err := d.DecodeElement(m.Dx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "MapOCX"}:
				m.MapOCX = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.MapOCX, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "CF"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.CF = append(m.CF, tmp)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "Camera"}:
				m.Camera = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.Camera, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "RecalcAlways"}:
				m.RecalcAlways = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.RecalcAlways, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "AutoScale"}:
				m.AutoScale = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.AutoScale, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "DDE"}:
				m.DDE = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.DDE, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "UIObj"}:
				m.UIObj = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.UIObj, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptText"}:
				m.ScriptText = new(string)
				if err := d.DecodeElement(m.ScriptText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptExtended"}:
				m.ScriptExtended = new(string)
				if err := d.DecodeElement(m.ScriptExtended, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptLanguage"}:
				m.ScriptLanguage = new(uint32)
				if err := d.DecodeElement(m.ScriptLanguage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ScriptLocation"}:
				m.ScriptLocation = new(uint32)
				if err := d.DecodeElement(m.ScriptLocation, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "FmlaTxbx"}:
				m.FmlaTxbx = new(string)
				if err := d.DecodeElement(m.FmlaTxbx, &el); err != nil {
					return err
				}
			default:
				goffice.Log("skipping unsupported element on CT_ClientData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ClientData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ClientData and its children
func (m *CT_ClientData) Validate() error {
	return m.ValidateWithPath("CT_ClientData")
}

// ValidateWithPath validates the CT_ClientData and its children, prefixing error messages with path
func (m *CT_ClientData) ValidateWithPath(path string) error {
	if m.ObjectTypeAttr == ST_ObjectTypeUnset {
		return fmt.Errorf("%s/ObjectTypeAttr is a mandatory field", path)
	}
	if err := m.ObjectTypeAttr.ValidateWithPath(path + "/ObjectTypeAttr"); err != nil {
		return err
	}
	if err := m.MoveWithCells.ValidateWithPath(path + "/MoveWithCells"); err != nil {
		return err
	}
	if err := m.SizeWithCells.ValidateWithPath(path + "/SizeWithCells"); err != nil {
		return err
	}
	if err := m.Locked.ValidateWithPath(path + "/Locked"); err != nil {
		return err
	}
	if err := m.DefaultSize.ValidateWithPath(path + "/DefaultSize"); err != nil {
		return err
	}
	if err := m.PrintObject.ValidateWithPath(path + "/PrintObject"); err != nil {
		return err
	}
	if err := m.Disabled.ValidateWithPath(path + "/Disabled"); err != nil {
		return err
	}
	if err := m.AutoFill.ValidateWithPath(path + "/AutoFill"); err != nil {
		return err
	}
	if err := m.AutoLine.ValidateWithPath(path + "/AutoLine"); err != nil {
		return err
	}
	if err := m.AutoPict.ValidateWithPath(path + "/AutoPict"); err != nil {
		return err
	}
	if err := m.LockText.ValidateWithPath(path + "/LockText"); err != nil {
		return err
	}
	if err := m.JustLastX.ValidateWithPath(path + "/JustLastX"); err != nil {
		return err
	}
	if err := m.SecretEdit.ValidateWithPath(path + "/SecretEdit"); err != nil {
		return err
	}
	if err := m.Default.ValidateWithPath(path + "/Default"); err != nil {
		return err
	}
	if err := m.Help.ValidateWithPath(path + "/Help"); err != nil {
		return err
	}
	if err := m.Cancel.ValidateWithPath(path + "/Cancel"); err != nil {
		return err
	}
	if err := m.Dismiss.ValidateWithPath(path + "/Dismiss"); err != nil {
		return err
	}
	if err := m.Visible.ValidateWithPath(path + "/Visible"); err != nil {
		return err
	}
	if err := m.RowHidden.ValidateWithPath(path + "/RowHidden"); err != nil {
		return err
	}
	if err := m.ColHidden.ValidateWithPath(path + "/ColHidden"); err != nil {
		return err
	}
	if err := m.MultiLine.ValidateWithPath(path + "/MultiLine"); err != nil {
		return err
	}
	if err := m.VScroll.ValidateWithPath(path + "/VScroll"); err != nil {
		return err
	}
	if err := m.ValidIds.ValidateWithPath(path + "/ValidIds"); err != nil {
		return err
	}
	if err := m.NoThreeD2.ValidateWithPath(path + "/NoThreeD2"); err != nil {
		return err
	}
	if err := m.Colored.ValidateWithPath(path + "/Colored"); err != nil {
		return err
	}
	if err := m.NoThreeD.ValidateWithPath(path + "/NoThreeD"); err != nil {
		return err
	}
	if err := m.FirstButton.ValidateWithPath(path + "/FirstButton"); err != nil {
		return err
	}
	if err := m.Horiz.ValidateWithPath(path + "/Horiz"); err != nil {
		return err
	}
	if err := m.MapOCX.ValidateWithPath(path + "/MapOCX"); err != nil {
		return err
	}
	if err := m.Camera.ValidateWithPath(path + "/Camera"); err != nil {
		return err
	}
	if err := m.RecalcAlways.ValidateWithPath(path + "/RecalcAlways"); err != nil {
		return err
	}
	if err := m.AutoScale.ValidateWithPath(path + "/AutoScale"); err != nil {
		return err
	}
	if err := m.DDE.ValidateWithPath(path + "/DDE"); err != nil {
		return err
	}
	if err := m.UIObj.ValidateWithPath(path + "/UIObj"); err != nil {
		return err
	}
	return nil
}
