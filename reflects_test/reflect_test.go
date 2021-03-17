package reflects_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRe1(t *testing.T) {
	var num = 1.2345
	fmt.Println("value:", reflect.ValueOf(num))
	fmt.Println("value:", reflect.TypeOf(num))
}

func TestRe2(t *testing.T) {
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

func TestRef3(t *testing.T) {
	type A int
	var a A = 1
	tt := reflect.TypeOf(a)
	t.Log(tt.Name()) // Name返回真实类型 A
	t.Log(tt.Kind()) // Kind返回基础类型 int
}

func TestRef4(t *testing.T) {
	// 传入的对象区分值类型和指针类型
	b:=2
	t1 := reflect.TypeOf(b)
	t2 := reflect.TypeOf(&b) // 指针类型
	fmt.Println(t1)
	fmt.Println(t2.Elem())
	t.Log(t1 == t2.Elem())
	//t.Log(t1.Kind())

}

type student struct {
	name string
	score int
}

func TestRefStruct(t *testing.T) {
	// 遍历结构体，需要先获取结构体指针的基类型
	var s student
	s.name = "wang"
	s.score = 100
	st := reflect.TypeOf(&s)
	if st.Kind() == reflect.Ptr{ // 指针类型
		t.Log(st.Kind())
		st = st.Elem()
		t.Log(st)
	}

	t.Log(st.NumField()) // field的个数

	for i:=0; i<st.NumField(); i++{
		structField := st.Field(i)  // 按照索引取值
		t.Log(structField.Name, structField.Type, structField)
	}
}

type user struct {
	name string `from:"user_name"`
	age  int    `from:"user_age" index:"int"`
}

func TestRefStructTag(t *testing.T) {
	// 提取 struct tag
	// 一般用于ORM映射，数据格式验证
	var user1 user
	ut := reflect.TypeOf(user1)
	for i:=0; i<ut.NumField(); i++{
		structField := ut.Field(i)
		t.Log(structField.Name)
		t.Log(structField.Tag.Get("from"))
		t.Log(structField.Tag.Get("index"))
		fmt.Println()
	}
}

type animal struct {
	Name string
	age int
}

func TestRefGetVal(t *testing.T) {
	// 使用反射获取和修改对象的值
	var cat animal
	v := reflect.TypeOf(&cat).Elem()
	v.FieldByName("Name")
	
}

type member struct {

}

func (m *member)Fun1(){  // 需要注意这个方法需要大写开头，否则使用反射调用时会报错
	fmt.Printf("Im running f1\n")
}

func (m *member)Fun2(){  // 需要注意这个方法需要大写开头，否则使用反射调用时会报错
	fmt.Printf("Im running f2\n")
}

func (m *member)Fun3(){  // 需要注意这个方法需要大写开头，否则使用反射调用时会报错
	fmt.Printf("Im running f3\n")
}

func OneFuncCaller(funcName string) []reflect.Value{
	// 使用反射调用方法
	var mem member
	v := reflect.ValueOf(&mem)
	m := v.MethodByName(funcName)
	params := []reflect.Value{}
	// 使用Call函数，只需按照参数列表传参即可
	return m.Call(params)
}

func MultipleFuncCaller(funcNames []string)map[string][]reflect.Value {
	// 多个函数调用
	var mem member
	res := make(map[string][]reflect.Value)
	v := reflect.ValueOf(&mem)
	for _, funcName :=range funcNames{
		m := v.MethodByName(funcName)
		res[funcName] = m.Call([]reflect.Value{})
	}
	return res
}



func TestRefCallFunc(t *testing.T) {
	OneFuncCaller("Fun1")
	res := MultipleFuncCaller([]string{"Fun1", "Fun2", "Fun3"})
	t.Log(res)
}
