package main

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello, World! ch"
	}

	


}