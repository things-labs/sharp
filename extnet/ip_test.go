package extnet

import (
	"net"
	"reflect"
	"testing"
)

func TestIPToNumber(t *testing.T) {
	type args struct {
		p net.IP
	}

	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{"valid IP", args{net.ParseIP("10.10.0.1")}, 0x0a0a0001, false},
		{"invalid IP", args{net.ParseIP("10.10.x")}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPToNumber(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IPToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumerToIP(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{"numb", args{0x0a0a0001}, net.ParseIP("10.10.0.1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToIP(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberToIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseIPToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{"ip string to number", args{"10.10.0.1"}, 0x0a0a0001, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIPToNumber(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIPToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseIPToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberToIPstring(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"number to ip string", args{0x0a0a0001}, "10.10.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToIPstring(tt.args.l); got != tt.want {
				t.Errorf("NumberToIPstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPMaskToString(t *testing.T) {
	type args struct {
		mask net.IPMask
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{net.IPMask(net.ParseIP("255.255.255.0"))}, "255.255.255.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPMaskToString(tt.args.mask); got != tt.want {
				t.Errorf("IPMaskToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseIPMask(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want net.IPMask
	}{
		{"string to ip mask", args{"255.255.255.0"}, net.IPMask(net.ParseIP("255.255.255.0"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseIPMask(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dot2Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}
