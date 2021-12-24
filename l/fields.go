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

func IntField(key, s string) logger.Field {
	return newField(logger.FieldTypeInt, key, s)
}

func Int64Field(key, s string) logger.Field {
	return newField(logger.FieldTypeInt64, key, s)
}

func UIntField(key, s string) logger.Field {
	return newField(logger.FieldTypeUint, key, s)
}

func UInt64Field(key, s string) logger.Field {
	return newField(logger.FieldTypeUint64, key, s)
}

func FloatField(key, s string) logger.Field {
	return newField(logger.FieldTypeFloat, key, s)
}

func Float64Field(key, s string) logger.Field {
	return newField(logger.FieldTypeFloat64, key, s)
}
