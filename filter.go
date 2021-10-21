// 切片过滤
package array

import (
	"reflect"
)

// Filter
// 将传入的 arr 使用 filter 进行过滤
// filter 的函数签名需要为 func(e Element) bool, 其中 Element 为 arr的元素类型
func Filter(arr interface{}, filter interface{}) interface{} {
	av := reflect.ValueOf(arr)
	at := av.Type()
	if at.Kind() != reflect.Slice {
		panic("arr must be slice:" + at.Kind().String())
	}
	fv := reflect.ValueOf(filter)
	ft := fv.Type()
	if ft.Kind() != reflect.Func {
		panic("filter must be func")
	}
	hold := make([]reflect.Value, 0, av.Cap())
	for i := 0; i < av.Len(); i++ {
		rst := fv.Call([]reflect.Value{av.Index(i)})
		if len(rst) != 1 {
			panic("filter must return bool")
		}
		if rst[0].Kind() != reflect.Bool {
			panic("filter must return bool")
		}
		v, _ := rst[0].Interface().(bool)
		if v {
			hold = append(hold, av.Index(i))
		}
	}
	rtn := reflect.MakeSlice(at, 0, len(hold))
	return reflect.Append(rtn, hold...).Interface()
}
