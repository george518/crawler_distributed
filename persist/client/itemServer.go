/************************************************************
** @Description: client
** @Author: haodaquan
** @Date:   2018-03-23 20:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 20:43
*************************************************************/
package client

import (
	"log"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler_distributed/config"
	"github.com/george518/crawler_distributed/rpcsupport"
)

func ItemServer(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item saver: got item #%d:%v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call(config.ItemServerRpc, item, &result)
			if err != nil {
				log.Printf("err #%d %v", itemCount, err)
			}
		}

	}()
	return out, nil
}
