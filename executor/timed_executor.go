package executor

import (
	"sync"

	"github.com/peter9207/dbcompare/queries"
)

type timedExecutor struct {
	Seconds int64
	db      queries.Runner
}

func NewTimedExecutor(seconds int64, db queries.Runner) (e *timedExecutor) {

	e = &timedExecutor{
		Seconds: seconds,
		db:      db,
	}
	return
}

func (e *timedExecutor) runRead(endCh chan bool) (err error) {
	for {
		err = e.db.PerformRead()
		select {
		case <-endCh:
			return
		default:

		}
	}

}

func (e *timedExecutor) runWrite(endCh chan bool) (err error) {

	for {
		err = e.db.PerformWrite()
		select {
		case <-endCh:
			return
		default:

		}
	}

}

func (e *timedExecutor) Run(readWorkers, writeWorkers int64) (errCh chan error) {

	var wg sync.WaitGroup

	var stopCh chan bool

	for i := int64(0); i < readWorkers; i++ {

		go func() {
			wg.Add(1)
			e.runRead(stopCh)
			wg.Done()
		}()

	}

	for i := int64(0); i < writeWorkers; i++ {

		go func() {
			wg.Add(1)
			e.runWrite(stopCh)
			wg.Done()
		}()

	}

	wg.Wait()

	return
}
