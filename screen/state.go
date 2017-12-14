package screen

func newState() *state {
	return &state{
		controls: make(map[ID]interface{}),
	}
}

type state struct {
	counter  int
	controls map[ID]interface{}
	// for now we have only one top object, no overlaying of dialogs over windows
	top ID
}

func (s *state) next(x int) ID {
	return ID(s.counter + x)
}

func (s *state) apply(e interface{}) {
	switch t := e.(type) {
	case *ButtonCreated:
		s.counter++
		s.controls[t.ID] = &buttonCtl{text: t.Text}
	case *DialogCreated:
		s.counter++
		s.controls[t.ID] = &dialogCtl{title: t.Title}
	case *DialogContentChanged:
		s.controls[t.ID].(*dialogCtl).content = t.ContentID
	case *DialogShown:
		s.top = t.ID
		s.controls[t.ID].(*dialogCtl).visible = true
	default:
		panic("Unknown event type")
	}
}

type buttonCtl struct {
	text string
}
type dialogCtl struct {
	title   string
	content ID
	visible bool
}
