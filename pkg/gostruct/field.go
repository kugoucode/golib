package gostruct

import "reflect"

// Field 定义字段类型
type Field reflect.StructField

// NewField 新增字段
// 	set field name and field type
func NewField(name string, ftype reflect.Type) *Field {
	return &Field{
		Name: name,
		Type: ftype,
	}
}

func NewStringField(name string) *Field {
	return &Field{
		Name: name,
		Type: reflect.TypeOf(""),
	}
}

func NewBoolField(name string) *Field {
	return &Field{
		Name: name,
		Type: reflect.TypeOf(true),
	}
}

func NewInt64Field(name string) *Field {
	return &Field{
		Name: name,
		Type: reflect.TypeOf(int64(0)),
	}
}

func NewFloat64Field(name string) *Field {
	return &Field{
		Name: name,
		Type: reflect.TypeOf(float64(0)),
	}
}

// SetTag 设置标签
func (f *Field) SetTag(tag reflect.StructTag) *Field {
	f.Tag = tag
	return f
}

// SetPkgPath
// PkgPath is the package path that qualifies a lower case (unexported)
// field name. It is empty for upper case (exported) field names.
// See https://golang.org/ref/spec#Uniqueness_of_identifiers
func (f *Field) SetPkgPath(pkg string) *Field {
	f.PkgPath = pkg
	return f
}

//  SetOffset offset within struct, in bytes
func (f *Field) SetOffset(offset uintptr) *Field {
	f.Offset = offset
	return f
}

// SetIndex index sequence for Type.FieldByIndex
func (f *Field) SetIndex(index []int) *Field {
	f.Index = index
	return f
}

// SetAnonymous is an embedded field
func (f *Field) SetAnonymous(t bool) *Field {
	f.Anonymous = t
	return f
}
