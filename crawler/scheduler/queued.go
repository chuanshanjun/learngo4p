package scheduler

import "chuanshan.github.com/learngo4p/crawler/engine"

type QueuedScheduler struct {
	workerChan  chan chan engine.Request
	requestChan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) ConfigureMaterWorkerChan(r chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	// 1 初始化chan
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		// 2 初始化队列
		var workerQ []chan engine.Request
		var requestQ []engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
