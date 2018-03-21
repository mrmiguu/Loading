package load

import (
	"strings"
	"sync"
	"time"
)

var (
	addtls  = make(map[chan<- bool]string)
	addtlsl sync.Mutex

	spinner = [...]rune{'/', '-', '\\', '|', '!'}
)

type Proc struct {
	initOnce, doneOnce sync.Once

	Name string

	frame int
}

func (p *Proc) init() {
	p.initOnce.Do(func() {
	})
}

func (p *Proc) Step() {
	p.init()
	p.print()
	p.frame = (p.frame + 1) % (len(spinner) - 1)
}

func (p Proc) print() {
	print(p.Name + ":" + string([]rune{spinner[p.frame]}) + "\r")
}

func (p *Proc) Done() {
	p.init()
	p.doneOnce.Do(func() {
		p.frame = len(spinner) - 1
		p.print()
		println()
	})
}

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
					if is {
						delete(addtls, a)
					}
				}
				addtlsl.Unlock()

				select {
				case a <- d:
				default:
					if is {
						a <- true
					}
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
