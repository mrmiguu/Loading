package done

import "strings"

func Yet(proc string) chan<- bool {
	c := make(chan bool)
	println(proc + ".")
	space := "   "
	spaces := strings.Repeat(" ", len(proc))
	var left bool
	go func() {
		i := 1
		for done := range c {
			if done {
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
