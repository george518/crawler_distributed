/************************************************************
** @Description: worker
** @Author: haodaquan
** @Date:   2018-03-23 21:26
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-03-23 21:26
*************************************************************/
package worker

import (
	"fmt"

	"log"

	"github.com/george518/crawler/engine"
	"github.com/george518/crawler/zhenai/parser"
	"github.com/george518/crawler_distributed/config"
	"github.com/pkg/errors"
)

type SerializedParse struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParse
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParse{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(
	r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserialize request %v", err)
			fmt.Print(err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(
	p SerializedParse) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if name, ok := p.Args.(string); ok {
			return parser.NewProfileParser(name), nil
		} else {
			return nil, fmt.Errorf("invalid args:%v", p.Args)
		}
	default:
		return nil, errors.New("unkown parser name")
	}
}
