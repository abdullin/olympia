package screen

func New(c chan interface{}) *Agg {
	return &Agg{c: c, s: newState()}
}

type Agg struct {
	c chan interface{}
	s *state
}

func (a *Agg) Apply(e interface{}) {
	a.s.apply(e)
	a.c <- e
}

type DialogShown struct {
	ID ID `json:"id"`
}
