package main

import (
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/abdullin/olympia/todo"

	"github.com/gorilla/websocket"
)

func main() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client subscribed")
		err = runTodoApp(conn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Client unsubscribed")

	})

	http.ListenAndServe(":3000", nil)
}

func runTodoApp(c *websocket.Conn) error {

	app := todo.Start()

	go func() {
		for i := 0; i < 999; i++ {

			time.Sleep(time.Second * 2)
			app.AddTask(fmt.Sprintf("Task %d", i), "Normal")
		}
	}()
	var err error

	if err = c.WriteJSON(app.GetScreen()); err != nil {
		return err
	}

	for _ = range app.Changed {

		if err = c.WriteJSON(app.GetScreen()); err != nil {
			return err
		}

	}
	return nil

}
