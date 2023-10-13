package es

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"go-ElasticSearch/es/esmodel"
	"log"
)

func DocCreate(model esmodel.ESModels) (*elastic.IndexResponse, error) {
	model.CreateTime() // 设置创建时间
	resp, err := ESClient.
		Index().Index(model.Index()). // 第一个index()可以理解为对index的操作
		BodyJson(model).
		Do(context.TODO())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("create doc success ", resp)
	return resp, nil
}

func DocUserBatchCreate(models []esmodel.User) (*elastic.BulkResponse, error) {
	if models == nil {
		return nil, errors.New("model is null")
	}

	bulk := ESClient.Bulk().Index(models[0].Index())
	for _, v := range models {
		v.CreateTime() // 设置创建时间
		req := elastic.NewBulkCreateRequest().Doc(v)
		bulk.Add(req)
	}

	resp, err := bulk.Do(context.TODO())
	if err != nil {
		return nil, err
	}
	log.Printf("%#v", resp)
	return resp, err
}

func DocDelete(id string, model esmodel.ESModels) (*elastic.DeleteResponse, error) {
	resp, err := ESClient.Delete().Index(model.Index()).Id(id).Do(context.TODO())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

func DocBatchDelete(ids []string, model esmodel.ESModels) (*elastic.BulkResponse, error) {
	bulk := ESClient.Bulk().Index(model.Index())

	for _, v := range ids {
		req := elastic.NewBulkDeleteRequest().Id(v)
		bulk.Add(req)
	}
	resp, err := bulk.Do(context.TODO())
	if err != nil {
		return nil, err
	}
	if resp.Succeeded() == nil {
		return nil, errors.New("elastic:delete nothing")
	}
	log.Println(resp.Succeeded())
	return resp, nil
}
