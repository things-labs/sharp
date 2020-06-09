package digit

import (
	"sort"
	"testing"
)

func TestSortUints(t *testing.T) {
	type args struct {
		a []uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(UintSlice(tt.args.a))

			if !sort.IsSorted(UintSlice(tt.args.a)) {
				t.Error("SortUints test failed")
			}
		})
	}
}

func TestSortInt8s(t *testing.T) {
	type args struct {
		a []int8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int8{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Int8Slice(tt.args.a))

			if !sort.IsSorted(Int8Slice(tt.args.a)) {
				t.Error("SortInt8s test failed")
			}
		})
	}
}

func TestSortUint8s(t *testing.T) {
	type args struct {
		a []uint8
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint8{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Uint8Slice(tt.args.a))

			if !sort.IsSorted(Uint8Slice(tt.args.a)) {
				t.Error("SortUint8s test failed")
			}
		})
	}
}

func TestSortInt16s(t *testing.T) {
	type args struct {
		a []int16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int16{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Int16Slice(tt.args.a))

			if !sort.IsSorted(Int16Slice(tt.args.a)) {
				t.Error("SortInt16s test failed")
			}
		})
	}
}

func TestSortUint16s(t *testing.T) {
	type args struct {
		a []uint16
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint16{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Uint16Slice(tt.args.a))

			if !sort.IsSorted(Uint16Slice(tt.args.a)) {
				t.Error("SortUint16s test failed")
			}
		})
	}
}

func TestSortInt32s(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int32{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Int32Slice(tt.args.a))

			if !sort.IsSorted(Int32Slice(tt.args.a)) {
				t.Error("SortInt32s test failed")
			}
		})
	}
}

func TestSortUint32s(t *testing.T) {
	type args struct {
		a []uint32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint32{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Uint32Slice(tt.args.a))

			if !sort.IsSorted(Uint32Slice(tt.args.a)) {
				t.Error("SortUint32s test failed")
			}
		})
	}
}

func TestSortInt64s(t *testing.T) {
	type args struct {
		a []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []int64{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Int64Slice(tt.args.a))

			if !sort.IsSorted(Int64Slice(tt.args.a)) {
				t.Error("SortInt64s test failed")
			}
		})
	}
}

func TestSortUint64s(t *testing.T) {
	type args struct {
		a []uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sort",
			args: args{a: []uint64{4, 2, 7, 9, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(Uint64Slice(tt.args.a))

			if !sort.IsSorted(Uint64Slice(tt.args.a)) {
				t.Error("SortUint64s test failed")
			}
		})
	}
}
