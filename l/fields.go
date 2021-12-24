package l

import "github.com/sgostarter/i/logger"

func AnyField(k string, v interface{}) logger.Field {
	return logger.FieldAny(k, v)
}

func ErrorField(err error) logger.Field {
	return logger.Field{
		T: logger.FieldTypeError,
		K: "err",
		V: err,
	}
}

func StringField(key, s string) logger.Field {
	return logger.FieldString(key, s)
}

func newField(t logger.FieldType, k string, v interface{}) logger.Field {
	return logger.Field{
		T: t,
		K: k,
		V: v,
	}
}

func IntField(k string, v int) logger.Field {
	return newField(logger.FieldTypeInt, k, v)
}

func Int64Field(k string, v int64) logger.Field {
	return newField(logger.FieldTypeInt64, k, v)
}

func UIntField(k string, v string) logger.Field {
	return newField(logger.FieldTypeUint, k, v)
}

func UInt64Field(k string, v uint64) logger.Field {
	return newField(logger.FieldTypeUint64, k, v)
}

func FloatField32(k string, v float32) logger.Field {
	return newField(logger.FieldTypeFloat32, k, v)
}

func Float64Field(k string, v float64) logger.Field {
	return newField(logger.FieldTypeFloat64, k, v)
}
