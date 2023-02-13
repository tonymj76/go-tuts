package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeOut      = errors.New("time out")
	ErrInterrupted  = errors.New("task is interrupted")
	ErrNotCompleted = errors.New("error while run task")
)

type runner struct {
	completed   chan error
	interrupted chan os.Signal
	timeOut     <-chan time.Time
	tasks       []func(int)
}

func New(elapse time.Duration) *runner {
	return &runner{
		completed:   make(chan error),
		interrupted: make(chan os.Signal, 1),
		timeOut:     time.After(elapse),
	}
}

func (r *runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *runner) Start() error {
	signal.Notify(r.interrupted, os.Interrupt)
	go func() {
		r.completed <- r.run()
	}()
	select {
	case err := <-r.completed:
		if err != nil {
			return fmt.Errorf("%v: %w", err, ErrNotCompleted)
		}
		return err
	case <-r.timeOut:
		return ErrTimeOut
	}
}

func (r *runner) run() error {
	for i, task := range r.tasks {
		if r.gotInterrupted() {
			return ErrInterrupted
		}
		task(i)
	}
	return nil
}

func (r *runner) gotInterrupted() bool {
	select {
	case c := <-r.interrupted:
		fmt.Println(c)
		signal.Stop(r.interrupted)
		return true
	default:
		return false
	}
}
