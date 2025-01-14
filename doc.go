/*
Package goffice provides creation, reading, and writing of ECMA 376 Office Open
XML documents, spreadsheets and presentations.  It is still early in
development, but is progressing quickly.  This library takes a slightly
different approach from others, in that it starts by trying to support all of
the ECMA-376 standard when marshaling/unmarshaling XML documents.  From there it
adds wrappers around the ECMA-376 derived types that provide a more convenient
interface.

The raw XML based types reside in the `schema/“ directory. These types are
always accessible from the wrapper types via a `X() method that returns the
raw type.  Except for the base documents (document.Document,
spreadsheet.Workbook and presentation.Presentation), the other wrapper types are
value types with non-pointer methods.  They exist solely to modify and return
data from one or more XML types.

The packages of interest are github.com/dhx007/goffice/document,
unidoc/goffice/spreadsheet and github.com/dhx007/goffice/presentation.
*/
package goffice // import "github.com/dhx007/goffice"
