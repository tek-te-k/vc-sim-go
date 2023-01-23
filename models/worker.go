package models

import (
	"vc-sim-go/state"
)

type Worker struct {
	ID          int
	State       state.WorkerState
	AssignedJob *Job
	DropoutRate float64
	JoiningRate float64
}

func NewWorker(id int, state state.WorkerState, assignedJob *Job) *Worker{
	return &Worker{
		ID: id,
		State: state,
		AssignedJob: assignedJob,
	}
}