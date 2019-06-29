package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/docker/go-healthcheck"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world! This is version 2!\n")
	fmt.Fprintln(w, "HOSTNAME:", os.Getenv("HOSTNAME"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	health.RegisterPeriodicFunc("minute_even", time.Second*5, currentMinuteEvenCheck)
}

func currentMinuteEvenCheck() error {
	m := time.Now().Minute()
	if m%2 == 0 {
		return errors.New("Current minute is even!")
	}
	return errors.New("Current minute is not even!")
}
