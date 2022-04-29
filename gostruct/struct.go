package gostruct

import "reflect"

type Struct struct {
	strct reflect.Type
	index map[string]int
}

// 生成有效的结构体对象实例
func (s *Struct) New() *Instance {
	instance := reflect.New(s.strct).Elem()
	return &Instance{instance, s.index}
}
