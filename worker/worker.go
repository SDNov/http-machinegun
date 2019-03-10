package worker

import (
	"encoding/json"
	"fmt"
	"github.com/SDNov/http-machinegun/stat"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	Path string
	Threads int
}

type Message struct {
	Key string
	Payload string
}

func (task Task) StartTask(client *http.Client) {
	path := "http://" + task.Path + "/api/v1/robots/"
	var wg sync.WaitGroup
	wg.Add(task.Threads)
	statistic := new(stat.Statistic)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("Sended %d requests, errors %d\n", statistic.SuccessCount, statistic.ErrorCount)
		}
	}()
	for i := 0; i < task.Threads; i++ {
		go func(iCode string) {
			defer wg.Done()
			url := path + iCode
			for {
				resp, err := client.Get(url)
				if err != nil {
					statistic.IncrementErrorCounter()
					fmt.Println(err)
					return
				}
				if resp.StatusCode != http.StatusOK {
					fmt.Println(resp.Status)
				} else {
					statistic.IncrementSuccessCounter()
					var messages []Message
					if err := json.NewDecoder(resp.Body).Decode(&messages); err != nil {

					}
				}
				if resp.Body != nil {
					err := resp.Body.Close()
					if err != nil {
						fmt.Println(err)
					}
				}
				time.Sleep(100 * time.Millisecond)
			}

		}(strconv.Itoa(i))
		time.Sleep(5 * time.Nanosecond)
	}
	wg.Wait()
	fmt.Printf("Sended %d requests, errors %d\n", statistic.SuccessCount, statistic.ErrorCount)
}
