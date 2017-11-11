package exam

import "github.com/abdullin/olympia/demo/exam/maths"

const TaskKindGreater string = "pick-greater"

type Task interface {
	//Kind() string
}

func SimpleExam() []Task {
	var tasks []Task

	for i := 0; i < 10; i++ {
		switch i % 2 {
		case 0:
			tasks = append(tasks, maths.Greater())
		default:
			tasks = append(tasks, maths.SimpleAddition())

		}
	}
	return tasks

}
