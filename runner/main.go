package main

import (
	"errors"
	"os"
	"time"
)

var (
	ErrInterrupt = errors.New("received Interrupt")
	ErrTimeout   = errors.New("received timeout")
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

func (r Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
