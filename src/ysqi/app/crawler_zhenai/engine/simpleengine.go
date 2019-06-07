package engine

import (
	"log"
)

// 简单无并发版本

type SimpleEngine struct {
}

// 程序入口
func (e *SimpleEngine) Run(seed ...Request) {
	var requests []Request
	for _, value := range seed {
		requests = append(requests, value)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, _ := worker(r)
		for _, items := range parseResult.Items {
			log.Println("Got item", items)
		}
		requests = append(requests, parseResult.Requests...)

	}

}
