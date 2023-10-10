package main

import (
	"fmt"
	"time"
)

type GeneralShutDowner interface {
	ShutDown(duration time.Duration) error
}

type ShutDownFunc func(time.Duration) error

func (f ShutDownFunc) ShutDown(duration time.Duration) error {
	return f(duration)
}

func runPShutDown() {
	fmt.Println("test")
}
