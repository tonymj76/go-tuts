package main

import (
	"fmt"
	"time"
	"context"
)

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
	fmt.Println(msg)
}

func main() {
	//Creating a background context which is the root of the context which does nothing 
	//Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. 
	//It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.
	bctx := context.Background()

	// //Now creating a child of the base context 
	// ctx, cancel := context.WithCancel(bctx)

	// //Afterfunc will wait till a sec then call cancel funtion that will terminate every process of the base parent children
	// time.AfterFunc(time.Second, cancel)

	ctx, cancel := context.WithTimeout(bctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}