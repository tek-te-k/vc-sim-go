package models

import (
	"vc-sim-go/state"
)

type Subjob struct {
	ID int
	State state.SubjobState
	AssignedWorker *Worker
}

func NewSubjob(id int, state state.SubjobState, assignedWorker *Worker) *Subjob {
	return &Subjob{
		ID: id,
		State: state,
		AssignedWorker: assignedWorker,
	}
}
