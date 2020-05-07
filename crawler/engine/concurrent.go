package engine

import (
	"log"
)

type Concurrent struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMaterWorkerChan(chan Request)
}

func (c *Concurrent) Run(seeds ...Request) {
	// 1 创建channel of in&out
	in := make(chan Request)
	out := make(chan ParseResult)
	// 2 准备好workerChan
	c.Scheduler.ConfigureMaterWorkerChan(in)
	// 3 createWorker
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}
	// 4 把种子装载进去
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}
	// 5 从out中取出值
	itemCount := 0
	for {
		// 我要能成功的从out中收到，那么我要将request成功的送走
		parseResult := <-out
		// 4.1 打印item
		for _, item := range parseResult.Items {
			log.Printf("Got item %v\n", item)
			itemCount++
		}
		// 4.2 继续将request放到scheduler
		for _, r := range parseResult.Requests {
			c.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
