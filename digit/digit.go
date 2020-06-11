// Package digit
package digit

// BreakUint16 break into bytes
func BreakUint16(v uint16) (LoByte, HiByte byte) {
	return byte(v), byte(v >> 8)
}

// BreakUint32  break into bytes
func BreakUint32(v uint32) (Byte0, Byte1, Byte2, Byte3 byte) {
	return byte(v), byte(v >> 8), byte(v >> 16), byte(v >> 24)
}

// BreakUint64 break into uint32
func BreakUint64(v uint64) (Lo32, Hi32 uint32) {
	return uint32(v), uint32(v >> 32)
}

// BuildUint16 combine into uint16
func BuildUint16(loByte, HiByte byte) uint16 {
	return uint16(loByte) | uint16(HiByte)<<8
}

// BuildUint32 combine into uint32
func BuildUint32(Byte0, Byte1, Byte2, Byte3 byte) uint32 {
	return uint32(Byte0) | (uint32(Byte1) << 8) | (uint32(Byte2) << 16) | (uint32(Byte3) << 24)
}

// BuildUint64 combine into uint64
func BuildUint64(Lo32, Hi32 uint32) uint64 {
	return uint64(Lo32) | uint64(Hi32)<<32
}

// ReverseBytes reverse []byte
func ReverseBytes(b []byte) []byte {
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}

	return b
}
