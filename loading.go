package load

import "strings"
import "time"

func Ing(proc string, d ...time.Duration) chan<- bool {
	c := make(chan bool)
	println(proc + ".")
	space := "   "
	spaces := strings.Repeat(" ", len(proc))
	var left bool
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
		i := 1
		for done := range c {
			if done {
				t.Stop()
				println(spaces + "!!!")
				return
			}
			strs := strings.Split(space, "")
			strs[i] = "."
			dot := strings.Join(strs, "")
			println(spaces + dot)
			if !left && i == 2 {
				left = true
			} else if left && i == 0 {
				left = false
			}
			if !left {
				i++
			} else {
				i--
			}
		}
	}()
	return c
}
