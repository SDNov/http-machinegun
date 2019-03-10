package main

import (
	"fmt"
	"github.com/SDNov/http-machinegun/config"
	"github.com/SDNov/http-machinegun/worker"

	"net/http"
)

func main() {

	configuration := &config.Config{}
	configuration.Parse()
	configuration.Print()

	transport := &http.Transport{
		MaxConnsPerHost: configuration.MaxConnsPerHost,
		MaxIdleConns:    configuration.MaxIdleConns,
		IdleConnTimeout: configuration.IdleConnTimeout,
	}
	client := &http.Client{
		Transport: transport,
	}

	task := worker.Task{configuration.Hosts[0], configuration.Threads}
	fmt.Println("Starting...")

	task.StartTask(client)

	fmt.Println("Finished!")
}
