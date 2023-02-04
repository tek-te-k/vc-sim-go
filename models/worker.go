package models

import (
	"errors"
	"vc-sim-go/state"
)

type Worker struct {
	ID          int
	State       state.WorkerState
	SecessionRate float64
	JoiningRate float64
}

func NewWorker(id int, state state.WorkerState) *Worker {
	return &Worker{
		ID:          id,
		State:       state,
	}
}

func (w *Worker) Secession() error {
	if w.State != state.RunningWorkerState && w.State != state.AvailableWorkerState {
		return errors.New("Worker is not available")
	}
	w.State = state.UnavailableWorkerState
	return nil
	// switch w.State {
	// case state.RunningWorkerState:
	// 	job := w.AssignedJob
	// 	w.State = state.UnavailableWorkerState
	// 	job.State = state.UnallocatedJobState
	// 	w.AssignedJob = nil
	// 	job.AssignedWorker = nil
	// case state.AvailableWorkerState:
	// 	w.State = state.UnavailableWorkerState
	// default:
	// 	return errors.New("Worker is not available")
	// }                                                                                                                     
	// return nil
}

func (w *Worker) Join() error {
	if w.State != state.UnavailableWorkerState {
		return errors.New("Worker is not available")
	}
	w.State = state.AvailableWorkerState
	return nil
}
