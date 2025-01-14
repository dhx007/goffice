package spreadsheet

import (
	"runtime"

	"github.com/dhx007/goffice"
	"github.com/dhx007/goffice/common"
	"github.com/dhx007/goffice/schema/soo/sml"
)

// New constructs a new workbook.
func New() *Workbook {
	wb := &Workbook{}
	wb.x = sml.NewWorkbook()

	runtime.SetFinalizer(wb, workbookFinalizer)

	wb.AppProperties = common.NewAppProperties()
	wb.CoreProperties = common.NewCoreProperties()
	wb.StyleSheet = NewStyleSheet(wb)

	wb.Rels = common.NewRelationships()
	wb.wbRels = common.NewRelationships()

	wb.Rels.AddRelationship(goffice.RelativeFilename(goffice.DocTypeSpreadsheet, "", goffice.ExtendedPropertiesType, 0), goffice.ExtendedPropertiesType)
	wb.Rels.AddRelationship(goffice.RelativeFilename(goffice.DocTypeSpreadsheet, "", goffice.CorePropertiesType, 0), goffice.CorePropertiesType)
	wb.Rels.AddRelationship(goffice.RelativeFilename(goffice.DocTypeSpreadsheet, "", goffice.OfficeDocumentType, 0), goffice.OfficeDocumentType)
	wb.wbRels.AddRelationship(goffice.RelativeFilename(goffice.DocTypeSpreadsheet, goffice.OfficeDocumentType, goffice.StylesType, 0), goffice.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddDefault("vml", goffice.VMLDrawingContentType)
	wb.ContentTypes.AddOverride(goffice.AbsoluteFilename(goffice.DocTypeSpreadsheet, goffice.OfficeDocumentType, 0), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride(goffice.AbsoluteFilename(goffice.DocTypeSpreadsheet, goffice.StylesType, 0), goffice.SMLStyleSheetContentType)

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride(goffice.AbsoluteFilename(goffice.DocTypeSpreadsheet, goffice.SharedStringsType, 0), goffice.SharedStringsContentType)
	wb.wbRels.AddRelationship(goffice.RelativeFilename(goffice.DocTypeSpreadsheet, goffice.OfficeDocumentType, goffice.SharedStringsType, 0), goffice.SharedStringsType)

	return wb
}
