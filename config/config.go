package config

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type hostArr []string

func (hosts *hostArr) String() string {
	return strings.Join(*hosts, ",")
}

func (hosts *hostArr) Set(value string) error {
	for _, element := range strings.Split(value, ",") {
		*hosts = append(*hosts, element)
	}
	return nil
}
type Config struct {
	Threads int
	Hosts   hostArr
	MaxConnsPerHost int
	MaxIdleConns int
	IdleConnTimeout time.Duration
}

func (config *Config) Parse() {
	flag.IntVar(&config.Threads, "threads", 1, "Number of clients to be emulated")
	flag.Var(&config.Hosts, "hosts", "Comma separated list of Hosts (address:port")
	flag.IntVar(&config.MaxConnsPerHost,
		"maxConnsPerHost",
		1,
		"MaxConnsPerHost optionally limits the total number of connections per host, including connections in the dialing, active, and idle states." +
		" On limit violation, dials will block. Zero means no limit. For HTTP/2, this currently only controls the number of new connections being created at a time, instead of the total number." +
		" In practice, hosts using HTTP/2 only have about one idle connection, though.")
	flag.IntVar(&config.MaxIdleConns,
		"maxIdleConns",
		1,
		"MaxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.")
	flag.DurationVar(&config.IdleConnTimeout,
		"idleConnTimeout",
		10 * time.Second,
		"IdleConnTimeout is the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself. Zero means no limit." +
		" Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\".")
	flag.Parse()
}

func (config Config) Print() {
	fmt.Printf("Threads: %d\n", config.Threads)
	fmt.Printf("Hosts: %s\n", config.Hosts.String())
	fmt.Printf("IdleConnTimeout: %s\n", config.IdleConnTimeout.String())
	fmt.Printf("MaxConnsPerHost: %d\n", config.MaxConnsPerHost)
	fmt.Printf("MaxIdleConns: %d\n", config.MaxIdleConns)
}
