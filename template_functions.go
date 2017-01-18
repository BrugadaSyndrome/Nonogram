package main

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func ascending(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i < end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func descending(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start - 1; i >= end; i-- {
			stream <- i
		}
		close(stream)
	}()
	return
}

func longestList(matrix [][]int) (max int) {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > max {
			max = len(matrix[i])
		}
	}
	return
}

func subtract(a, b int) (dif int) {
	dif = a - b
	return
}
