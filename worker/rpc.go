/************************************************************
** @Description: worker
** @Author: haodaquan
** @Date:   2018-03-23 23:16
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 23:16
*************************************************************/
package worker

import "github.com/george518/crawler/engine"

type CrawlerService struct{}

func (CrawlerService) Process(
	req Request,
	result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return nil
	}
	*result = SerializeResult(engineResult)
	return nil
}
