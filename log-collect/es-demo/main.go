package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = "http://117.51.148.112:9200/"

// Tweet ...
type Tweet struct {
	User    string
	Message string
}

func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		fmt.Println("new client err: ", err)
		os.Exit(-1)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func main() {
	tweet := Tweet{User: "olivere name", Message: "Take Five"}
	_, err := client.Index().Index("jw").Type("employee").Id("4").BodyJson(tweet).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("insert succ")
}
