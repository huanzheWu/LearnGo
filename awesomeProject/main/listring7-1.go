package main

import (
	"os"
	"time"
	"fmt"
	"errors"
	"os/signal"
)

type Runner struct {
	interrupt chan os.Signal
	complete chan error
	timeout <- chan time.Time
	task []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")
func New( d time.Duration) *Runner{
	return &Runner{
		interrupt:make(chan os.Signal,1),
		complete:make(chan error),
		timeout:time.After(d),
	}
}

func (r* Runner) Add(tasks ...func(int)){
	r.task = append(r.task,tasks...)
}

func (r *Runner) Start()error{
	signal.Notify(r.interrupt,os.Interrupt)
	go func(){
		r.complete<-r.run()
	}()
}

func main() {

}
