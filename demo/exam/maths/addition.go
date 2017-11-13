package maths

import (
	"fmt"
	"math/rand"
	"strconv"
)

type MultipleChoice struct {
	Question string
	Options  []string

	CorrectAnswer  string
	ProvidedAnswer string
}

func SimpleAddition() *MultipleChoice {

	p1 := rand.Intn(4) + 1
	p2 := rand.Intn(5) + 1

	q := fmt.Sprintf("%d + %d = ?", p1, p2)

	answer := strconv.Itoa(p1 + p2)

	options := []string{
		strconv.Itoa(p1 + p2 - 1),

		strconv.Itoa(p1 + p2 + 1),

		strconv.Itoa(p1 + p2 + 3),

		answer,
	}
	for i := range options {
		j := rand.Intn(i + 1)
		options[i], options[j] = options[j], options[i]
	}
	return &MultipleChoice{
		Question:      q,
		Options:       options,
		CorrectAnswer: answer,
	}

}
