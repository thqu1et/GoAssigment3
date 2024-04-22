package main

import (
	"fmt"
	"main.go/FetchingFiles"
)

func main() {
	var apis = []string{"https://jsonplaceholder.typicode.com/posts/8"}

	ch := make(chan string)

	for _, url := range apis {
		go FetchingFiles.FetchData(url, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}
}
