package graphQL

import (
	"github.com/graph-gophers/graphql-go"
	"io/ioutil"
)

var Schema *graphql.Schema

func init() {
	b, err := ioutil.ReadFile("./graphQL/schema/schema.graphql")
	if err != nil {
		panic(err)
	}
	Schema, err = graphql.ParseSchema(string(b), &Resolver{})
	if err != nil {
		panic(err)
	}
}
