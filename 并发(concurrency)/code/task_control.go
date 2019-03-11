package main

import(
	"fmt"
	"errors"
	"time"
	"os/signal"
	"os"
)

var ErrTimeOut=errors.New("Running timeout")
var ErrInterrupt=errors.New("interrupt program")

type Runner struct{
	tasks []func(int)
	complete chan error
	timeout <-chan time.Time
	interrupt chan os.Signal
}

func NewRuner(tm time.Duration) *Runner{
	return &Runner{
		complete: make(chan error),
		timeout: time.After(tm),
		interrupt:	make(chan os.Signal,1),
	}
}

func (r *Runner) Add(tasks ...func(int)){
	r.tasks=append(r.tasks,tasks...)
}

func (r *Runner) run() error{
	for id,task:=range r.tasks{
		if r.inInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) inInterrupt() bool{
	select{
	case<-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
func (r *Runner) Start() error{
	signal.Notify(r.interrupt,os.Interrupt)

	go func(){
		r.complete<- r.run()
	}()

	select{
	case err:=<-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func main(){
	fmt.Println("...beginning run...")
	timeout:=3*time.Second
	run:=NewRuner(timeout)

	run.Add(createTask(),createTask(),createTask())

	if err:=run.Start();err!=nil{
		switch err{
		case ErrTimeOut:
			fmt.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			fmt.Println(err)
			os.Exit(2)
		}
	}
	fmt.Println("task finish")
}

func createTask() func(int){
	return func(id int){
		fmt.Printf("beginning Task:%d\n",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}