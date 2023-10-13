package es

import (
	"crypto/tls"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

var ESClient *elastic.Client

func InitES() {
	//cert, err := os.ReadFile("D:\\elastic\\elasticsearch-8.9.0\\ssl.pem")
	//if err != nil {
	//	log.Fatalf("error reading cert file: %s", err)
	//}
	//
	//log.Println(string(cert))
	//// 创建证书池并添加证书
	//certPool := x509.NewCertPool()
	//if ok := certPool.AppendCertsFromPEM(cert); !ok {
	//	log.Fatalf("failed to append cert")
	//}
	//
	//// 创建 TLS 配置并设置证书池
	//tlsConfig := &tls.Config{
	//	RootCAs:    certPool,
	//	ServerName: "localhost",
	//}
	//
	//// 创建 HTTP 客户端并配置 TLS 配置
	//httpClient := &http.Client{
	//	Transport: &http.Transport{
	//		TLSClientConfig: tlsConfig,
	//	},
	//}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 忽略证书验证
		},
	}
	httpClient := &http.Client{Transport: tr}
	es, err := elastic.NewClient(
		elastic.SetURL("https://localhost:9200"),
		elastic.SetBasicAuth("elastic", "07IDHsSqhM8nDMgXS9GP"),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
	)

	if err != nil {
		log.Fatalf("error creating the client %s\n", err)
	}
	log.Println("es init success")
	ESClient = es
}
