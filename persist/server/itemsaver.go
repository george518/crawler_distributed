/************************************************************
** @Description: server
** @Author: haodaquan
** @Date:   2018-03-22 23:39
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-22 23:39
*************************************************************/
package main

import (
	"fmt"

	"flag"

	"github.com/george518/crawler_distributed/config"
	"github.com/george518/crawler_distributed/persist"
	"github.com/george518/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	err := ServeRpc(fmt.Sprintf(":%d",
		*port),
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func ServeRpc(host, index string) error {
	Client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: Client,
		Index:  index,
	})

}
