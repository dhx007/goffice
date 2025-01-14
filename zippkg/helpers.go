package zippkg

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/dhx007/goffice"

	"github.com/dhx007/goffice/algo"
	"github.com/dhx007/goffice/schema/soo/pkg/relationships"
)

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	sp := strings.Split(path, "/")
	pathPortion := strings.Join(sp[0:len(sp)-1], "/")
	filePortion := sp[len(sp)-1]
	pathPortion += "/_rels/"
	filePortion += ".rels"
	return pathPortion + filePortion
}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *zip.File, dest interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("error reading %s: %s", f.Name, err)
	}
	defer rc.Close()
	dec := xml.NewDecoder(rc)
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("error decoding %s: %s", f.Name, err)
	}

	// this ensures that relationship ID is increasing, which we apparently rely
	// on....
	if ds, ok := dest.(*relationships.Relationships); ok {
		for i, r := range ds.Relationship {
			switch r.TypeAttr {
			// Common
			case goffice.OfficeDocumentTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.OfficeDocumentType
			case goffice.StylesTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.StylesType
			case goffice.ThemeTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.ThemeType
			case goffice.SettingsTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.SettingsType
			case goffice.ImageTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.ImageType
			case goffice.CommentsTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.CommentsType
			case goffice.ThumbnailTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.ThumbnailType
			case goffice.DrawingTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.DrawingType
			case goffice.ChartTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.ChartType
			case goffice.ExtendedPropertiesTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.ExtendedPropertiesType
			case goffice.CustomXMLTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.CustomXMLType

			// SML
			case goffice.WorksheetTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.WorksheetType
			case goffice.SharedStringsTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.SharedStringsType
			case goffice.TableTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.TableType

			// WML
			case goffice.HeaderTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.HeaderType
			case goffice.FooterTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.FooterType
			case goffice.NumberingTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.NumberingType
			case goffice.FontTableTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.FontTableType
			case goffice.WebSettingsTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.WebSettingsType
			case goffice.FootNotesTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.FootNotesType
			case goffice.EndNotesTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.EndNotesType

			// PML
			case goffice.SlideTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.SlideType

			// VML
			case goffice.VMLDrawingTypeStrict:
				ds.Relationship[i].TypeAttr = goffice.VMLDrawingType
			}
		}

		sort.Slice(ds.Relationship, func(i, j int) bool {
			lhs := ds.Relationship[i]
			rhs := ds.Relationship[j]
			return algo.NaturalLess(lhs.IdAttr, rhs.IdAttr)
		})
	}
	return nil
}

// AddFileFromDisk reads a file from disk and adds it at a given path to a zip file.
func AddFileFromDisk(z *zip.Writer, zipPath, diskPath string) error {
	w, err := z.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", zipPath, err)
	}
	f, err := os.Open(diskPath)
	if err != nil {
		return fmt.Errorf("error opening %s: %s", diskPath, err)
	}
	_, err = io.Copy(w, f)
	return err
}

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes(z *zip.Writer, zipPath string, data []byte) error {
	w, err := z.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", zipPath, err)
	}
	_, err = io.Copy(w, bytes.NewReader(data))
	return err
}

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *zip.File, path string) (string, error) {
	tmpFile, err := ioutil.TempFile(path, "zz")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()
	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()
	_, err = io.Copy(tmpFile, rc)
	if err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}
