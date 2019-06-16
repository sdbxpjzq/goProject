package engine

//接口
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//接口
type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler //Sheduler
	WorkerCount int       //worker的数量
	ItemChan    chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//worker公用一个in，out
	//in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		createWorker2(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//参数seeds的request，要分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			// 存储
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker2(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
