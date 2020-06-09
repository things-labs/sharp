// Package digit
package digit

import "testing"

func TestContainInt(t *testing.T) {
	type args struct {
		x int
		y []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []int{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []int{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint(t *testing.T) {
	type args struct {
		x uint
		y []uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []uint{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []uint{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt8(t *testing.T) {
	type args struct {
		x int8
		y []int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []int8{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []int8{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt8(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint8(t *testing.T) {
	type args struct {
		x uint8
		y []uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []uint8{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []uint8{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint8(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt16(t *testing.T) {
	type args struct {
		x int16
		y []int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []int16{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []int16{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt16(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint16(t *testing.T) {
	type args struct {
		x uint16
		y []uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []uint16{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []uint16{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint16(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt32(t *testing.T) {
	type args struct {
		x int32
		y []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []int32{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []int32{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt32(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint32(t *testing.T) {
	type args struct {
		x uint32
		y []uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []uint32{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []uint32{2, 4, 6, 7, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint32(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainInt64(t *testing.T) {
	type args struct {
		x int64
		y []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []int64{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []int64{5, 2, 4, 7, 6, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainInt64(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainUint64(t *testing.T) {
	type args struct {
		x uint64
		y []uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []uint64{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4,
				y: []uint64{5, 2, 4, 7, 6, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainUint64(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainFloat64(t *testing.T) {
	type args struct {
		x float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []float64{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: 4.4,
				y: []float64{2.3, 4.4, 6.7, 7.2, 1.9, 3.5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainFloat64(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("ContainFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContain(t *testing.T) {
	type args struct {
		x interface{}
		y []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				x: 4,
				y: []interface{}{},
			},
			want: false,
		},
		{
			name: "t1",
			args: args{
				x: "iiinsomnia",
				y: []interface{}{1, "test", "iiinsomnia", 2.9, true},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}
