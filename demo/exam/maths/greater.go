package maths

import (
	"math/rand"
	"strconv"
)

type PickGreater struct {
	First  string
	Second string

	CorrectAnswer  string
	ProvidedAnswer string
}

func Greater() *PickGreater {

	first := rand.Intn(99) + 1
	second := rand.Intn(99) + 1

	for second == first {
		second = rand.Intn(99) + 1
	}

	t := &PickGreater{
		First:  strconv.Itoa(first),
		Second: strconv.Itoa(second),
	}
	if first > second {
		t.CorrectAnswer = t.First
	} else if first < second {
		t.CorrectAnswer = t.Second
	}
	return t
}
