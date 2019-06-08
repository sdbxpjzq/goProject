package main

import (
	"ysqi/app/crawler_zhenai/engine"
	"ysqi/app/crawler_zhenai/persist"
	"ysqi/app/crawler_zhenai/scheduler"
	"ysqi/app/crawler_zhenai/zhenai/parser"
)

func main() {

	ItemChan, err := persist.ItemSaver("datint_profile_1")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    ItemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// 简单并发版本
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 10	,
	//}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// 简单无并发版本
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

}
