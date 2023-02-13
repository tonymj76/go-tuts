package runner

import (
	"fmt"
	"testing"
	"time"
)

func does(c int) func(int) {
	return func(i int) {
		fmt.Printf("multipling %d with %d is %d \n", i, c, c*i)
	}
}

func TestRunner_Start(t *testing.T) {
	backTask := New(time.Minute * 1)
	var do []func(int)
	for i := 1; i < 20; i++ {
		do = append(do, does(i))
	}

	backTask.AddTasks(do...)
	if len(backTask.tasks) != 19 {
		t.Error("length don't match task\n")
	}
	if err := backTask.Start(); err != nil {
		t.Errorf("task failed at %v\n", err)
	}
}
