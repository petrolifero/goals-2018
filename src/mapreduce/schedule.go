package mapreduce

import (
	"fmt"
	"sync"
)

const (
	Idle = iota
	InProgress
	Completed
)

type TasksState struct {
	State int
	Worker Worker
}

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}
	var waitgroup sync.Waitgroup
	tasksState := make(TasksState,ntasks)
	for i:=0 ; i<ntasks; i++ {
		tasksState[i].State=Idle
	}
	waitgroup.Add(ntasks)
	for i:=0; i<ntasks; i++ {
		go func() {
			reply := 0
			call("localhost","Worker.DoTask",,&reply)
			waitgroup.Done()
		}
	}
	waitgroup.Wait()
	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//
	fmt.Printf("Schedule: %v phase done\n", phase)
}
