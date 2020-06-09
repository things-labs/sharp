// Package digit
package digit

import (
	"reflect"
)

// ContainInt checks if x exists in []ints and returns TRUE if x is found.
func ContainInt(x int, y []int) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint checks if x exists in []uints and returns TRUE if x is found.
func ContainUint(x uint, y []uint) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt8 checks if x exists in []int8s and returns TRUE if x is found.
func ContainInt8(x int8, y []int8) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint8 checks if x exists in []uint8s and returns TRUE if x is found.
func ContainUint8(x uint8, y []uint8) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt16 checks if x exists in []int16s and returns TRUE if x is found.
func ContainInt16(x int16, y []int16) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint16 checks if x exists in []uint16s and returns TRUE if x is found.
func ContainUint16(x uint16, y []uint16) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt32 checks if x exists in []int32s and returns TRUE if x is found.
func ContainInt32(x int32, y []int32) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint32 checks if x exists in []uint32s and returns TRUE if x is found.
func ContainUint32(x uint32, y []uint32) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainInt64 checks if x exists in []int64s and returns TRUE if x is found.
func ContainInt64(x int64, y []int64) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainUint64 checks if x exists in []uint64s and returns TRUE if x is found.
func ContainUint64(x uint64, y []uint64) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// ContainFloat64 checks if x exists in []float64s and returns TRUE if x is found.
func ContainFloat64(x float64, y []float64) bool {
	for _, v := range y {
		if x == v {
			return true
		}
	}
	return false
}

// Contain checks if x exists in a slice and returns TRUE if x is found.
func Contain(x interface{}, y []interface{}) bool {
	for _, v := range y {
		if reflect.DeepEqual(x, v) {
			return true
		}
	}
	return false
}
