package models

import (
	"vc-sim-go/state"
)

type Job struct {
	ID      int
	State   state.JobState
	Subjobs []*Subjob
}

func NewJob(id int, state state.JobState, subjobs []*Subjob) *Job {
	return &Job{
		ID: id,
		State: state,
		Subjobs: subjobs,
	}
}
