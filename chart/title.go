package chart

import (
	"github.com/dhx007/goffice/color"
	"github.com/dhx007/goffice/drawing"
	"github.com/dhx007/goffice/measurement"
	"github.com/dhx007/goffice/schema/soo/dml"
	crt "github.com/dhx007/goffice/schema/soo/dml/chart"
)

type Title struct {
	x *crt.CT_Title
}

func MakeTitle(x *crt.CT_Title) Title {
	return Title{x}
}

// X returns the inner wrapped XML type.
func (t Title) X() *crt.CT_Title {
	return t.x
}

func (t Title) InitializeDefaults() {
	t.SetText("Title")
	t.RunProperties().SetSize(16 * measurement.Point)
	t.RunProperties().SetSolidFill(color.Black)
	t.RunProperties().SetFont("Calib ri")
	t.RunProperties().SetBold(false)
}

func (t Title) SetText(s string) {
	if t.x.Tx == nil {
		t.x.Tx = crt.NewCT_Tx()
	}
	if t.x.Tx.Choice.Rich == nil {
		t.x.Tx.Choice.Rich = dml.NewCT_TextBody()
	}
	var pr *dml.CT_TextParagraph
	if len(t.x.Tx.Choice.Rich.P) == 0 {
		pr = dml.NewCT_TextParagraph()
		t.x.Tx.Choice.Rich.P = []*dml.CT_TextParagraph{pr}
	} else {
		pr = t.x.Tx.Choice.Rich.P[0]
	}

	var tr *dml.EG_TextRun
	if len(pr.EG_TextRun) == 0 {
		tr = dml.NewEG_TextRun()
		pr.EG_TextRun = []*dml.EG_TextRun{tr}
	} else {
		tr = pr.EG_TextRun[0]
	}

	if tr.R == nil {
		tr.R = dml.NewCT_RegularTextRun()
	}
	tr.R.T = s
}

func (t Title) ParagraphProperties() drawing.ParagraphProperties {
	if t.x.Tx == nil {
		t.SetText("")
	}
	if t.x.Tx.Choice.Rich.P[0].PPr == nil {
		t.x.Tx.Choice.Rich.P[0].PPr = dml.NewCT_TextParagraphProperties()
	}
	return drawing.MakeParagraphProperties(t.x.Tx.Choice.Rich.P[0].PPr)
}

func (t Title) RunProperties() drawing.RunProperties {
	if t.x.Tx == nil {
		t.SetText("")
	}
	if t.x.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr == nil {
		t.x.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr = dml.NewCT_TextCharacterProperties()
	}
	return drawing.MakeRunProperties(t.x.Tx.Choice.Rich.P[0].EG_TextRun[0].R.RPr)
}
