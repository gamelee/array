package array

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestFilter(t *testing.T) {
	type args struct {
		arr    interface{}
		filter interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"[]string]",
			args{
				[]string{"a", "A", "b", "B", "C", "c"},
				func(str string) bool {
					return strings.ToLower(str) == str
				},
			},
			[]string{"a", "b", "c"},
		},
		{
			"[]int",
			args{
				[]int{1, 2, 3, 4, 5, 6},
				func(num int) bool {
					return num%2 == 1
				},
			},
			[]int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if rtn:=Filter(tt.args.arr, tt.args.filter); !reflect.DeepEqual(rtn, tt.want) {
				t.Errorf("source = %v Filter() = %v, want %v", tt.args.arr, tt.args.arr, tt.want)
			}
		})
	}
}

func TestSkip(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	skip := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] % 2 == 1{
			if skip > 0 {
				arr[i-skip] = arr[i]
			}
		} else {
			skip ++
		}
	}
	(*reflect.SliceHeader)(unsafe.Pointer(&arr)).Len = len(arr) - skip
	fmt.Printf("%v", arr)
}
