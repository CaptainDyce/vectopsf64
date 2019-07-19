package intsupp

/////////////////////////////////////////////
// plain predicates
/////////////////////////////////////////////
type Predicate func(value int) bool

func And(ps ...Predicate) Predicate {
	return func(val int) bool {
		for _, p := range ps {
			if !p(val) {
				return false
			}
		}
		return true
	}
}

func Or(ps ...Predicate) Predicate {
	return func(val int) bool {
		for _, p := range ps {
			if p(val) {
				return true
			}
		}
		return false
	}
}

func Xor(l Predicate, r Predicate) Predicate {
	return func(val int) bool {
		return l(val) != r(val)
	}
}

// -> p'(val) = !p(val)
func Not(p Predicate) Predicate {
	return func(val int) bool {
		return !p(val)
	}
}

// -> p(val) = val > value
func GT(value int) Predicate {
	return func(val int) bool {
		return val > value
	}
}

// -> p(val) = val >= value
func GTE(value int) Predicate {
	return func(val int) bool {
		return val >= value
	}
}

// -> p(val) = val < value
func LT(value int) Predicate {
	return func(val int) bool {
		return val < value
	}
}

// -> p(val) = val <= value
func LTE(value int) Predicate {
	return func(val int) bool {
		return val <= value
	}
}

// -> p(val) = val == value
func EQ(value int) Predicate {
	return func(val int) bool {
		return val == value
	}
}

/////////////////////////////////////////////
// fluent predicates
/////////////////////////////////////////////

type FluentPredicate interface {
	Test(value int) bool
	Neg() FluentPredicate
	And(ps ...Predicate) FluentPredicate
	Or(ps ...Predicate) FluentPredicate
	Xor(p Predicate) FluentPredicate
	Pred() Predicate
}

type pred struct {
	p Predicate
}

func Fluent(p Predicate) FluentPredicate {
	return pred{p}
}

func (s pred) Test(value int) bool {
	return s.p(value)
}

func (s pred) Neg() FluentPredicate {
	return pred{Not(s.p)}
}

func (s pred) And(ps ...Predicate) FluentPredicate {
	return pred{And(s.p, And(ps...))}
}

func (s pred) Or(ps ...Predicate) FluentPredicate {
	return pred{Or(s.p, Or(ps...))}
}

func (s pred) Xor(p Predicate) FluentPredicate {
	return pred{Xor(s.p, p)}
}

func (s pred) Pred() Predicate {
	return s.p
}
