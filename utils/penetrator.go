package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func doRequestToService(client *http.Client, url string) {

	host := "http://127.0.0.1:8888/shorten"

	jsonPayload := []byte(`{"url":"` + url + `"}`)

	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
}

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        10000,
		MaxIdleConnsPerHost: 10000,
		MaxConnsPerHost:     10000,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  true,
	},
}

func main() {
	start := time.Now()
	file, err := os.Open("utils/files/top-1m.csv")

	if err != nil {
		fmt.Println("Could not read file", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Could not close file", err)
			return
		}
	}(file)

	jobs := make(chan string, 10000)
	wg := &sync.WaitGroup{}

	workerCount := 500
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, wg)
	}

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Could not read line", err)
			continue
		}
		jobs <- record[1]
	}
	close(jobs)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Exceution time: %s\n", elapsed)
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		doRequestToService(client, url)
	}
}

// first try last year 2025
// 3s for CSV file processing
// Execution time: 17m5.91306025s for CSV file processing and single http request
// Execution time: 14m55.75640825s for CSV file processing and 5 workers http request
// Execution time: 14m30.869779833s
