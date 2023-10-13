package main

import (
	"go-ElasticSearch/es"
	"go-ElasticSearch/es/esmodel"
	"log"
	"time"
)

func main() {
	es.InitES()
	user := esmodel.User{
		ID:       1,
		Name:     "mutezebra",
		Tag:      []string{"爱睡觉"},
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	resp, _ := es.CreateDoc(&user)
	log.Println(resp)
}
