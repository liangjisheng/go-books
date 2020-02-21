package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

// Tweet ...
type Tweet struct {
	User    string
	Message string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://117.51.148.112:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	fmt.Println("conn es succ")
	tweet := Tweet{User: "olivere name", Message: "Take Five"}
	_, err = client.Index().
		Index("jw").
		// Type("tweet").
		Type("employee").
		Id("1").
		BodyJson(tweet).
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("insert succ")
}
