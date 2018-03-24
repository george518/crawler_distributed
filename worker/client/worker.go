/************************************************************
** @Description: client
** @Author: haodaquan
** @Date:   2018-03-24 20:23
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-24 20:23
*************************************************************/
package client

import (
	"net/rpc"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler_distributed/config"
	"github.com/george518/crawler_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var result worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlerServiceRPC, sReq, &result)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeResult(result), nil
	}
}
