package vectopsf64

import (
	"fmt"
	fs "github.com/CaptainDyce/vectopsf64/f64supp"
	is "github.com/CaptainDyce/vectopsf64/intsupp"
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
	for i, _ := range s {
		s[i] = f(i)
	}
	return s
}

func (s VectOp) ApplyOp(f fs.Operator) VectOp {
	for i, val := range s {
		s[i] = f(val)
	}
	return s
}

func (s VectOp) ApplyOpi(f fs.IndexedOperator) VectOp {
	for i, val := range s {
		s[i] = f(i, val)
	}
	return s
}

func (s VectOp) Ident() VectOp {
	return s.Apply(fs.CoerceInt)
}

func (s VectOp) Setl(value float64) VectOp {
	for i, _ := range s {
		s[i] = value
	}
	return s
}

func (s VectOp) Setv(v []float64) VectOp {
	s.accept(v)
	for i, _ := range s {
		s[i] = v[i]
	}
	return s
}

func (s VectOp) SetMaskl(value float64, p is.Predicate) VectOp {
	for i, _ := range s {
		if p(i) {
			s[i] = value
		}
	}
	return s
}

func (s VectOp) SetMaskv(v []float64, p is.Predicate) VectOp {
	s.accept(v)
	for i, _ := range s {
		if p(i) {
			s[i] = v[i]
		}
	}
	return s
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
	s.accept(v)
	return s.PlusOp(fs.Get(v))
}

// -> s'[i] = s[i] + value
func (s VectOp) Plusl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return val + value })
}

// -> s'[i] = s[i] + o(i)
func (s VectOp) PlusOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val + o(i) })
}

// -> s'[i] = s[i] + o(i, s[i])
func (s VectOp) PlusOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val + o(i, val) })
}

/////////////////
// Minus
/////////////////
// -> s'[i] = s[i] - v[i]
func (s VectOp) Minusv(v []float64) VectOp {
	s.accept(v)
	return s.MinusOp(fs.Get(v))
}

// -> s'[i] = s[i] - value
func (s VectOp) Minusl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return val - value })
}

// -> s'[i] = s[i] + o(i)
func (s VectOp) MinusOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val - o(i) })
}

// -> s'[i] = s[i] + o(i, s[i])
func (s VectOp) MinusOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val - o(i, val) })
}

/////////////////
// Times
/////////////////
// -> s'[i] = s[i] * v[i]
func (s VectOp) Timesv(v []float64) VectOp {
	s.accept(v)
	return s.TimesOp(fs.Get(v))
}

// -> s'[i] = s[i] * value
func (s VectOp) Timesl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return val * value })
}

// -> s'[i] = s[i] * o(i)
func (s VectOp) TimesOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val * o(i) })
}

// -> s'[i] = s[i] * o(i, s[i])
func (s VectOp) TimesOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val * o(i, val) })
}

/////////////////
// Div
/////////////////
// -> s'[i] = s[i] / v[i]
func (s VectOp) Divv(v []float64) VectOp {
	s.accept(v)
	return s.DivOp(fs.Get(v))
}

// -> s'[i] = s[i] / value
func (s VectOp) Divl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return val / value })
}

// -> s'[i] = s[i] / o(i)
func (s VectOp) DivOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val / o(i) })
}

// -> s'[i] = s[i] / o(i, s[i])
func (s VectOp) DivOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return val / o(i, val) })
}

/////////////////
// Pow
/////////////////
// -> s'[i] = s[i] ^ v[i] (as in e.g. 2^3...)
func (s VectOp) Powv(v []float64) VectOp {
	s.accept(v)
	return s.PowOp(fs.Get(v))
}

// -> s'[i] = s[i] ^ value
func (s VectOp) Powl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return math.Pow(val, value) })
}

// -> s'[i] = s[i] ^ o(i)
func (s VectOp) PowOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Pow(val, o(i)) })
}

// -> s'[i] = s[i] ^ o(i, s[i])
func (s VectOp) PowOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Pow(val, o(i, val)) })
}

/////////////////
// Max
/////////////////
// -> s'[i] = max(s[i], v[i])
func (s VectOp) Maxv(v []float64) VectOp {
	s.accept(v)
	return s.MaxOp(fs.Get(v))
}

// -> s'[i] = max(s[i], value)
func (s VectOp) Maxl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return math.Max(val, value) })
}

// -> s'[i] = max(s[i], o(i))
func (s VectOp) MaxOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Max(val, o(i)) })
}

// -> s'[i] = max(s[i], o(i, s[i]))
func (s VectOp) MaxOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Max(val, o(i, val)) })
}

/////////////////
// Min
/////////////////
// -> s'[i] = min(s[i], v[i])
func (s VectOp) Minv(v []float64) VectOp {
	s.accept(v)
	return s.MinOp(fs.Get(v))
}

// -> s'[i] = min(s[i], value)
func (s VectOp) Minl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return math.Min(val, value) })
}

// -> s'[i] = min(s[i], o(i))
func (s VectOp) MinOp(o fs.IndexedFunc) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Min(val, o(i)) })
}

// -> s'[i] = min(s[i], o(i, s[i]))
func (s VectOp) MinOpi(o fs.IndexedOperator) VectOp {
	return s.ApplyOpi(func(i int, val float64) float64 { return math.Min(val, o(i, val)) })
}

///////////////////////////
// misc...
///////////////////////////

func (s VectOp) Rev() VectOp {
	size := len(s)
	for i := 0; i < size/2; i++ {
		tmp := s[i]
		s[i] = s[size-i-1]
		s[size-i-1] = tmp
	}
	return s
}

// -> s'[i] = -s[i]
func (s VectOp) Neg() VectOp {
	return s.ApplyOp(fs.Neg)
}

// -> s'[i] = abs(s[i])
func (s VectOp) Abs() VectOp {
	return s.ApplyOp(math.Abs)
}

// -> s'[i] = value / s[i]
func (s VectOp) Idivl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return value / val })
}

// -> s'[i] = 1 / s[i]
func (s VectOp) Inv() VectOp {
	return s.Idivl(1)
}

// -> s'[i] = log(s[i])
func (s VectOp) Log() VectOp {
	return s.ApplyOp(math.Log)
}

// -> s'[i] = exp(s[i])
func (s VectOp) Exp() VectOp {
	return s.ApplyOp(math.Exp)
}

// -> s'[i] = value ^ s[i]
func (s VectOp) Expl(value float64) VectOp {
	return s.ApplyOp(func(val float64) float64 { return math.Pow(value, val) })
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
	return s.Reduce(fs.Plus, 0)
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
	return s.Reduce(math.Max, math.Inf(-1))
}

// -> min(s[i]), +inf if empty
func (s VectOp) Min() float64 {
	return s.Reduce(math.Min, math.Inf(+1))
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
