/************************************************************
** @Description: crawler_distributed
** @Author: haodaquan
** @Date:   2018-03-23 21:07
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 21:07
*************************************************************/
package main

import (
	"net/rpc"

	"log"

	"flag"

	"strings"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/scheduler"
	"github.com/george518/crawler/zhenai/parser"
	"github.com/george518/crawler_distributed/config"
	Itemsaver "github.com/george518/crawler_distributed/persist/client"
	"github.com/george518/crawler_distributed/rpcsupport"
	Worker "github.com/george518/crawler_distributed/worker/client"
)

var (
	itemServerHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts( comma separated)")
)

func main() {
	flag.Parse()
	//itemChan, err := persist.ItemServer("dating_profile")
	itemChan, err := Itemsaver.ItemServer(*itemServerHost)
	if err != nil {
		panic(err)
	}
	//TODO 验证合法性
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := Worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(
	hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s:%v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out

}
