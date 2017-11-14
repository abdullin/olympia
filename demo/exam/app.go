package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/abdullin/olympia/demo/exam/maths"
	"github.com/abdullin/olympia/demo/lib"
	"github.com/abdullin/olympia/forms"
)

func main() {

	rand.Seed(time.Now().Unix())
	app := Start()

	lib.ConfigureSingleton("/todo", app)

	http.ListenAndServe(":3001", nil)

}

type App struct {
	Screen        string
	Tasks         []Task
	Scores        []int
	Answers       []string
	Changes       chan bool
	Position      int
	ShowingAnswer bool

	Grades []float32
}

func (a *App) Changed() chan bool {
	return a.Changes
}

func (a *App) GradeTest() {

	var score float32
	for _, s := range a.Scores {
		score += float32(s)
	}
	a.Grades = append(a.Grades, score/float32(len(a.Scores)))
}

func (a *App) Dispatch(kind string, args map[string]interface{}) {

	switch kind {
	case "start-test":
		a.Position = 0
		a.Tasks = SimpleExam()
		a.ShowingAnswer = false
		l := len(a.Tasks)
		a.Scores = make([]int, l, l)
		a.Answers = make([]string, l, l)

		a.Screen = "exam"
		a.Changes <- true
		return

	case "correct":
		a.Scores[a.Position] = 1
		a.ShowingAnswer = true

		a.Answers[a.Position] = args["provided"].(string)
		a.Changes <- true

		go func() {
			time.Sleep(time.Second)

			if a.Position == len(a.Tasks)-1 {
				a.Screen = "answers"
				a.GradeTest()

			} else {
				a.ShowingAnswer = false
				a.Position++
			}
			a.Changes <- true
		}()

		return

	case "incorrect":
		a.ShowingAnswer = true

		a.Answers[a.Position] = args["provided"].(string)
		a.Changes <- true

		go func() {
			time.Sleep(time.Second * 1)

			if a.Position == len(a.Tasks)-1 {
				a.Screen = "answers"
				a.GradeTest()
			} else {
				a.ShowingAnswer = false
				a.Position++
			}
			a.Changes <- true

		}()

		return

	}

	log.Print("Unknown dispatch", kind, args)

}
func (a *App) GetScreen() interface{} {
	switch a.Screen {
	case "main":
		return renderMainWindow(a)
	case "exam":
		return renderExamWindow(a)
	case "answers":
		return renderAnswers(a)
	}

	return nil
}

func Start() *App {
	// todo - load state from the DB
	changed := make(chan bool, 100)
	return &App{
		Screen:  "main",
		Changes: changed,
	}
}

func renderExamWindow(a *App) interface{} {

	nop := forms.NewAction("nop")

	title := fmt.Sprintf("%d / %d", a.Position+1, len(a.Tasks))
	w := forms.NewWindow(title)

	t := a.Tasks[a.Position]

	switch k := t.(type) {
	case *maths.MultipleChoice:

		g := forms.NewGrid()

		title := forms.NewTitle(k.Question, 1)
		g.AddRowItems(nil, title, nil).SetSpans(0, 6, 0)

		r := g.AddRow()
		r.AddCol(nil)

		w.AddMenu(k.Question, nil)

		for _, o := range k.Options {

			b := forms.NewButton(o, nop)
			b.Fill = true
			b.Large = true

			if a.ShowingAnswer {
				if o == k.CorrectAnswer {
					b.Style = "success"
				} else {
					b.Style = "danger"
				}
			} else {
				if o == k.CorrectAnswer {
					b.Action = forms.NewAction("correct")

				} else {
					b.Action = forms.NewAction("incorrect")
				}
				b.Action.Args["provided"] = o
			}

			r.AddCol(b)
		}

		r.AddCol(nil)
		w.Content = g

	case *maths.PickGreater:

		w.AddMenu("Выберите наибольшее число", nil)
		g := forms.NewGrid()

		title := forms.NewTitle("Что больше?", 1)
		g.AddRowItems(nil, title, nil).SetSpans(0, 6, 0)

		b := forms.NewButton(k.First, nop).EnlargeAndFill()

		b2 := forms.NewButton(k.Second, nop).EnlargeAndFill()

		g.AddRowItems(nil, b, b2, nil)

		if k.CorrectAnswer == k.First {

			if a.ShowingAnswer {
				b.Style = "success"
				b2.Style = "danger"
			} else {
				b.Action = forms.NewAction("correct")
				b2.Action = forms.NewAction("incorrect")
				b.Action.Args["provided"] = b.Text
				b2.Action.Args["provided"] = b2.Text
			}

		} else {
			if a.ShowingAnswer {
				b2.Style = "success"
				b.Style = "danger"
			} else {

				b2.Action = forms.NewAction("correct")
				b.Action = forms.NewAction("incorrect")
				b.Action.Args["provided"] = b.Text
				b2.Action.Args["provided"] = b2.Text
			}
		}
		w.Content = g
	}

	return w
}

func renderAnswers(a *App) interface{} {
	w := forms.NewWindow("Результаты")

	next := forms.NewButton("Начать новый тест", forms.NewAction("start-test")).EnlargeAndFill()

	t := renderSummaryTable(a)

	history := renderHistoryChart(a)
	stats := renderExamChart(a)

	g := forms.NewGrid()
	g.AddRowItems(nil, forms.NewTitle("Ответы", 1), nil).SetSpans(0, 4, 0)
	g.AddRowItems(nil, t, nil).SetSpans(0, 10, 0)
	g.AddRowItems(nil, stats, history, nil).SetSpans(0, 5, 5, 0)
	g.AddRowItems(nil, next, nil)

	w.Content = g

	return w

}

func renderSummaryTable(a *App) *forms.DataTable {

	t := forms.NewTable()
	t.AddTextColumn("Задание")
	t.AddTextColumn("Ответ")
	t.AddTextColumn("Правильный")

	for i, ts := range a.Tasks {
		switch k := ts.(type) {
		case *maths.PickGreater:

			r := t.AddRow("Больше", a.Answers[i], k.CorrectAnswer)

			if a.Scores[i] == 0 {
				r.Style = "danger"
			} else {
				r.Style = "success"
			}
		case *maths.MultipleChoice:
			r := t.AddRow("Сложение", a.Answers[i], k.CorrectAnswer)

			if a.Scores[i] == 0 {
				r.Style = "danger"
			} else {
				r.Style = "success"
			}
		}
	}
	return t
}

func renderExamChart(a *App) *forms.PieChart {

	correct := 0
	incorrect := 0

	for _, a := range a.Scores {
		if a == 0 {
			incorrect++
		} else {
			correct++
		}
	}

	data := make([]forms.PiePoint, 2, 2)
	data[0].X = "Правильно"
	data[0].Y = float32(correct)

	data[1].X = "Неправильно"
	data[1].Y = float32(incorrect)

	chart := forms.NewPieChart(data)
	return chart
}

func renderHistoryChart(a *App) interface{} {

	var lines []forms.LinePoint

	for i, s := range a.Grades {

		var l forms.LinePoint
		l.X = strconv.Itoa(i + 1)
		l.Y = s
		lines = append(lines, l)
	}
	return forms.NewLineChart(lines)

}

func renderMainWindow(a *App) interface{} {
	w := forms.NewWindow("Тесты")

	w.AddMenu("1Б", nil)

	g := forms.NewGrid()

	startAction := forms.NewAction("start-test")

	b := forms.NewButton("Новый тест", startAction).EnlargeAndFill()

	g.AddRowItems(nil, b, nil)

	w.Content = g
	return w

}
