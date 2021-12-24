package logger

type FieldType int

const (
	FieldTypeAny     FieldType = 0
	FieldTypeError   FieldType = 1
	FieldTypeString  FieldType = 2
	FieldTypeInt     FieldType = 3
	FieldTypeInt64   FieldType = 4
	FieldTypeUint    FieldType = 5
	FieldTypeUint64  FieldType = 6
	FieldTypeFloat   FieldType = 7
	FieldTypeFloat64 FieldType = 8
)

type Field struct {
	T FieldType
	K string
	V interface{}
}

func FieldAny(k string, v interface{}) Field {
	return Field{
		T: FieldTypeAny,
		K: k,
		V: v,
	}
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
