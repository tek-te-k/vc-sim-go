package models

import (
	"vc-sim-go/state"
)

type Subjob struct {
	ID int
	State state.SubjobState
	AssignedWorker []*Worker
}

func NewSubjob(id int, state state.SubjobState) *Subjob {
	return &Subjob{
		ID: id,
		State: state,
		AssignedWorker: make([]*Worker, 0),
	}
}
