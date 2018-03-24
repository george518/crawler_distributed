/************************************************************
** @Description: persist
** @Author: haodaquan
** @Date:   2018-03-22 23:30
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-22 23:30
*************************************************************/
package persist

import (
	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/persist"

	"log"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(item, s.Client, s.Index)
	log.Printf("item save %v", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Panicf("item :%v,err:%v", item, err)
	}
	return err

}
