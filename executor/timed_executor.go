package executor

import "github.com/peter9207/dbcompare/queries"

type timedExecutor struct {
	Ratio   float64
	Seconds int64
	db      *queries.Runner
}

func NewTimedExecutor(ratio float64, seconds int64, db *queries.Runner) (e *timedExecutor) {

	e = &timedExecutor{
		Ratio:   ratio,
		Seconds: seconds,
		db:      db,
	}
	return
}

func (e *timedExecutor) runRead() (err error) {
	_, err = e.db.PerformRead()
	return
}

func (e *timedExecutor) runWrite() (err error) {

	_, err = e.db.PerformWrite()
	return
}

func (e *timedExecutor) Run(workers int64) (errCh chan error) {

	return
}
