package downloadFiles

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func Download(url string, wg *sync.WaitGroup, worker chan struct{}) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading:", err)
		return
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	out, err := os.Create("downloaded_" + filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
