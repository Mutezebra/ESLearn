package es

import (
	"context"
	"fmt"
	"go-ElasticSearch/es/esmodel"
	"log"
)

func CreateIndex(indexName string, model esmodel.ESModels) (err error) {
	exist := IndexExist(indexName)
	if exist {
		err = DeleteIndex(indexName)
		if err != nil {
			return err
		}
	}
	_, err = ESClient.CreateIndex(indexName).
		BodyString(model.Mapping()).
		Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("create index success")
	return nil
}

func DeleteIndex(indexName string) (err error) {
	_, err = ESClient.DeleteIndex(indexName).Do(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("delete index success")
	return err
}

func IndexExist(indexName string) bool {
	exist, _ := ESClient.IndexExists(indexName).Do(context.TODO())
	return exist
}
