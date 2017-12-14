package screen

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	c := make(chan interface{})

	go printEvents(c)

	a := New(c)
	h := a.NewButton("World!")
	d := a.NewDialog("Hello")
	d.SetContent(h.ID)
	d.Show()

	time.Sleep(time.Millisecond)

}

func printEvents(c chan interface{}) {
	for e := range c {
		b, err := json.Marshal(e)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T: %s\n", e, string(b))
		}

	}

}
