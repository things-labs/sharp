package ternary

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfInt(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfInt64(condition bool, trueVal, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfFloat(condition bool, trueVal, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfString(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
