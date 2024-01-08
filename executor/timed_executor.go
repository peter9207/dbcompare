package executor

import (
	"fmt"
	"sync"
	"time"

	"github.com/peter9207/dbcompare/queries"
)

type timedExecutor struct {
	Seconds int
	db      queries.Runner
}

func NewTimedExecutor(seconds int64, db queries.Runner) (e *timedExecutor) {

	e = &timedExecutor{
		Seconds: int(seconds),
		db:      db,
	}
	return
}

func (e *timedExecutor) runRead(readCountCh chan bool, endCh chan bool) (err error) {
	for {
		err = e.db.PerformRead()
		if err != nil {
			fmt.Println("encoutered Error", err.Error())
		}
		readCountCh <- true
		select {
		case <-endCh:
			fmt.Println("stopping read worker")
			return
		default:
		}
	}

}

func (e *timedExecutor) runWrite(countCh chan bool, endCh chan bool) (err error) {

	for {
		err = e.db.PerformWrite()
		if err != nil {
			fmt.Println("encoutered Error", err.Error())
		}
		countCh <- true

		select {
		case <-endCh:
			fmt.Println("stopping write worker")
			return
		default:
		}
	}

}

type RunResult struct {
	ReadCount  int64 `json: "read"`
	WriteCount int64 `json: "write"`
}

func (e *timedExecutor) Run(readWorkers, writeWorkers int64) (result RunResult, errCh chan error) {

	var wg sync.WaitGroup
	result = RunResult{}

	var stopCh chan bool
	stopCh = make(chan bool, readWorkers+writeWorkers)

	fmt.Println("workers", readWorkers, writeWorkers)

	readCountCh := make(chan bool, 100)
	writeCountCh := make(chan bool, 100)

	var readCount int64
	var writeCount int64

	go func() {
		for {
			select {
			case <-readCountCh:
				readCount++
			case <-writeCountCh:
				writeCount++
			}
		}
	}()

	for i := int64(0); i < readWorkers; i++ {

		wg.Add(1)
		go func() {
			e.runRead(readCountCh, stopCh)
			wg.Done()
		}()

	}

	for i := int64(0); i < writeWorkers; i++ {
		wg.Add(1)
		go func() {
			e.runWrite(writeCountCh, stopCh)
			wg.Done()
		}()
	}

	go func(signalCh chan bool) {
		// time.Sleep(2 * time.Second)
		time.Sleep(time.Duration(e.Seconds) * time.Second)
		for i := int64(0); i < writeWorkers+readWorkers; i++ {
			signalCh <- true
		}

	}(stopCh)

	wg.Wait()

	result.ReadCount = readCount
	result.WriteCount = writeCount

	return
}
