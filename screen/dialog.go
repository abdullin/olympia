package screen

func (a *Agg) NewDialog(title string) *DialogRef {
	id := a.s.next(1)
	a.Apply(&DialogCreated{id, title})

	return &DialogRef{id, a}
}

func (d *DialogRef) SetContent(i ID) {
	s := d.a.s

	ctl, found := s.controls[d.ID]
	if !found {
		panic("unknown dialog")
	}

	dc := ctl.(*dialogCtl)
	if dc.content == i {
		return
	}
	d.a.Apply(&DialogContentChanged{ID: d.ID, ContentID: i, PreviousID: dc.content})
}

func (d *DialogRef) Show() {
	s := d.a.s

	_, found := s.controls[d.ID]
	if !found {
		panic("unknown dialog")
	}

	if s.top == d.ID {
		return
	}
	// TODO: hide previous

	d.a.Apply(&DialogShown{ID: d.ID})
}

type DialogRef Ref

type DialogCreated struct {
	ID    ID     `json:"id"`
	Title string `json:"title"`
}

type DialogContentChanged struct {
	ID         ID `json:"id"`
	ContentID  ID `json:"contentId"`
	PreviousID ID `json:"previousId"`
}
