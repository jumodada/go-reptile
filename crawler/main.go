package main

import (
	"fmt"
	"strconv"
)

func main() {
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 100,
	//}
	//
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println(i)
	}
}
