package executor

type timedExecutor struct {
	Ratio   float64
	Seconds int64
}

func NewTimedExecutor(ratio float64, seconds int64) (e *timedExecutor) {

	e = &timedExecutor{
		Ratio:   ratio,
		Seconds: seconds,
	}
	return
}

func (e *timedExecutor) Run() (errCh chan error) {

	return
}
