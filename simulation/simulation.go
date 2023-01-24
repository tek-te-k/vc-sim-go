package simulation

import (
	"vc-sim-go/models"
	"vc-sim-go/state"
)

type Simulator struct {
	Workers []*models.Worker
	Jobs []*models.Job
	ParallelismNum int
}

func NewSimulator(workers []*models.Worker, jobs []*models.Job, parallelismNum int) *Simulator {
	return &Simulator{
		Workers: workers,
		Jobs: jobs,
		ParallelismNum: parallelismNum,
	}
}

func (s Simulator) SetWorkersState(joiningRate float64) {
	for i := range s.Workers {
		if (float64(i) <= float64(len(s.Workers)) * joiningRate) {
			s.Workers[i].State = state.AvailableWorkerState
			continue
		}
		s.Workers[i].State = state.UnavailableWorkerState
	}
}

func (s Simulator) SetWorkersParticipationRate(dropoutRate float64, joiningRate float64) {
	for i := range s.Workers {
		s.Workers[i].DropoutRate = dropoutRate
		s.Workers[i].JoiningRate = joiningRate
	}
}

func (s Simulator) Simulate() {
	finishedJobCount := 0
	for finishedJobCount != len(s.Jobs) {
		for i := 0; i < s.ParallelismNum; i++ {
			dist := 0
			fail := 0
			s.assignJobs()
		}
	}
}

func (s Simulator) assignJobs() {
	subjobNum := len(s.Jobs) * s.ParallelismNum
	for i := 0; i < subjobNum; i++ {
		if s.Jobs[i].State == state.UnallocatedJobState {
			for j := range s.Workers {
				if s.Workers[j].State == state.AvailableWorkerState {
				}
			}
		}
	}
}