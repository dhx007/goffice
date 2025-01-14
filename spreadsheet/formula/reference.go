package formula

// ReferenceType is a type of reference
//
//go:generate stringer -type=ReferenceType
type ReferenceType byte

const (
	ReferenceTypeInvalid ReferenceType = iota
	ReferenceTypeCell
	ReferenceTypeHorizontalRange
	ReferenceTypeVerticalRange
	ReferenceTypeNamedRange
	ReferenceTypeRange
	ReferenceTypeSheet
)

type Reference struct {
	Type  ReferenceType
	Value string
}

var ReferenceInvalid = Reference{Type: ReferenceTypeInvalid}

func MakeRangeReference(ref string) Reference {
	return Reference{Type: ReferenceTypeRange, Value: ref}
}
