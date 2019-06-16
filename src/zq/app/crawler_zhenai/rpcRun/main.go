package main

import (
	"zq/app/crawler_zhenai/engine"
	"zq/app/crawler_zhenai/rpcRun/client"
	"zq/app/crawler_zhenai/scheduler"
	"zq/app/crawler_zhenai/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"

	itemChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
