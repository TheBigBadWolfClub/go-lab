package riddles

import (
	"reflect"
	"testing"
)

func Test_factorial(t *testing.T) {
	type args struct {
		i uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Factorial of 0",
			args: args{i: 0},
			want: 1,
		}, {
			name: "Factorial of 1",
			args: args{i: 1},
			want: 1,
		}, {
			name: "Factorial of 2",
			args: args{i: 2},
			want: 2,
		}, {
			name: "Factorial of 5",
			args: args{i: 5},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Factorial(tt.args.i)
			if got := u; got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "sum of 0",
			args: args{n: 0},
			want: 0,
		}, {
			name: "sum of 1",
			args: args{n: 1},
			want: 1,
		}, {
			name: "sum of 2",
			args: args{n: 2},
			want: 3,
		}, {
			name: "sum of 5",
			args: args{n: 5},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.n); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciNumber(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "fibonacci number of 0",
			args: args{n: 0},
			want: 0,
		}, {
			name: "fibonacci number of 1",
			args: args{n: 1},
			want: 1,
		}, {
			name: "fibonacci number of 2",
			args: args{n: 2},
			want: 1,
		}, {
			name: "fibonacci number 5",
			args: args{n: 5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciNumber(tt.args.n); got != tt.want {
				t.Errorf("FibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciSequence(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "fibonacci of 0",
			args: args{n: 0},
			want: []uint{0, 1},
		}, {
			name: "fibonacci of 1",
			args: args{n: 1},
			want: []uint{0, 1},
		}, {
			name: "fibonacci of 2",
			args: args{n: 2},
			want: []uint{0, 1, 1},
		}, {
			name: "fibonacci of 5",
			args: args{n: 5},
			want: []uint{0, 1, 1, 2, 3, 5},
		}, {
			name: "fibonacci of 10",
			args: args{n: 10},
			want: []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := FibonacciSequence(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreatestCommonDivisor(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "GreatestCommonDivisor of 0, 0",
			args: args{x: 0, y: 0},
			want: 0,
		}, {
			name: "GreatestCommonDivisor of 0, 1",
			args: args{x: 0, y: 1},
			want: 1,
		}, {
			name: "GreatestCommonDivisor of 1, 1",
			args: args{x: 1, y: 1},
			want: 1,
		}, {
			name: "GreatestCommonDivisor of 5, 10",
			args: args{x: 5, y: 10},
			want: 5,
		}, {
			name: "GreatestCommonDivisor of 10, 5",
			args: args{x: 10, y: 5},
			want: 5,
		}, {
			name: "GreatestCommonDivisor of 10, 10",
			args: args{x: 10, y: 10},
			want: 10,
		}, {
			name: "GreatestCommonDivisor of 36, 54",
			args: args{x: 36, y: 54},
			want: 18,
		}, {
			name: "GreatestCommonDivisor of 54, 36",
			args: args{x: 54, y: 36},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreatestCommonDivisor(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GreatestCommonDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
