// Package extnet 扩展net包, 提供ip到数值的转换
package extnet

import (
	"encoding/binary"
	"errors"
	"net"
)

// IPToNumber net.IP转换为数值
func IPToNumber(p net.IP) (uint32, error) {
	ip := p.To4()
	if ip == nil || len(ip) < 4 {
		return 0, errors.New("invalid ipv4 address")
	}

	return binary.BigEndian.Uint32(ip), nil
}

// NumberToIP 数值转换为net.IP
func NumberToIP(l uint32) net.IP {
	return net.IPv4(byte(l>>24), byte(l>>16), byte(l>>8), byte(l))
}

// ParseIPToNumber 点分十进制字符串转换数值
func ParseIPToNumber(s string) (uint32, error) {
	return IPToNumber(net.ParseIP(s))
}

// NumberToIPstring 数值转换为点分十进制字符串
func NumberToIPstring(l uint32) string {
	return NumberToIP(l).String()
}

// IPMaskToString mask 转为点分十进制作字符串
func IPMaskToString(mask net.IPMask) string {
	return net.IP(mask).String()
}

// ParseIPMask parses s as an IP address, returning the result.
// The string s can be in dotted decimal ("192.0.2.1")
// or IPv6 ("2001:db8::68") form.
// If s is not a valid textual representation of an IP address,
// ParseIP returns nil.
func ParseIPMask(s string) net.IPMask {
	return net.IPMask(net.ParseIP(s))
}
