package task

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	. "viking/goplus/core"
	"viking/goplus/runtime"
)

const (
	INIT		int = iota
	RUNNING
	FINISH
)


type Task struct {
	function			reflect.Value
	arguments			[]reflect.Value
	state				int
	result				interface{}
	error				interface{}
	wg					sync.WaitGroup
	stack				runtime.Stack
}

func (h *Task) State() int { return h.state }

func (h *Task) Error() interface{} { return h.error }

func (h *Task) Result() interface{} { return h.result }

func (h *Task) Finish() bool { return h.state == FINISH}

func (h *Task) StateDesc() string {
	switch h.state {
	case INIT: return "Init"
	case RUNNING: return "Running"
	case FINISH: return "Finish"
	default: return "Invalid"
	}
}

func Create(f interface{}, args... interface{}) *Task {
	if f == nil { panic(errors.New("Can not create task with nil function")) }
	vf := reflect.ValueOf(f)
	if vf.Kind() != reflect.Func { panic(errors.New(fmt.Sprintf("Can not create task with invalid function type %v", vf.Kind()))) }
	vargs := make([]reflect.Value, len(args))
	for i, a := range (args){ vargs[i] = reflect.ValueOf(a) }
	return &Task{vf, vargs,INIT, nil, nil, sync.WaitGroup{}, runtime.Trace(3)}
}

func (h *Task) Start(){
	if h.state != INIT { panic(errors.New(fmt.Sprintf("Task can not start, current state is %v", h.StateDesc()))) }
	h.state = RUNNING
	h.wg.Add(1)
	go func() {
		defer func(){
			h.wg.Done()
			h.state = FINISH
		}()
		defer CatchCallback(func(info *PanicInfo){
			h.error = info.Error()
			log.Fatal(strings.Join([]string{
				"A task error occurred as below",
				fmt.Sprintf("%v", h.stack),
			}, "\n"))
		})
		result := h.function.Call(h.arguments)
		h.result = result
	}()
}

func (h *Task) Wait(){
	if h.state == INIT { panic(errors.New("Can't wait for a task that hasn't started")) }
	if h.state == FINISH { return }
	h.wg.Wait()
}

func (h *Task) WaitResult() interface{}{
	h.Wait()
	return h.result
}