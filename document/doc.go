/*
Package document provides creation, reading, and writing of ECMA 376 Open
Office XML documents.

Example:

	doc := document.New()
	para := doc.AddParagraph()
	run := para.AddRun()
	run.SetText("foo")
	doc.SaveToFile("foo.docx")
*/
package document
