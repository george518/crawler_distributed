/************************************************************
** @Description: main
** @Author: haodaquan
** @Date:   2018-03-22 23:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-22 23:43
*************************************************************/
package main

import (
	"testing"

	"time"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/model"
	"github.com/george518/crawler_distributed/config"
	"github.com/george518/crawler_distributed/rpcsupport"
)

func TestServeRpc(t *testing.T) {
	const host = ":1234"
	go ServeRpc(host, "test1")
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(client)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108835456",
		Id:   "108835456",
		Type: "zhenai",
		PayLoad: model.Profile{
			Name:       "缱绻",
			Gender:     "女",
			Age:        28,
			Height:     165,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "--",
			Hukou:      "江苏徐州",
			Xingzuo:    "狮子座",
			House:      "--",
			Car:        "未购车",
			City:       "江苏徐州",
		},
	}
	result := ""
	err = client.Call(config.ItemServerRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s,err:%s", result, err)
	}

}
