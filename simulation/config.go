package simulation

type Config struct {
	WorkerLimit int
	JobLimit int
	JoiningRate float64
	DropoutRate float64
	InitialJoiningRate float64
	LoopCount int
	ParallelismNum int
	Redundancy int
}

func NewConfig(workerLimit int, jobLimit int, joiningRate float64, dropoutRate float64, initialJoiningRate float64, loopCount int, parallelismNum int, redundancy int) *Config {
	return &Config{
		WorkerLimit: workerLimit,
		JobLimit: jobLimit,
		JoiningRate: joiningRate,
		DropoutRate: dropoutRate,
		InitialJoiningRate: initialJoiningRate,
		LoopCount: loopCount,
		ParallelismNum: parallelismNum,
		Redundancy: redundancy,
	}
}
