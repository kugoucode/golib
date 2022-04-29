package gostruct

import "reflect"

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt64(name string, value int64) {
	i.Field(name).SetInt(value)
}

func (i *Instance) SetFloat64(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) Interface() interface{} {
	return i.internal.Interface()
}

// Addr 方法,可用于取地址,常用于调用指针接收者的方法
func (i *Instance) Addr() interface{} {
	return i.internal.Addr().Interface()
}
