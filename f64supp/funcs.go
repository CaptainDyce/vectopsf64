package f64supp

///////////////////////////
// interface...
///////////////////////////

type IndexedFunc func(index int) float64
type Operator func(value float64) float64
type IndexedOperator func(index int, value float64) float64
type Consumer func(value float64)
type IndexedConsumer func(index int, value float64)
type ReduceOperator func(vleft float64, vright float64) float64

///////////////////////////
// building...
///////////////////////////

// -> f(i) = c
func Constant(constant float64) IndexedFunc {
	return func(i int) float64 {
		return constant
	}
}

func CoerceInt(i int) float64 {
	return float64(i)
}

func CoerceInts(values []int) IndexedFunc {
	return func(i int) float64 {
		return float64(values[i])
	}
}

// -> f(i) = values[i]
func Get(values []float64) IndexedFunc {
	return func(i int) float64 {
		return values[i]
	}
}

func Plus(vleft float64, vright float64) float64 {
	return vleft + vright
}

func Minus(vleft float64, vright float64) float64 {
	return vleft - vright
}

func Times(vleft float64, vright float64) float64 {
	return vleft * vright
}

func Div(vleft float64, vright float64) float64 {
	return vleft / vright
}

func Neg(value float64) float64 {
	return -value
}
