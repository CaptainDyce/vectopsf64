package vectopsf64

import (
	"fmt"
	fs "github.com/CaptainDyce/f64supp"
	is "github.com/CaptainDyce/intsupp"
	"math"
)

type VectOp []float64

///////////////////////////
// bootstrap...
///////////////////////////

func On(v []float64) VectOp {
	return v
}

func OnSize(size int) VectOp {
	return make(VectOp, size)
}

func OnConst(value float64, size int) VectOp {
	return OnSize(size).Setl(value)
}

func OnIdent(size int) VectOp {
	return OnSize(size).Ident()
}

func OnInts(values []int) VectOp {
	return OnSize(len(values)).Apply(fs.CoerceInts(values))
}

///////////////////////////
// operations...
///////////////////////////

func (s VectOp) Apply(f fs.IndexedFunc) VectOp {
	return fs.Apply(s, f)
}

func (s VectOp) ApplyOp(f fs.Operator) VectOp {
	return fs.ApplyOp(s, f)
}

func (s VectOp) ApplyOpi(f fs.IndexedOperator) VectOp {
	return fs.ApplyOpi(s, f)
}

func (s VectOp) Ident() VectOp {
	return s.Apply(fs.CoerceInt)
}

func (s VectOp) Setl(value float64) VectOp {
	return fs.Setl(s, value)
}

func (s VectOp) Setv(v []float64) VectOp {
	return fs.Setv(s, v)
}

func (s VectOp) SetMaskl(value float64, p is.Predicate) VectOp {
	return fs.SetMaskl(s, value, p)
}

func (s VectOp) SetMaskv(v []float64, p is.Predicate) VectOp {
	return fs.SetMaskv(s, v, p)
}

func (s VectOp) accept(v []float64) {
	if len(v) < len(s) {
		panic(fmt.Sprintf("invalid array size %d (out of bounds for %d-element vector)", len(v), len(s)))
	}
}

/////////////////
// Plus
/////////////////
// -> s'[i] = s[i] + v[i]
func (s VectOp) Plusv(v []float64) VectOp {
	return fs.Plusv(s, v)
}

// -> s'[i] = s[i] + value
func (s VectOp) Plusl(value float64) VectOp {
	return fs.Plusl(s, value)
}

// -> s'[i] = s[i] + o(i)
func (s VectOp) PlusOp(o fs.IndexedFunc) VectOp {
	return fs.PlusOp(s, o)
}

// -> s'[i] = s[i] + o(i, s[i])
func (s VectOp) PlusOpi(o fs.IndexedOperator) VectOp {
	return fs.PlusOpi(s, o)
}

/////////////////
// Minus
/////////////////
// -> s'[i] = s[i] - v[i]
func (s VectOp) Minusv(v []float64) VectOp {
	return fs.Minusv(s, v)
}

// -> s'[i] = s[i] - value
func (s VectOp) Minusl(value float64) VectOp {
	return fs.Minusl(s, value)
}

// -> s'[i] = s[i] + o(i)
func (s VectOp) MinusOp(o fs.IndexedFunc) VectOp {
	return fs.MinusOp(s, o)
}

// -> s'[i] = s[i] + o(i, s[i])
func (s VectOp) MinusOpi(o fs.IndexedOperator) VectOp {
	return fs.MinusOpi(s, o)
}

/////////////////
// Times
/////////////////
// -> s'[i] = s[i] * v[i]
func (s VectOp) Timesv(v []float64) VectOp {
	return fs.Timesv(s, v)
}

// -> s'[i] = s[i] * value
func (s VectOp) Timesl(value float64) VectOp {
	return fs.Timesl(s, value)
}

// -> s'[i] = s[i] * o(i)
func (s VectOp) TimesOp(o fs.IndexedFunc) VectOp {
	return fs.TimesOp(s, o)
}

// -> s'[i] = s[i] * o(i, s[i])
func (s VectOp) TimesOpi(o fs.IndexedOperator) VectOp {
	return fs.TimesOpi(s, o)
}

/////////////////
// Div
/////////////////
// -> s'[i] = s[i] / v[i]
func (s VectOp) Divv(v []float64) VectOp {
	return fs.Divv(s, v)
}

// -> s'[i] = s[i] / value
func (s VectOp) Divl(value float64) VectOp {
	return fs.Divl(s, value)
}

// -> s'[i] = s[i] / o(i)
func (s VectOp) DivOp(o fs.IndexedFunc) VectOp {
	return fs.DivOp(s, o)
}

// -> s'[i] = s[i] / o(i, s[i])
func (s VectOp) DivOpi(o fs.IndexedOperator) VectOp {
	return fs.DivOpi(s, o)
}

/////////////////
// Pow
/////////////////
// -> s'[i] = s[i] ^ v[i] (as in e.g. 2^3...)
func (s VectOp) Powv(v []float64) VectOp {
	return fs.Powv(s, v)
}

// -> s'[i] = s[i] ^ value
func (s VectOp) Powl(value float64) VectOp {
	return fs.Powl(s, value)
}

// -> s'[i] = s[i] ^ o(i)
func (s VectOp) PowOp(o fs.IndexedFunc) VectOp {
	return fs.PowOp(s, o)
}

// -> s'[i] = s[i] ^ o(i, s[i])
func (s VectOp) PowOpi(o fs.IndexedOperator) VectOp {
	return fs.PowOpi(s, o)
}

/////////////////
// Max
/////////////////
// -> s'[i] = max(s[i], v[i])
func (s VectOp) Maxv(v []float64) VectOp {
	return fs.Maxv(s, v)
}

// -> s'[i] = max(s[i], value)
func (s VectOp) Maxl(value float64) VectOp {
	return fs.Maxl(s, value)
}

// -> s'[i] = max(s[i], o(i))
func (s VectOp) MaxOp(o fs.IndexedFunc) VectOp {
	return fs.MaxOp(s, o)
}

// -> s'[i] = max(s[i], o(i, s[i]))
func (s VectOp) MaxOpi(o fs.IndexedOperator) VectOp {
	return fs.MaxOpi(s, o)
}

/////////////////
// Min
/////////////////
// -> s'[i] = min(s[i], v[i])
func (s VectOp) Minv(v []float64) VectOp {
	return fs.Minv(s, v)
}

// -> s'[i] = min(s[i], value)
func (s VectOp) Minl(value float64) VectOp {
	return fs.Minl(s, value)
}

// -> s'[i] = min(s[i], o(i))
func (s VectOp) MinOp(o fs.IndexedFunc) VectOp {
	return fs.MinOp(s, o)
}

// -> s'[i] = min(s[i], o(i, s[i]))
func (s VectOp) MinOpi(o fs.IndexedOperator) VectOp {
	return fs.MinOpi(s, o)
}

///////////////////////////
// misc...
///////////////////////////

func (s VectOp) Rev() VectOp {
	return fs.Rev(s)
}

// -> s'[i] = -s[i]
func (s VectOp) Neg() VectOp {
	return fs.Negv(s)
}

// -> s'[i] = abs(s[i])
func (s VectOp) Abs() VectOp {
	return fs.Abs(s)
}

// -> s'[i] = value / s[i]
func (s VectOp) Idivl(value float64) VectOp {
	return fs.Idivl(s, value)
}

// -> s'[i] = 1 / s[i]
func (s VectOp) Inv() VectOp {
	return fs.Idivl(s, 1)
}

// -> s'[i] = log(s[i])
func (s VectOp) Log() VectOp {
	return fs.Log(s)
}

// -> s'[i] = exp(s[i])
func (s VectOp) Exp() VectOp {
	return fs.Exp(s)
}

// -> s'[i] = value ^ s[i]
func (s VectOp) Expl(value float64) VectOp {
	return fs.Expl(s, value)
}

///////////////////////////
// slicing...
///////////////////////////

// -> s[0..n[
func (s VectOp) Head(n int) VectOp {
	return On(s[:n])
}

// -> s[n..end[
func (s VectOp) Tail(n int) VectOp {
	return On(s[n:])
}

func (s VectOp) Slice(start int, end int) VectOp {
	return On(s[start:end])
}

/////////////////
// Terminals
/////////////////

// -> index of the first occurrence of the specified value in this vector, or < 0 if absent...
func (s VectOp) IndexOfVal(value float64) int {
	return s.IndexOf(fs.EQ(value))
}

func (s VectOp) IndexOf(p fs.Predicate) int {
	for i, val := range s {
		if p(val) {
			return i
		}
	}
	return -1
}

// --> indexes of all occurrences of the specified value in this vector, or empty...
func (s VectOp) IndexesOfVal(value float64) []int {
	return s.IndexesOf(fs.EQ(value))
}

// -> indexes of all occurrences of the specified value in this vector, or empty...
func (s VectOp) IndexesOf(p fs.Predicate) []int {
	var indexes []int
	for i, val := range s {
		if p(val) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// -> s[last]
func (s VectOp) Last() float64 {
	return s[len(s)-1]
}

// -> f[v[n - 1], ... f[v[2], f[v[1], f[v[0], identity]]]]
func (s VectOp) Reduce(f fs.ReduceOperator, identity float64) float64 {
	result := identity
	for _, val := range s {
		result = f(val, result)
	}
	return result
}

// -> sum(s[i])
func (s VectOp) Sum() float64 {
	sum := 0.0
	for _, val := range s {
		sum += val
	}
	return sum
}

// -> s.v
func (s VectOp) Dot(v []float64) float64 {
	s.accept(v)
	sum := 0.0
	for i, val := range s {
		sum += val * v[i]
	}
	return sum
}

// -> max(s[i]), -inf if empty
func (s VectOp) Max() float64 {
	max := math.Inf(-1)
	for _, val := range s {
		max = math.Max(max, val)
	}
	return max
}

// -> min(s[i]), +inf if empty
func (s VectOp) Min() float64 {
	min := math.Inf(+1)
	for _, val := range s {
		min = math.Min(min, val)
	}
	return min
}

// -> c(s[i])...
func (s VectOp) ForEach(c fs.Consumer) {
	for _, val := range s {
		c(val)
	}
}

// -> c(i, s[i])...
func (s VectOp) ForEachIndexed(c fs.IndexedConsumer) {
	for i, val := range s {
		c(i, val)
	}
}

// -> <-chan = <-s[i]...
func (s VectOp) Stream() <-chan float64 {
	ch := make(chan float64)
	go feed(ch, s)
	return ch
}

func feed(ch chan float64, v VectOp) {
	for _, val := range v {
		ch <- val
	}
	close(ch)
}
