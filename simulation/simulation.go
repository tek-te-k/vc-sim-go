package simulation

import (
	"crypto/rand"
	"log"
	"math/big"
	"vc-sim-go/models"
	"vc-sim-go/state"
)

type Simulator struct {
	Workers        []*models.Worker
	Jobs           []*models.Job
	ParallelismNum int
}

func NewSimulator(workers []*models.Worker, jobs []*models.Job, parallelismNum int) *Simulator {
	return &Simulator{
		Workers:        workers,
		Jobs:           jobs,
		ParallelismNum: parallelismNum,
	}
}

func (s Simulator) SetWorkersState(joiningRate float64) {
	for i := range s.Workers {
		if float64(i) <= float64(len(s.Workers))*joiningRate {
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

func (s Simulator) areAllJobsFinished() bool {
	for i := range s.Jobs {
		if s.Jobs[i].State != state.FinishedJobState {
			return false
		}
	}
	return true
}

func (s Simulator) Simulate() {
	for s.areAllJobsFinished() {
		for i := 0; i < s.ParallelismNum; i++ {
			s.assignJobs()
			s.participationEvent()
			s.dropoffJobs()
		}
	}
}

func (s Simulator) assignJobs() {
	subjobNum := len(s.Jobs) * s.ParallelismNum
	for i := 0; i < subjobNum; i++ {
		if s.Jobs[i].State == state.UnallocatedJobState {
			for j := range s.Workers {
				if s.Workers[j].State == state.AvailableWorkerState {
					if s.Workers[j].AssignedJob != nil || s.Jobs[j].AssignedWorker != nil {
						log.Fatal("Worker or Job is already assigned")
					}
					s.Workers[j].State = state.RunningWorkerState
					s.Jobs[j].State = state.ProcessingJobState
					s.Workers[j].AssignedJob = s.Jobs[j]
					s.Jobs[j].AssignedWorker = s.Workers[j]
				}
			}
		}
	}
}

func (s Simulator) participationEvent() {
	for i := range s.Workers {
		if s.Workers[i].State == state.RunningWorkerState || s.Workers[i].State == state.AvailableWorkerState {
			n, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				log.Fatal(err)
			}
			if n.Int64() <= int64(s.Workers[i].DropoutRate*100) {
				err := s.Workers[i].Dropout()
				if err != nil {
					log.Fatal(err)
				}
			}
		} else if s.Workers[i].State == state.UnavailableWorkerState {
			n, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				log.Fatal(err)
			}
			if n.Int64() <= int64(s.Workers[i].JoiningRate*100) {
				s.Workers[i].Join()
			}
		}
	}
}

func (s Simulator) dropoffJobs() {
	for i := range s.Workers {
		if s.Workers[i].State == state.RunningWorkerState {
			if s.Workers[i].State != state.RunningWorkerState {
				log.Fatal("Worker is not running")
			}
			job := s.Workers[i].AssignedJob
			s.Workers[i].State = state.AvailableWorkerState
			job.State = state.FinishedJobState
			s.Workers[i].AssignedJob = nil
			job.AssignedWorker = nil
		}
	}
}
