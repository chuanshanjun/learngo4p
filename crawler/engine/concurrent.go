package engine

type Concurrent struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

// 函数也可以定义为一种类型
type Processor func(Request) (ParseResult, error)

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

func (e *Concurrent) Run(seeds ...Request) {
	// 1 创建channel of out
	out := make(chan ParseResult)
	e.Scheduler.Run()
	// 2 createWorker
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	// 4 把种子装载进去
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	// 5 从out中取出值
	//itemCount := 0
	for {
		// 我要能成功的从out中收到，那么我要将request成功的送走
		parseResult := <-out
		// 4.1 打印item
		for _, item := range parseResult.Items {
			//log.Printf("Got item %v\n", item)

			if item.Payload != nil {
				go func() {
					e.ItemChan <- item
				}()
			}

			//if _, ok := item.Payload.(model.Profile); ok {
			//	go func() {
			//		e.ItemChan <- item
			//	}()
			//}
			//itemCount++
		}
		// 4.2 继续将request放到scheduler
		for _, r := range parseResult.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
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

func (e *Concurrent) createWorker(in chan Request, out chan ParseResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in)
			request := <-in
			// 此处的Worker将其拆成RPC请求
			//result, err := Worker(request)
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
