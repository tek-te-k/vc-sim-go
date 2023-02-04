package simulation

import (
	"crypto/rand"
	"log"
	"math/big"
	"vc-sim-go/models"
	"vc-sim-go/state"
)

type Result struct {
	TotalCycle int
}

type Simulator struct {
	Workers []*models.Worker
	Jobs    []*models.Job
	Config  Config
	Result  Result
}

func NewSimulator(workers []*models.Worker, jobs []*models.Job, config Config) *Simulator {
	return &Simulator{
		Workers:        workers,
		Jobs:           jobs,
		Config:         config,
		Result: Result{
			TotalCycle: 0,
		},
	}
}

func (s *Simulator) SetWorkersState() {
	for i := range s.Workers {
		if float64(i) < float64(len(s.Workers))*s.Config.InitialJoiningRate {
			s.Workers[i].State = state.AvailableWorkerState
			continue
		}
		s.Workers[i].State = state.UnavailableWorkerState
	}
}

func (s *Simulator) SetWorkersParticipationRate() {
	for i := range s.Workers {
		s.Workers[i].SecessionRate = s.Config.SecessionRate
		s.Workers[i].JoiningRate = s.Config.JoiningRate
	}
}

func (s *Simulator) areAllJobsFinished() bool {
	for _, job := range s.Jobs {
		if job.State != state.FinishedJobState {
			return false
		}
	}
	return true
}

func (s *Simulator) Simulate() int {
	cycle := 0
	log.Println(s.Jobs[1])
	for !s.areAllJobsFinished() {
		s.assignJobs()
		s.workerSecessionEvent()
		s.finishJobs()
		s.workerJoinEvent()
		cycle++
		// log.Println(s.Jobs[1])
	}
	return cycle
}

func (s *Simulator) assignJobs() {
	label:
	for _, job := range s.Jobs {
		if job.State != state.UnallocatedJobState {
			continue
		}
		for _, subjob := range job.Subjobs {
			if subjob.State != state.UnallocatedSubjobState {
				continue
			}
			for i := 0; i < s.Config.Redundancy; i++ {
				for _, worker := range s.Workers {
					if worker.State != state.AvailableWorkerState {
						continue
					}
					worker.State = state.RunningWorkerState
					subjob.AssignedWorker = append(subjob.AssignedWorker, worker)
					subjob.State = state.ProcessingSubjobState
					break
				}
			}
			if subjob.State == state.UnallocatedSubjobState {
				break label
			}
		}
		job.State = state.ProcessingJobState
	}
}


func (s *Simulator) workerSecessionEvent() {
	for _, job := range s.Jobs {
		for _, subjob := range job.Subjobs {
			if subjob.State != state.ProcessingSubjobState {
				continue
			}
			// subjobFailedRate := 1 - math.Pow((1 - s.Config.SecessionRate), len(subjob.AssignedWorker))
			// n, err := rand.Int(rand.Reader, big.NewInt(100))
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// if n.Int64() < int64(subjobFailedRate*100) {
			// 	job.Failed()
			// 	continue
			// }
			for _, aw := range subjob.AssignedWorker {
				n, err := rand.Int(rand.Reader, big.NewInt(100))
				if err != nil {
					log.Fatal(err)
				}
				if n.Int64() < int64(aw.SecessionRate*100) {
					err := aw.Secession()
					if err != nil {
						log.Fatal(err)
					}
					if job.State == state.UnallocatedJobState {
						job.Failed()
					}
				}
			}
		}
	}
}

func (s *Simulator) finishJobs() {
	for _, job := range s.Jobs {
		if job.State != state.ProcessingJobState {
			continue
		}
		for _, subjob := range job.Subjobs {
			if subjob.State != state.ProcessingSubjobState {
				continue
			}
			for _, aw := range subjob.AssignedWorker {
				aw.State = state.AvailableWorkerState
			}
			subjob.State = state.FinishedSubjobState
		}
		job.State = state.FinishedJobState
	}
}

func (s *Simulator) workerJoinEvent() {
	for _, worker := range s.Workers {
		if worker.State == state.AvailableWorkerState || worker.State == state.RunningWorkerState {
			continue
		}
		n, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			log.Fatal(err)
		}
		if n.Int64() < int64(worker.JoiningRate*100) {
			err := worker.Join()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
