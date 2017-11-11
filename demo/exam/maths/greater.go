package maths

import (
	"math/rand"
)

type PickGreater struct {
	First  int
	Second int

	CorrectAnswer  int
	ProvidedAnswer int
}

func Greater() *PickGreater {

	t := &PickGreater{
		First:  rand.Intn(100),
		Second: rand.Intn(100),
	}
	if t.First > t.Second {
		t.CorrectAnswer = t.First
	} else if t.First < t.Second {
		t.CorrectAnswer = t.Second
	} else {
		t.First++
		t.CorrectAnswer = t.First
	}
	return t
}
