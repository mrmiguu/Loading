package load

import "strings"
import "time"

func New(proc string, d ...time.Duration) chan<- bool {
	c := make(chan bool)
	println(proc + "...")
	spaces := strings.Repeat(" ", len(proc))
	go func() {
		var t time.Ticker
		if len(d) > 1 {
			panic("too many arguments")
		}
		if len(d) > 0 {
			t = *time.NewTicker(d[0])
			go func() {
				for range t.C {
					c <- false
				}
			}()
		}
		for done := range c {
			if done {
				t.Stop()
				println(spaces + "!!!")
				return
			}
			println(spaces + "...")
		}
	}()
	return c
}
