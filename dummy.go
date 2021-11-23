package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	DefaultTimeoutSeconds = 5
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Println("got a request")

	fmt.Fprint(w, "This is dummy service\n")
}

func main() {
	http.HandleFunc("/", HandleRoot)

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println(sig)

		timeout := os.Getenv("DUMMY_TIMEOUT")
		nTimeout, err := strconv.Atoi(timeout)
		if err != nil {
			fmt.Printf("invalid timeout value '%s', using default %v seconds\n", timeout, DefaultTimeoutSeconds)

			nTimeout = DefaultTimeoutSeconds
		}

		fmt.Println("exiting ...")

		time.Sleep(time.Duration(nTimeout) * time.Second)

		os.Exit(0)
	}()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
