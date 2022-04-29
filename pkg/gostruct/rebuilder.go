package gostruct

import (
	"math"
	"math/cmplx"
	"reflect"
	"unsafe"
)

// ParserStruct2Map 解析结构体为切片Map
func ParserStruct2Map(t reflect.Type) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tmp := make(map[string]interface{}, 0)
		tmp["name"] = f.Name
		tmp["type"] = f.Type.Kind().String()
		tmp["pkg_path"] = f.PkgPath
		tmp["index"] = f.Index
		tmp["offset"] = f.Offset
		tmp["tag"] = f.Tag
		tmp["anonymous"] = f.Anonymous

		if f.Type.Kind() == reflect.Struct {
			tmp["childStruct"] = ParserStruct2Map(f.Type)
		}

		ret = append(ret, tmp)
	}
	return ret
}

// ReBuildStruct 重建结构体
func ReBuildStruct(items []map[string]interface{}) (addr interface{}, structs interface{}) {
	if len(items) == 0 {
		return nil, nil
	}

	builder := New()
	for _, item := range items {
		f := BuildField(item)
		if f == nil {
			continue
		}
		builder.AddField(*f)
	}

	tStruct := builder.Build()

	addr = tStruct.New().Addr()
	structs = tStruct.New().Interface()

	return addr, structs
}

// BuildField 构造struct单字段
func BuildField(item map[string]interface{}) *Field {
	name, nhas := item["name"]
	typ, thas := item["type"]
	if !nhas || !thas {
		return &Field{}
	}

	var field *Field
	if typ.(string) == "struct" {
		_, tmp := ReBuildStruct(item["childStruct"].([]map[string]interface{}))
		field = NewField(name.(string), reflect.TypeOf(tmp))
	} else {
		field = NewField(name.(string), TypeName2ReflectType(typ.(string)))
	}

	if tag, has := item["tag"]; has {
		field = field.SetTag(tag.(reflect.StructTag))
	}

	if offset, has := item["offset"]; has {
		field = field.SetOffset(offset.(uintptr))
	}

	if pkg_path, has := item["pkg_path"]; has {
		field = field.SetPkgPath(pkg_path.(string))
	}

	if index, has := item["index"]; has {
		field = field.SetIndex(index.([]int))
	}

	if anonymous, has := item["anonymous"]; has {
		field = field.SetAnonymous(anonymous.(bool))
	}

	return field
}

// GetReflectTypeByValue 通过类型名获取对应的反射类型
func TypeName2ReflectType(typeVal string) reflect.Type {
	var ret reflect.Type
	switch typeVal {
	case "bool":
		ret = reflect.TypeOf(true)
	case "int":
		ret = reflect.TypeOf(int(0))
	case "int8":
		ret = reflect.TypeOf(int8(0))
	case "int16":
		ret = reflect.TypeOf(int16(0))
	case "int32":
		ret = reflect.TypeOf(int32(0))
	case "int64":
		ret = reflect.TypeOf(int64(0))
	case "uint":
		ret = reflect.TypeOf(int(0))
	case "uint8":
		ret = reflect.TypeOf(int8(0))
	case "uint16":
		ret = reflect.TypeOf(int16(0))
	case "uint32":
		ret = reflect.TypeOf(int32(0))
	case "uint64":
		ret = reflect.TypeOf(int64(0))
	case "float32":
		ret = reflect.TypeOf(float32(0))
	case "float64":
		ret = reflect.TypeOf(float64(0))
	case "slice":
		ret = reflect.TypeOf([]string{})
	case "map":
		ret = reflect.TypeOf(map[string]interface{}{})
	case "func":
		ret = reflect.TypeOf(func() {})
	case "chan":
		ret = reflect.TypeOf(make(chan int))
	case "ptr":
		p := int(0)
		ret = reflect.TypeOf(&p)
	case "array":
		ret = reflect.TypeOf([3]string{})
	// case "interface":
	// 	var any interface{}
	// 	ret = reflect.TypeOf(any).Key()
	case "uintptr":
		ret = reflect.TypeOf(uintptr(unsafe.Pointer(new(struct{}))))
	case "complex64":
		ret = reflect.TypeOf(math.Pi)
	case "complex128":
		ret = reflect.TypeOf(cmplx.Exp(math.Pi))
	case "unsafe.Pointer", "unsafe_pointer", "pointer": // 这里不确定具体输出值
		ret = reflect.TypeOf(unsafe.Pointer(new(struct{})))
	default:
		ret = reflect.TypeOf("")
	}

	return ret
}
