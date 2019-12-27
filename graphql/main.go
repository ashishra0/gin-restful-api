package grapqhl

import (
	"fmt"
	"go_microservice/repository"
	"log"

	"github.com/shahidhk/gql"
)

func MutateQuery(v map[string]interface{}) {
	client := gql.NewClient("http://localhost:8080/v1/graphql", nil)
	var articles = repository.New()
	err := client.Execute(gql.Request{Query: `mutation update_articles($id: Int, $title: String, $body: String, $count: Int) {
		update_articles(
			where: {id: {_eq: $id}},
			_set: {
				count: $count
			}
		) {
			affected_rows
		}
	}`, Variables: v}, &articles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(articles)
}
