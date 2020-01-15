package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func downloadFile(url, filename string, wg *sync.WaitGroup, retries int) {
	defer wg.Done()

	for i := 0; i <= retries; i++ {
		out, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error downloading file %s: %v\n", filename, err)
			out.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		_, err = io.Copy(out, resp.Body)
		out.Close()
		resp.Body.Close()

		if err != nil {
			fmt.Printf("Error saving file %s: %v\n", filename, err)
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Download completed: %s\n", filename)
		return
	}

	fmt.Printf("Failed to download %s after %d retries\n", filename, retries)
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Printf("Error opening urls.txt: %v\n", err)
		return
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading urls.txt: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	retries := 3
	for i, url := range urls {
		wg.Add(1)
		go downloadFile(url, fmt.Sprintf("file%d", i+1), &wg, retries)
	}

	wg.Wait()

	fmt.Println("All downloads completed!")
}
