package screen

func (a *Agg) NewButton(text string) *ButtonRef {
	id := a.s.next(1)

	a.Apply(&ButtonCreated{id, text})
	return &ButtonRef{id, a}

}

type ButtonCreated struct {
	ID   ID     `json:"id"`
	Text string `json:"text"`
}

type ButtonRef Ref
