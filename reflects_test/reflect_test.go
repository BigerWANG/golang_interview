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
