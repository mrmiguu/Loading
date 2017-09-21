package load

import "strings"
import "time"

func Ing(proc string, d ...time.Duration) chan<- bool {
	c := make(chan bool)
	println(proc + "...")
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
			if !left && i == 2 {
				left = true
				strs[i] = `|`
				i--
			} else if left && i == 0 {
				left = false
				strs[i] = `|`
				i++
			} else if !left {
				strs[i] = `\`
				i++
			} else if left {
				strs[i] = `/`
				i--
			}
			dot := strings.Join(strs, "")
			println(spaces + dot)
		}
	}()
	return c
}
