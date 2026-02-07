package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

/*
Learning goals
- file handling
- go routines
- http calls

iteration one
- load file top-1m.csv
- read it line per line
- do a post request again the server (do i need an extra package then, to run the penetrator service)

iteration one
- load file top-1m.csv
- use go routines to speed up the processing of the url shortener service
- read it line per line
- do a post request again the server (do i need an extra package then, to run the penetrator service)
*/

/*
Results
- process file wihout post request and: Exceution time: 3.004051375s
*/
type PostRequest struct {
	Url string `json:"url"`
}

func doRequestToService(client *http.Client, url string) {

	host := "http://localhost:8888/shorten"

	payload := PostRequest{Url: url}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Could not encode json payload", err)
	}

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
	Timeout: time.Second * 10,
}

func main() {
	start := time.Now()
	// think about command line app
	file, err := os.Open("files/top-1m.csv")

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

	jobs := make(chan string) // each line is a job
	wg := &sync.WaitGroup{}

	workerCount := 100
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
		fmt.Println(record)
		jobs <- record[1]

	}
	elapsed := time.Since(start)
	fmt.Printf("Exceution time: %s\n", elapsed)
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		doRequestToService(client, url)
		fmt.Printf("[Worker %d] processed: %s\n", id, url)
	}
}

// 3s for CSV file processing
// Exceution time: 17m5.91306025s for CSV file processing and single http request
// Exceution time: 14m55.75640825s for CSV file processing and 5 workers http request
// Exceution time: 14m30.869779833s
