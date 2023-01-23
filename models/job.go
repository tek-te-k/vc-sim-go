package models

import (
	"vc-sim-go/state"
)

type Job struct {
	ID             int
	GroupID        int
	// TODO state になんの種類があるか確認し、int で管理するのをやめてみる
	State          state.JobState
	// TODO Worker との二重結合を解消し、親子関係をはっきりさせる
	AssignedWorker *Worker
	IsAssigned     bool
}
