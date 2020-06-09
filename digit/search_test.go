// Package digit
package digit

import (
	"testing"
)

func TestSearchUints(t *testing.T) {
	type args struct {
		a []uint
		x uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUints(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchUints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt8s(t *testing.T) {
	type args struct {
		a []int8
		x int8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int8{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInt8s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchInt8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint8s(t *testing.T) {
	type args struct {
		a []uint8
		x uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint8{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUint8s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchUint8s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt16s(t *testing.T) {
	type args struct {
		a []int16
		x int16
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int16{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInt16s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchInt16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint16s(t *testing.T) {
	type args struct {
		a []uint16
		x uint16
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint16{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUint16s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchUint16s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt32s(t *testing.T) {
	type args struct {
		a []int32
		x int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int32{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInt32s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchInt32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint32s(t *testing.T) {
	type args struct {
		a []uint32
		x uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint32{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUint32s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchUint32s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchInt64s(t *testing.T) {
	type args struct {
		a []int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []int64{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInt64s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchInt64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchUint64s(t *testing.T) {
	type args struct {
		a []uint64
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				a: []uint64{4, 1, 7, 2, 9},
				x: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchUint64s(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("SearchUint64s() = %v, want %v", got, tt.want)
			}
		})
	}
}
