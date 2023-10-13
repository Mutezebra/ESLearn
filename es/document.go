package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-ElasticSearch/es/esmodel"
	"log"
)

func CreateDoc(model esmodel.ESModels) (*elastic.IndexResponse, error) {
	// 第一个index()可以理解为对index的操作
	resp, err := ESClient.
		Index().Index(model.Index()).
		BodyJson(&model).
		Do(context.TODO())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("create doc success ", resp)
	return resp, nil
}
