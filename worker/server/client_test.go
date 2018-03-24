/************************************************************
** @Description: main
** @Author: haodaquan
** @Date:   2018-03-23 23:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 23:30
*************************************************************/
package main

import (
	"testing"

	"time"

	"fmt"

	"github.com/george518/crawler_distributed/config"
	"github.com/george518/crawler_distributed/rpcsupport"
	"github.com/george518/crawler_distributed/worker"
)

func TestCrawlerService(t *testing.T) {
	const host = ":9100"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1254739352",
		Parser: worker.SerializedParse{
			Name: config.ParseProfile,
			Args: "等待",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlerServiceRPC, req, &result)
	fmt.Print(result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Print(result)
	}
}
