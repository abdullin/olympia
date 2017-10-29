package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/abdullin/olympia/pubsub"
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

	app := todo.Start()
	hub := pubsub.NewHub()

	go func() {
		for _ = range app.Changed {
			hub.Publish([]string{"render"}, true)
		}
	}()

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client subscribed")

		var c pubsub.Channel

		c, err = hub.Subscribe([]string{"render"})
		if err != nil {
			fmt.Println(err)
			return
		}

		defer c.Close()

		go runRenderLoop(conn, app, c)

		err = runActionLoop(conn, app)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Client unsubscribed")

	})

	http.ListenAndServe(":3001", nil)
}

type Action struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

func runActionLoop(c *websocket.Conn, app *todo.App) error {
	for {
		action := &Action{}
		if err := c.ReadJSON(action); err != nil {
			log.Printf("Error reading json: %s", err)
			return err
		}
		app.Dispatch(action.Type, action.Args)
	}
}

func runRenderLoop(ws *websocket.Conn, app *todo.App, c pubsub.Channel) error {

	var err error

	if err = ws.WriteJSON(app.GetScreen()); err != nil {
		return err
	}

	for _ = range c.Read() {

		if err = ws.WriteJSON(app.GetScreen()); err != nil {
			return err
		}

	}
	return nil
}
