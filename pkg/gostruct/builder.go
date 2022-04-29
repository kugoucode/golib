package gostruct

import "reflect"

type Builder struct {
	field []reflect.StructField
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) AddField(f Field) *Builder {
	b.field = append(
		b.field,
		reflect.StructField(f))

	return b
}

func (b *Builder) AddString(name string) *Builder {
	return b.AddField(*NewStringField(name))
}

func (b *Builder) AddBool(name string) *Builder {
	return b.AddField(*NewBoolField(name))
}

func (b *Builder) AddInt64(name string) *Builder {
	return b.AddField(*NewInt64Field(name))
}

func (b *Builder) AddFloat64(name string) *Builder {
	return b.AddField(*NewFloat64Field(name))

}

func (b *Builder) Build() Struct {
	strct := reflect.StructOf(b.field)

	// 收集字段在结构体中的顺序位置
	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		index[strct.Field(i).Name] = i
	}

	return Struct{strct, index}
}
