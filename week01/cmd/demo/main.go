package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"github.com/canteen_ns/go-container-lab/week01/pkg/sysinfo"
)

func main() {
	go func() {
		fmt.Println("pprof on :6060")
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	cores, err := sysinfo.CPUCores()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Week01-Day1: %d cores\n", cores)
	busyLoop()
}

func busyLoop() {
	for {
		i := 1
		_ = i * 2
	}
}
