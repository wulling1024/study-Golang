package main

func main() {
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)
}

func gen(number ...int) <-chan int {
	out := make(chan int, len(number))
	defer close(out)
	for _, n := range number {
		out <- n
	}
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
