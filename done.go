package done

func Yet(proc string) chan<- bool {
  c := make(chan bool)
  println(proc+".")
  spaces := strings.Repeat(" ",len(proc))
  go func() {
    i := 1
    for done := range c {
      if done {
        println(spaces+"!!!")
        return
      }
      println(spaces+"...")
    }
  }()
  return c
}
