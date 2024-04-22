package FetchingFiles

import (
	"fmt"
	"net/http"
)

func FetchData(apiURL string, ch chan<- string) {
	resp, err := http.Get(apiURL)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching data from %s: %s", apiURL, err)
		return
	}
	defer resp.Body.Close()

	data := make([]byte, 100)
	_, err = resp.Body.Read(data)
	if err != nil {
		ch <- fmt.Sprintf("Error reading data from %s: %s", apiURL, err)
		return
	}

	ch <- fmt.Sprintf("Data from %s: %s", apiURL, string(data))
}
