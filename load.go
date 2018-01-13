package load

import (
	"strings"
	"sync"
	"time"
)

var (
	addtls  = make(map[chan<- bool]string)
	addtlsl sync.Mutex
)

// New adds a process tracker.
func New(PROC string, addtl ...chan<- bool) (DONE chan<- bool) {
	addtlsl.Lock()
	for _, a := range addtl {
		_, found := addtls[a]
		if !found {
			addtls[a] = PROC
		}
	}
	addtlsl.Unlock()

	c := make(chan bool)
	DONE = c
	println(PROC + "...")
	spaces := strings.Repeat(" ", len(PROC))
	go func() {
		var t time.Ticker
		for is := range c {
			for _, a := range addtl {
				addtlsl.Lock()
				proc, found := addtls[a]
				d := false
				if found && proc == PROC { // found
					d = is
					delete(addtls, a)
				}
				addtlsl.Unlock()

				select {
				case a <- d:
				default:
				}
			}
			if is {
				t.Stop()
				close(c)
				println(spaces + "!!!")
				return
			}
			println(spaces + "...")
		}
	}()
	return
}

// Is is deprecated. Pass true into done channel instead.
func Is(done chan<- bool) {
	done <- true
}
