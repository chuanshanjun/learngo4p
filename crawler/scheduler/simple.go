package scheduler

import "chuanshan.github.com/learngo4p/crawler/engine"

type Simple struct {
	WorkChan chan engine.Request
}

func (s *Simple) Submit(request engine.Request) {
	// 15-2 简单调度器12:25讲解，这边为何要用go func
	// request送给worker送成功的前提是必须要有worker在那等待，
	// 此时我10个worker都在做其他事情，都不空
	// chan的发送一定是要两方都要存在的，一方不在它就会等待
	// 所以成功的前提是有worker在等这个request，worker要把手头的活做完才能返回(也就是把request及item成功的送给engine)
	// 表示他里面的request要通过scheduler成功的将request送给其他的worker
	//s.WorkChan <- request
	go func() {
		s.WorkChan <- request
	}()
}

// 配置chan
func (s *Simple) ConfigureMaterWorkerChan(c chan engine.Request) {
	s.WorkChan = c
}
