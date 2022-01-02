package riddles

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "order slice",
			args: args{
				arr:  []int{3, 2, 1},
				low:  0,
				high: 2,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := partition(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
