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
	Run()
	WorkerReady(chan Request)
}

func (c *Concurrent) Run(seeds ...Request) {
	// 1 创建channel of out
	out := make(chan ParseResult)
	c.Scheduler.Run()
	// 2 createWorker
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(out, c.Scheduler)
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

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
