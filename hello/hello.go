package hello

func Add(a int, b int) int {
	c1 := make(chan int)
	go func(a int, b int) int {
		c1 <- a + b
	}()
	select {
	case msg1 := <-c1:
		return c1
	}
}
