package intsupp

///////////////////////////
// interface...
///////////////////////////

type IndexedFunc func(index int) int
type Operator func(value int) int
type IndexedOperator func(index int, value int) int
type Consumer func(value int)
type IndexedConsumer func(index int, value int)
type ReduceOperator func(vleft int, vright int) int

///////////////////////////
// building...
///////////////////////////

// -> f(i) = c
func Constant(constant int) IndexedFunc {
	return func(i int) int {
		return constant
	}
}

func Identity(i int) int {
	return i
}

// -> f(i) = values[i]
func Get(values []int) IndexedFunc {
	return func(i int) int {
		return values[i]
	}
}

func Mod(value int) Operator {
	return func(val int) int {
		return val % value
	}
}

func Plus(vleft int, vright int) int {
	return vleft + vright
}

func Minus(vleft int, vright int) int {
	return vleft - vright
}

func Times(vleft int, vright int) int {
	return vleft * vright
}

func Div(vleft int, vright int) int {
	return vleft / vright
}

func Max(vleft int, vright int) int {
	if vleft > vright {
		return vleft
	}
	return vright
}

func Min(vleft int, vright int) int {
	if vleft < vright {
		return vleft
	}
	return vright
}

func Neg(value int) int {
	return -value
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
