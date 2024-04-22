package main

import (
	"main.go/downloadFiles"
	"sync"
)

const maxWorkers = 5

func main() {
	urls := []string{"https://oldmy.sdu.edu.kz/index.php", "https://sdu.edu.kz", "https://www.rii.edu.kz"}

	var wg sync.WaitGroup
	maxWorkers := 2
	worker := make(chan struct{}, maxWorkers)

	for _, url := range urls {
		wg.Add(1)
		worker <- struct{}{}
		go func(url string) {
			defer func() { <-worker }()
			downloadFiles.Download(url, &wg, worker)
		}(url)
	}

	wg.Wait()
	close(worker)
}
