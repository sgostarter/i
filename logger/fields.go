package logger

type FieldType int

const (
	FieldTypeError  FieldType = 1
	FieldTypeString FieldType = 2
)

type Field struct {
	T FieldType
	K string
	V interface{}
}

func FieldError(key string, err error) Field {
	return Field{
		T: FieldTypeError,
		K: key,
		V: err,
	}
}

func FieldString(key, s string) Field {
	return Field{
		T: FieldTypeString,
		K: key,
		V: s,
	}
}
