package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"vc-sim-go/models"
	"vc-sim-go/state"

	"github.com/joho/godotenv"
)

func initWorkers(workerCount int) ([]*models.Worker, error) {
	workers := make([]*models.Worker, workerCount)
	for i := range workers {
		workers[i] = models.NewWorker(i, state.UnavailableWorkerState, nil)
	}
	return workers, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	workerLimit, err := strconv.Atoi(os.Getenv("WORKER_LIMIT"))
	if err != nil {
		log.Fatal("Error loading workerLimit")
	}
	jobLimit, err := strconv.Atoi(os.Getenv("JOB_LIMIT"))
	if err != nil {
		log.Fatal("Error loading jobLimit")
	}
	joiningRate, err := strconv.ParseFloat(os.Getenv("JOINING_RATE"), 32)
	if err != nil {
		log.Fatal("Error loading joiningRate")
	}
	dropoutRate, err := strconv.ParseFloat(os.Getenv("DROPOUT_RATE"), 32)
	if err != nil {
		log.Fatal("Error loading dropoutRate")
	}
	initialJoiningRate, err := strconv.ParseFloat(os.Getenv("INITIAL_JOINING_RATE"), 32)
	if err != nil {
		log.Fatal("Error loading initialJoiningRate")
	}
	loopCount, err := strconv.Atoi(os.Getenv("LOOP_COUNT"))
	if err != nil {
		log.Fatal("Error loading loopCount")
	}
	log.Println(fmt.Sprintf(`
ワーカ数: %d,
ジョブ数: %d,
参加率: %.3f,
離脱率: %.3f,
初期のワーカの参加率: %.3f,
並列数: %d`,
		workerLimit,
		jobLimit,
		joiningRate,
		dropoutRate,
		initialJoiningRate,
		loopCount,
	))
	workers, err := initWorkers(workerLimit)
	if err != nil {
		log.Fatal("init failed")
	}
	log.Println(workers)
}
