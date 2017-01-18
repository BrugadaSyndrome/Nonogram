package main

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
		for i := end - 1; i >= 0; i-- {
			stream <- i
		}
		close(stream)
	}()
	return
}

func longestRow(matrix [][]int) (max int) {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > max {
			max = len(matrix[i])
		}
	}
	return
}
