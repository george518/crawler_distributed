/************************************************************
** @Description: server
** @Author: haodaquan
** @Date:   2018-03-23 23:26
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 23:26
*************************************************************/
package main

import (
	"fmt"

	"flag"

	"github.com/george518/crawler_distributed/rpcsupport"
	"github.com/george518/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlerService{})
	if err != nil {
		panic(err)
	}

}
