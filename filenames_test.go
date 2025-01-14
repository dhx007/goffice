package goffice_test

import (
	"testing"

	"github.com/dhx007/goffice"
)

func TestWMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, goffice.CorePropertiesType, "docProps/core.xml"},
		{0, goffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, goffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, goffice.StylesType, "word/styles.xml"},

		{0, goffice.OfficeDocumentType, "word/document.xml"},
		{0, goffice.FontTableType, "word/fontTable.xml"},
		{0, goffice.EndNotesType, "word/endnotes.xml"},
		{0, goffice.FootNotesType, "word/footnotes.xml"},
		{0, goffice.NumberingType, "word/numbering.xml"},
		{0, goffice.WebSettingsType, "word/webSettings.xml"},
		{0, goffice.SettingsType, "word/settings.xml"},
		{23, goffice.HeaderType, "word/header23.xml"},
		{15, goffice.FooterType, "word/footer15.xml"},
		{1, goffice.ThemeType, "word/theme/theme1.xml"},
	}
	for _, tc := range td {
		abs := goffice.AbsoluteFilename(goffice.DocTypeDocument, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestSMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, goffice.CorePropertiesType, "docProps/core.xml"},
		{0, goffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, goffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, goffice.StylesType, "xl/styles.xml"},

		{0, goffice.OfficeDocumentType, "xl/workbook.xml"},
		{15, goffice.ChartType, "xl/charts/chart15.xml"},
		{12, goffice.DrawingType, "xl/drawings/drawing12.xml"},
		{13, goffice.TableType, "xl/tables/table13.xml"},
		{2, goffice.CommentsType, "xl/comments2.xml"},
		{15, goffice.WorksheetType, "xl/worksheets/sheet15.xml"},
		{2, goffice.VMLDrawingType, "xl/drawings/vmlDrawing2.vml"},
		{0, goffice.SharedStringsType, "xl/sharedStrings.xml"},
		{1, goffice.ThemeType, "xl/theme/theme1.xml"},
		{2, goffice.ImageType, "xl/media/image2.png"},
	}
	for _, tc := range td {
		abs := goffice.AbsoluteFilename(goffice.DocTypeSpreadsheet, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestPMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, goffice.CorePropertiesType, "docProps/core.xml"},
		{0, goffice.ExtendedPropertiesType, "docProps/app.xml"},
		{0, goffice.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, goffice.StylesType, "ppt/styles.xml"},

		{0, goffice.OfficeDocumentType, "ppt/presentation.xml"},
		{4, goffice.SlideType, "ppt/slides/slide4.xml"},
		{5, goffice.SlideLayoutType, "ppt/slideLayouts/slideLayout5.xml"},
		{6, goffice.SlideMasterType, "ppt/slideMasters/slideMaster6.xml"},
		{7, goffice.ThemeType, "ppt/theme/theme7.xml"},
	}
	for _, tc := range td {
		abs := goffice.AbsoluteFilename(goffice.DocTypePresentation, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}
