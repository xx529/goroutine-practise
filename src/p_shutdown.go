package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

type GeneralShutDowner interface {
	ShutDown(duration time.Duration) error
}

type ShutDownFunc func(time.Duration) error

func (f ShutDownFunc) ShutDown(duration time.Duration) error {
	return f(duration)
}

func ConcurrentShutdown(waitTimeout time.Duration, shutDowners ...GeneralShutDowner) error {
	c := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, s := range shutDowners {
			wg.Add(1)
			go func(s GeneralShutDowner) {
				defer wg.Done()
				e := s.ShutDown(waitTimeout)
				if e != nil {
					log.Println("error")
				}

			}(s)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(waitTimeout)
	defer timer.Stop()

	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("time out")
	}
}

func ShutDownMaker(t int) func(duration time.Duration) error {
	return func(duration time.Duration) error {
		time.Sleep(time.Second * time.Duration(t))
		return nil
	}
}

func runPShutDown() {
	f1 := ShutDownFunc(ShutDownMaker(3))
	f2 := ShutDownFunc(ShutDownMaker(4))
	f3 := ShutDownFunc(ShutDownMaker(5))

	r := ConcurrentShutdown(time.Second*10, f1, f2, f3)
	log.Println(r)
}
