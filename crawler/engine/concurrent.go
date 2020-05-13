package engine

import "chuanshan.github.com/learngo4p/crawler/model"

type Concurrent struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	Submit(Request)
	ConfigureMaterWorkerChan(chan Request)
	Run()
	ReadyNotifier
	WorkerChan() chan Request
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var duplicateUrls = make(map[string]bool)

func (c *Concurrent) Run(seeds ...Request) {
	// 1 创建channel of out
	out := make(chan ParseResult)
	c.Scheduler.Run()
	// 2 createWorker
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}
	// 4 把种子装载进去
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}
	// 5 从out中取出值
	//itemCount := 0
	for {
		// 我要能成功的从out中收到，那么我要将request成功的送走
		parseResult := <-out
		// 4.1 打印item
		for _, item := range parseResult.Items {
			//log.Printf("Got item %v\n", item)
			if _, ok := item.Payload.(model.Profile); ok {
				go func() {
					c.ItemChan <- item
				}()
			}
			//itemCount++
		}
		// 4.2 继续将request放到scheduler
		for _, r := range parseResult.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			c.Scheduler.Submit(r)
		}
	}
}

var visitedUrls = make(map[string]bool)

// 教程的写法
func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}

func isDuplicate(url string) bool {
	_, ok := duplicateUrls[url]
	if ok {
		//log.Printf("This url is exist %s", url)
		return true
	} else {
		duplicateUrls[url] = true
		return false
	}
}

func createWorker(in chan Request, out chan ParseResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
