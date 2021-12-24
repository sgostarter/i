package l

type FieldType int

const (
	FieldTypeAny     FieldType = 0
	FieldTypeError   FieldType = 1
	FieldTypeString  FieldType = 2
	FieldTypeInt     FieldType = 3
	FieldTypeInt64   FieldType = 4
	FieldTypeUint    FieldType = 5
	FieldTypeUint64  FieldType = 6
	FieldTypeFloat32 FieldType = 7
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

func AnyField(k string, v interface{}) Field {
	return FieldAny(k, v)
}

func ErrorField(err error) Field {
	return Field{
		T: FieldTypeError,
		K: "err",
		V: err,
	}
}

func StringField(key, s string) Field {
	return FieldString(key, s)
}

func newField(t FieldType, k string, v interface{}) Field {
	return Field{
		T: t,
		K: k,
		V: v,
	}
}

func IntField(k string, v int) Field {
	return newField(FieldTypeInt, k, v)
}

func Int64Field(k string, v int64) Field {
	return newField(FieldTypeInt64, k, v)
}

func UIntField(k string, v string) Field {
	return newField(FieldTypeUint, k, v)
}

func UInt64Field(k string, v uint64) Field {
	return newField(FieldTypeUint64, k, v)
}

func FloatField32(k string, v float32) Field {
	return newField(FieldTypeFloat32, k, v)
}

func Float64Field(k string, v float64) Field {
	return newField(FieldTypeFloat64, k, v)
}
