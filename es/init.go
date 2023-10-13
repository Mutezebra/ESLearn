package es

import (
	"crypto/tls"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

var ESClient *elastic.Client

func InitES() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 忽略证书验证
		},
	}
	httpClient := &http.Client{Transport: tr}
	es, err := elastic.NewClient(
		elastic.SetURL("https://localhost:9200"),
		elastic.SetBasicAuth("elastic", "password"),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
	)

	if err != nil {
		log.Fatalf("error creating the client %s\n", err)
	}
	log.Println("es init success")
	ESClient = es
}
