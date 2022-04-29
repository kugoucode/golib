// 参考：https://github.com/itsubaki/gostruct
// Go: 运行时结构体构造[reflect]： https://zhuanlan.zhihu.com/p/352808352
package main

import (
	"encoding/json"
	"fmt"
	"gostruct/gostruct"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 自定义校验struct结构体，遵循Validator配置规范
func vfunc() interface{} {
	type Class struct {
		Class string `json:"class"`
		Num   string `json:"num"`
	}

	return &struct {
		Name    string `json:"name,omitempty" validate:"required"`
		Age     string `json:"age" validate:"numeric,omitempty"`
		Classes Class  `json:"classes" validate:"omitempty"`
		// Classes struct {
		// 	Class string `json:"class"`
		// 	Num   string `json:"num"`
		// }
	}{}
}

// 解析并重建结构体，实现对已有结构体的更改
func Rebuilder() {
	person := vfunc()
	t := reflect.TypeOf(person)
	mmap := gostruct.ParserStruct2Map(t)
	// fmt.Printf("Debug-Map:%+v\n\n", mmap)

	mstruct, _ := gostruct.ReBuildStruct(mmap)

	fmt.Printf("Struct-new: [%T]\t[%+v]\n", mstruct, mstruct)
	fmt.Printf("Struct-org:[%T]\t[%+v]\n\n", person, person)

	// Output:
	// 	Struct-new: [*struct { Name string "json:\"name,omitempty\" validate:\"required\""; Age string "json:\"age\" validate:\"numeric,omitempty\""; Classes struct { Class string "json:\"class\""; Num string "json:\"num\"" } "json:\"classes\" validate:\"omitempty\"" }]      [&{Name: Age: Classes:{Class: Num:}}]
	// Struct-org:[*struct { Name string "json:\"name,omitempty\" validate:\"required\""; Age string "json:\"age\" validate:\"numeric,omitempty\""; Classes main.Class "json:\"classes\" validate:\"omitempty\"" }]        [&{Name: Age: Classes:{Class: Num:}}]

	str := `{"name":"kugouming","age":"10years","classes":{"class":"www","num":"sdsdsd"}}`
	json.Unmarshal([]byte(str), &mstruct)
	fmt.Printf("\n%+v\n\n", mstruct)
	// Output:
	// &{Name:kugouming Age:"10years" Classes:{Class:www Num:sdsdsd}}

	validate := validator.New()
	err := validate.Struct(mstruct)
	fmt.Println(err)
	// Output:
	// Key: 'Age' Error:Field validation for 'Age' failed on the 'numeric' tag

	fmt.Println("\n\n######################################################\n\n")
}

func Builder() {
	builder := gostruct.New()
	fieldName := gostruct.NewStringField("Name")
	fieldName = fieldName.SetTag(`json:"name" validate:"required"`)
	builder.AddField(*fieldName)

	fieldAge := gostruct.NewInt64Field("Age")
	fieldAge = fieldAge.SetTag(`json:"age" validate:"required"`)
	builder.AddField(*fieldAge)

	person := builder.Build()
	s := person.New().Addr()
	str := `{"name":"kugouming","age":"sdsd"}`
	json.Unmarshal([]byte(str), s)

	fmt.Printf("Type: [%T]\nStruct: [%+v]\n\n", s, s)
	// Output:
	// Type: [*struct { Name string "json:\"name\" validate:\"required\""; Age int64 "json:\"age\" validate:\"required\"" }]
	// Struct: [&{Name:kugouming Age:0}]

	validate := validator.New()
	err := validate.Struct(s)
	fmt.Println(err)
	// Out:
	// Key: 'Age' Error:Field validation for 'Age' failed on the 'required' tag

	fmt.Println("\n\n######################################################\n\n")
}

func main() {
	// 结构体重构
	Rebuilder()

	// 自定义结构体
	Builder()
}
